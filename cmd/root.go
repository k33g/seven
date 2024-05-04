// Package cmd :
package cmd

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"seven/cli"
	"seven/types"
)

//go:embed config/version.txt
var version []byte

//go:embed config/about.txt
var about []byte

//go:embed config/help.md
var help []byte

// Parse de CLI args and flags
func Parse(command string, args []string) error {

	switch command {
	case "directive":
		// 🚧 simplify Seven - wip
	case "apply":
		// options / flags
		flagSet := flag.NewFlagSet("apply", flag.ExitOnError)

		manifest := flagSet.String("manifest", "", "Seven yaml manifest")
		envfile := flagSet.String("envfile", ".env", "Environent variable file") // not used
		config := flagSet.String("config", os.Getenv("SEVENCONFIG"), "Configuration file - by default use the SEVENCONFIG environment variable")
		output := flagSet.String("output", "", "Output file")
		logs := flagSet.Bool("logs", false, "Show logs")

		/*
		  If you define an environment variable before running the command,
		  it will override the value from the .env file
		*/
		// -------------------------------------
		//   Prompt Construction
		// -------------------------------------
		question := flagSet.String("question", "", "Question to ask")
		context := flagSet.String("context", "", "Prompt context")
		system := flagSet.String("system", "", "Prompt system")

		// -------------------------------------
		//   Prompt Settings Configuration
		// -------------------------------------
		// TODO: add more settings
		temperature := flagSet.Float64("temperature", 0.0, "Temperature")
		stopWords := flagSet.String("stopWords", "", "Stop words")

		// -------------------------------------
		//   LLM Settings Configuration
		// -------------------------------------
		// TODO: add more settings
		/* See:
		From the LangChainGo source code:
		https://github.com/tmc/langchaingo/blob/main/llms/ollama/options.go#L230
		WithPredictRepeatLastN Sets how far back for the model
		to look back to prevent repetition
		(Default: 64, 0 = disabled, -1 = num_ctx).
		*/
		modelName := flagSet.String("llm", "", "LLM name")
		predictRepeatLastN := flagSet.Int("predictRepeatLastN", 64, "PredictRepeatLastN")


		//flagSet.Parse(args[0:])
		flagSet.Parse(args)

		params := types.Parameters{
			ManifestPath: *manifest,
			EnvfilePath:  *envfile,
			OutputPath:   *output,
			ShowLogs:     *logs,
			ConfigPath:   *config,
		}

		prompParams := types.PromptParameters{
			Question:    *question,
			Context:     *context,
			System:      *system,
			Temperature: *temperature,
			StopWords:   *stopWords,
		}

		modelParams := types.ModelParameters{
			Name:               *modelName,
			PredictRepeatLastN: *predictRepeatLastN,
		}

		cli.Run(params, prompParams, modelParams)

		return nil

	case "listen":
		fmt.Println("🚧 it's a work in progress...")
		return nil

	case "version":
		fmt.Println(string(version))
		return nil

	case "about":
		fmt.Println(string(about))
		return nil

	case "help":
		fmt.Println(string(help))
		return nil

	default:
		return fmt.Errorf("😡 invalid command")
	}
}
