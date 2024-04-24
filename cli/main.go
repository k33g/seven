package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"seven/model"
	"seven/tools"
	"seven/types"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/memory"
	"go.etcd.io/bbolt"
)

func Run(manifestPath string, envfilePath string, question string, outputPath string, showLogs bool, promptContext string, promptSystem string, temperature float64, stopWords string, predictRepeatLastN int, configPath string) {
	ctx := context.Background()

	seven, err := YamlFileToSeven(configPath)
	if err != nil {
		log.Fatal("😡 error loading config file: ", err)
	}

	manifest, err := YamlFileToManifest(manifestPath)
	if err != nil {
		log.Fatal("😡 error loading .yaml file: ", err)
	}

	/*
	  Environment variables substitution
	*/
	if manifest.Prompt.System != "" {
		manifest.Prompt.System, err = tools.EnvSubst(manifest.Prompt.System)
		if err != nil {
			log.Fatal("😡 error with envsubst: ", err)
		}
	}
	if manifest.Prompt.Human != "" {
		manifest.Prompt.Human, err = tools.EnvSubst(manifest.Prompt.Human)
		if err != nil {
			log.Fatal("😡 error with envsubst: ", err)
		}
	}
	if manifest.Prompt.Context != "" {
		manifest.Prompt.Context, err = tools.EnvSubst(manifest.Prompt.Context)
		if err != nil {
			log.Fatal("😡 error with envsubst: ", err)
		}
	}

	if manifest.BeforeScript != "" {
		cmd := exec.Command("sh", "-c", manifest.BeforeScript)
		err = cmd.Run()
		if err != nil {
			log.Fatal("😡 error running before script: ", err)
		}
	}

	// Ollama Settings
	if predictRepeatLastN != 64 {
		manifest.Model.Settings.PredictRepeatLastN = predictRepeatLastN
	}

	ollamaURL := seven.Config.Ollama.URL
	if ollamaURL == "" {
		ollamaURL = "http://localhost:11434"
		//ollamaURL = "http://host.docker.internal:11434"
	}

	// Prompt Settings
	if temperature != 0 {
		manifest.Prompt.Settings.Temperature = temperature
	}

	if stopWords != "" {
		manifest.Prompt.Settings.StopWords = strings.Split(stopWords, ",")
	}

	if showLogs {
		fmt.Println("🤭 predictRepeatLastN:", manifest.Model.Settings.PredictRepeatLastN)
		fmt.Println("🌡️  temperature:", manifest.Prompt.Settings.Temperature)
		fmt.Println("🚫 stop words:", manifest.Prompt.Settings.StopWords)
		fmt.Println()
	}

	// ! here test if there is a list of tools / functions
	var llm *ollama.LLM
	if manifest.Functions != nil {
		llm, err = ollama.New(
			ollama.WithModel(manifest.Model.Name),
			ollama.WithServerURL(ollamaURL),
			ollama.WithPredictRepeatLastN(manifest.Model.Settings.PredictRepeatLastN),
			ollama.WithFormat("json"),
		)
	} else {
		llm, err = ollama.New(
			ollama.WithModel(manifest.Model.Name),
			ollama.WithServerURL(ollamaURL),
			ollama.WithPredictRepeatLastN(manifest.Model.Settings.PredictRepeatLastN),
		)
	}

	if err != nil {
		log.Fatal("😡 error loading the LLM: ", err)
	}

	if question != "" {
		manifest.Prompt.Human = question + "\n"
	}
	if promptContext != "" {
		manifest.Prompt.Context = promptContext + "\n"
	}
	if promptSystem != "" {
		manifest.Prompt.System = promptSystem + "\n"
	}

	if showLogs {
		fmt.Println("🧠 model:", manifest.Model.Name)
		if manifest.Model.Memory {
			fmt.Println("🧠 memory:", manifest.Model.MemoryStore)
		}
	}

	var memDb *bbolt.DB
	var history *memory.ChatMessageHistory
	// TODO:
	// ? how to manage the remove of memory db file
	// * Initialize the conversational history
	if manifest.Model.Memory {
		history, memDb, err = model.GetConversationalHistory(ctx, manifest, showLogs)
		if err != nil {
			log.Fatal("😡 error getting the conversation history: ", err)
		}
	}

	switch composition := manifest.PromptComposedBy(); composition {

	case types.HumanOnly:
		_, err := GenerateCompletionWithHumanPrompt(ctx, llm, manifest, history, memDb, outputPath, showLogs)
		if err != nil {
			log.Fatal("😡 error generating completion: ", err)
		}

	case types.SystemOnly:
		log.Fatal("😡 prompt composed only by system message")
	case types.ContextOnly:
		log.Fatal("😡 prompt composed only by context message")
	case types.SystemAndContext:
		log.Fatal("😡 prompt composed only by system and context message")

	case types.SystemAndHuman:
		_, err := GenerateCompletionWithSystemHumanPrompt(ctx, llm, manifest, history, memDb, outputPath, showLogs)
		if err != nil {
			log.Fatal("😡 error generating completion: ", err)
		}

	case types.ContextAndHuman:
		log.Fatal("😡 prompt composed only by context and human message")

	case types.SystemAndContextAndHuman:

		//var prompt prompts.ChatPromptTemplate
		// TODO:
		// - documents from database ?
		// - documents from textfiles ?
		// ! or handle this with embeddings

		if manifest.Documents != nil {
			// ! conversational memory not implemented for this part
			_, err = GenerateCompletionWithStuffDocuments(ctx, llm, manifest, showLogs, outputPath)
			if err != nil {
				log.Fatal("😡 error generating completion: ", err)
			}
		} else {
			// ! conversational memory is implemented for this part
			_, err = GenerateCompletionWithSystemContextHumanPrompt(ctx, llm, manifest, history, memDb, outputPath, showLogs)
			if err != nil {
				log.Fatal("😡 error generating completion: ", err)
			}
		}

	case types.Functions:

		/*
			if showLogs {
				fmt.Println("🧮 functions:")
				fmt.Println(manifest.Functions)
			}
		*/

		var functionsList = []llms.FunctionDefinition{}

		for _, f := range manifest.Functions {

			functionsList = append(functionsList, llms.FunctionDefinition{
				Name:        f.Name,
				Description: f.Description,
				Parameters:  f.Parameters, //? pretty print or not???
			})

		}

		toolsList, err := json.Marshal(functionsList)
		if err != nil {
			log.Fatal("😡 error when generation json string tools list:", err)
		}
		toolsListStr := string(toolsList)

		var messagesContent []llms.MessageContent

		systemMessage := strings.Replace(manifest.Prompt.System, "{{ .functions }}", toolsListStr, 1)

		messagesContent = append(
			messagesContent,
			llms.TextParts(llms.ChatMessageTypeSystem, systemMessage),
		)
		// it exists a ChatMessageTypeTool

		if manifest.Prompt.Context != "" {
			messagesContent = append(
				messagesContent,
				llms.TextParts(llms.ChatMessageTypeSystem, manifest.Prompt.Context),
			)
		}

		messagesContent = append(
			messagesContent,
			llms.TextParts(llms.ChatMessageTypeHuman, manifest.Prompt.Human),
		)

		//TODO: how to stream?
		response, err := llm.GenerateContent(ctx, messagesContent)
		if err != nil {
			log.Fatal("😡 issue when generating content:", err)
		}

		firstChoice := response.Choices[0].Content

		selectedTool := map[string]any{}
		errTool := json.Unmarshal([]byte(firstChoice), &selectedTool)

		if errTool != nil {
			log.Fatal("😡 issue when unmarshalling:", errTool)
		}

		/*
			args, err := tools.AnyToJsonString(selectedTool["arguments"])
			if err != nil {
				log.Fatal("😡 error when generation json string:", err)
			}
			args, err = tools.CleanJsonString(args)
			if err != nil {
				log.Fatal("😡 error when cleaning json string:", err)
			}
		*/

		if showLogs {
			fmt.Println("🛠️ selected tool:")
			fmt.Println(" 🔵 function  ->", selectedTool["function"])
			fmt.Println(" 🟢 arguments ->", selectedTool["arguments"])
			fmt.Println()
		}

		// Call the function
		for _, f := range manifest.Functions {
			if f.Name == selectedTool["function"].(string) {
				if showLogs {
					fmt.Println("🚀 run:", f.Cmd)
				}
				args, err := tools.AnyToJsonString(selectedTool["arguments"])
				if err != nil {
					log.Fatal("😡 error when generation json string:", err)
				}

				parts := strings.Fields(f.Cmd)
				data, err := exec.Command(parts[0], parts[1], args).Output()
				/*
					cmd := exec.Command(f.Cmd, args)
					stdout, err := cmd.Output()
				*/

				if err != nil {
					log.Fatal("😡 error when executing the command:", err)
					return
				}
				// Print the output
				//fmt.Println(string(stdout))
				fmt.Println(string(data))
				break
			}
		}

	// TODO:

	// - some gardening 🌸
	// - with summarization
	// - with more settings

	case types.NoComposition:
		log.Fatal("😡 something bad is happening")
	}

	if manifest.AfterScript != "" {
		//fmt.Println("🚀 after script:")
		cmd := exec.Command("sh", "-c", manifest.AfterScript)
		stdout, err := cmd.Output()
		//err = cmd.Run()
		if err != nil {
			log.Fatal("😡 error running after script: ", err)
		}
		fmt.Println(string(stdout))
	}
}
