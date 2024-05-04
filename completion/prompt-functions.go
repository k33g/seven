package completion

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"seven/tools"
	"seven/types"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func GenerateWithFunction(ctx context.Context, llm *ollama.LLM, manifest types.Manifest, params types.Parameters) (string, error) {

	var functionsList = []llms.FunctionDefinition{}

	// Read the functions list from the manifest
	for _, f := range manifest.Functions {
		functionsList = append(functionsList, llms.FunctionDefinition{
			Name:        f.Name,
			Description: f.Description,
			Parameters:  f.Parameters, //? pretty print or not???
		})
	}

	// Create  the tools list
	toolsList, err := json.Marshal(functionsList)
	if err != nil {
		return "", err
		//log.Fatal("😡 error when generation json string tools list:", err)
	}
	toolsListStr := string(toolsList)

	var messagesContent []llms.MessageContent

	systemMessage := strings.Replace(manifest.Prompt.System, "{{ .functions }}", toolsListStr, 1)

	messagesContent = append(
		messagesContent,
		llms.TextParts(llms.ChatMessageTypeSystem, systemMessage),
	)
	// TODO: it exists a ChatMessageTypeTool, test it

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
		return "", err
		//log.Fatal("😡 issue when generating content:", err)
	}
	firstChoice := response.Choices[0].Content

	selectedTool := map[string]any{}
	errTool := json.Unmarshal([]byte(firstChoice), &selectedTool)

	if errTool != nil {
		return "", err
		//log.Fatal("😡 issue when unmarshalling:", errTool)
	}

	// TODO: put this in a function
	if params.ShowLogs {
		fmt.Println("🛠️ selected tool:")
		fmt.Println(" 🔵 function  ->", selectedTool["function"])
		fmt.Println(" 🟢 arguments ->", selectedTool["arguments"])
		fmt.Println()
	}

	var completion string = ""
	// Call the function
	for _, f := range manifest.Functions {
		if f.Name == selectedTool["function"].(string) {
			if params.ShowLogs {
				fmt.Println("🚀 run:", f.Cmd)
			}
			args, err := tools.AnyToJsonString(selectedTool["arguments"])
			if err != nil {
				return "", err
				//log.Fatal("😡 error when generation json string:", err)
			}

			parts := strings.Fields(f.Cmd)
			data, err := exec.Command(parts[0], parts[1], args).Output()
			/*
				cmd := exec.Command(f.Cmd, args)
				stdout, err := cmd.Output()
			*/

			if err != nil {
				return "", err
				///log.Fatal("😡 error when executing the command:", err)
			}

			// Completion result
			completion = string(data)
			fmt.Println(completion)
			break
		}
	}

	return completion, nil
}
