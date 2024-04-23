// Package cmd :
package cmd

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"seven/cli"
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

	case "apply":
		// options / flags
		flagSet := flag.NewFlagSet("apply", flag.ExitOnError)
		manifest := flagSet.String("manifest", "", "Seven yaml manifest")
		envfile := flagSet.String("envfile", ".env", "Environmant variable file")
		/*
		  If you define an environment variable before running the command,
		  it will override the value from the .env file
		*/
		question := flagSet.String("question", "", "Question to ask")
		/*
		  if empty use the yaml manifest
		*/
		output := flagSet.String("output", "", "Output file")

		logs := flagSet.Bool("logs", false, "Show logs")

		context := flagSet.String("context", "", "Prompt context")
		system := flagSet.String("system", "", "Prompt system")

		config := flagSet.String("config", os.Getenv("SEVENCONFIG"), "Configuration file - by default use the SEVENCONFIG environment variable")

		// -------------------------------------
		//   Prompt Settings Configuration
		// -------------------------------------
		// TODO: add more settings
		/* See:
			- robot/01-simple.yaml
			- Source code:
			  - types/manifest.go
			  - usage (initialization): cli/main.go
		*/
		temperature := flagSet.Float64("temperature", 0.0, "Temperature")
		stopWords := flagSet.String("stopWords", "", "Stop words")

		// -------------------------------------
		//   LLM Settings Configuration
		// -------------------------------------
		// TODO: add more settings
		/* See:
			- robot/01-simple.yaml
			- Source code:
			  - types/manifest.go
			  - usage (initialization): cli/main.go

			From the LangChainGo source code:
			https://github.com/tmc/langchaingo/blob/main/llms/ollama/options.go#L230
			WithPredictRepeatLastN Sets how far back for the model 
			to look back to prevent repetition
			(Default: 64, 0 = disabled, -1 = num_ctx).
		*/
		predictRepeatLastN := flagSet.Int("predictRepeatLastN", 64, "PredictRepeatLastN")

		//flagSet.Parse(args[0:])
		flagSet.Parse(args)

		cli.Run(
			*manifest,
			*envfile,
			*question,
			*output,
			*logs,
			*context,
			*system,
			*temperature,
			*stopWords,
			*predictRepeatLastN,
			*config,
		)

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
