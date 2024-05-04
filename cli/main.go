package cli

import (
	"context"
	"log"
	"os"
	"seven/cli/config"
	"seven/cli/environment"
	"seven/cli/logs"
	"seven/cli/scripts"
	"seven/completion"
	"seven/model"
	"seven/prompt"
	"seven/types"
)

func Run(params types.Parameters, promptParams types.PromptParameters, modelParams types.ModelParameters) {

	ctx := context.Background()

	manifest, seven := config.GetManifestAndConfig(params)

	/*
		Create environment variables from the variables section of the manifest:

		```yaml
		variables:
		  - MESSAGE=Your name is Seven
		  - CHARACTER=KeegOrg
		  - AUTHOR=@k33g_org
		```
	*/
	environment.ReadVariablesSection(manifest)

	/*
		Model Settings - Ollama Settings
		Update the settings of the model if we override them with the option flags:
		--llm (for the llm name)
		--predictRepeatLastN

		```yaml
		model:
		  name: phi3
		  memory: false
		  stream: true
		settings:
			predictRepeatLastN: 64
		```
	*/
	model.UpdateModelSettingsWithParams(modelParams, &manifest)

	/*
		Prompt Settings
		Update the settings of the prompt if we override them with the option flags:
		--temperature
		--stopWords

		```yaml
		prompt:
		  settings:
		    temperature: 0.8
		    stopWords: []
		```
	*/
	prompt.UpdateSettingsWithParams(promptParams, &manifest)

	logs.ShowModelPromptInfo(params, manifest, seven)

	scripts.RunBeforeScriptIfExists(manifest)

	/*
		Environment variables substitution
		into the prompt messages
	*/
	prompt.EnvVarSubstMessages(&manifest)

	/*
		Load the LLM
		if isFunc == true -> we are using functions
	*/
	isFunc, llm, err := model.GetLLM(manifest, seven)

	logs.ShowIfWeAreUsingFunctions(params, isFunc)

	if err != nil {
		log.Fatal("😡 error loading the LLM: ", err)
	}

	/*
		Update the messages if we override them with the option flags:
		--question
		--context
		--system

		```yaml
		prompt:
		  settings:
			temperature: 0.8
			stopWords: []

		system: |
			You are an expert of Generative AI.
		context: |
			<context>
			</context>
		human: |
			How to prevent the "rise of the machines" by mastering the use of generative AI.
		```
	*/
	prompt.UpdateMessagesWithParams(promptParams, &manifest)

	logs.ShowMemoryUsageInfo(params, manifest)

	/*
		Initialize the conversational history:

		```yaml
		model:
		  name: phi3
		  memory: true
		  stream: true
		settings:
			predictRepeatLastN: 64
		```
	*/
	history, memDb, err := model.InitializeConversationalHistory(ctx, params, manifest)
	if err != nil {
		log.Fatal("😡 error initializing the conversational history: ", err)
	}

	switch composition := manifest.PromptComposedBy(); composition {

	case types.HumanOnly:
		completionStr, err := completion.GenerateWithHumanPrompt(ctx, llm, manifest, history, memDb, params.OutputPath, params.ShowLogs)
		if err != nil {
			log.Fatal("😡 error generating completion: ", err)
		}
		// Completion result
		os.Setenv("SEVEN_COMPLETION", completionStr)

	case types.Summary:
		completionStr, err := completion.GenerateSummary(ctx, llm, manifest, params.ShowLogs, params.OutputPath)
		if err != nil {
			log.Fatal("😡 error generating summary: ", err)
		}
		// Completion result
		os.Setenv("SEVEN_COMPLETION", completionStr)

	case types.SystemOnly:
		log.Fatal("😡 prompt composed only by system message")
	case types.ContextOnly:
		log.Fatal("😡 prompt composed only by context message")
	case types.SystemAndContext:
		log.Fatal("😡 prompt composed only by system and context message")

	case types.SystemAndHuman:
		completionStr, err := completion.GenerateWithSystemHumanPrompt(ctx, llm, manifest, history, memDb, params.OutputPath, params.ShowLogs)
		if err != nil {
			log.Fatal("😡 error generating completion: ", err)
		}
		// Completion result
		os.Setenv("SEVEN_COMPLETION", completionStr)

	case types.ContextAndHuman:
		log.Fatal("😡 prompt composed only by context and human message")

	case types.SystemAndContextAndHuman:

		var completionStr = ""
		if manifest.Documents != nil {
			// ! conversational memory not implemented for this part
			completionStr, err = completion.GenerateWithStuffDocuments(ctx, llm, manifest, params.ShowLogs, params.OutputPath)
			if err != nil {
				log.Fatal("😡 error generating completion: ", err)
			}
		} else {
			// ! conversational memory is implemented for this part
			completionStr, err = completion.GenerateWithSystemContextHumanPrompt(ctx, llm, manifest, history, memDb, params.OutputPath, params.ShowLogs)
			if err != nil {
				log.Fatal("😡 error generating completion: ", err)
			}
		}
		// Completion result
		os.Setenv("SEVEN_COMPLETION", completionStr)

	case types.Functions:
		//! no stream
		completionStr, err := completion.GenerateWithFunction(ctx, llm, manifest, params)
		if err != nil {
			log.Fatal("😡 error generating completion: ", err)
		}
		os.Setenv("SEVEN_COMPLETION", completionStr)

	case types.NoComposition:
		log.Fatal("😡 something bad is happening")
	}

	scripts.RunAfterScriptIfExists(manifest)
}
