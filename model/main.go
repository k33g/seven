package model

import (
	"context"
	"fmt"
	"os"
	"seven/types"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/schema"
)

func GenerateFromSinglePrompt(ctx context.Context, llm *ollama.LLM, manifest types.Manifest, promptText string, outputPath string) (string, error) {
	var completion string
	if manifest.Model.Stream {

		_, err := llms.GenerateFromSinglePrompt(ctx, llm, promptText,
			llms.WithTemperature(manifest.Prompt.Settings.Temperature),
			llms.WithStopWords(manifest.Prompt.Settings.StopWords),
			llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
				fmt.Print(string(chunk))
				completion += string(chunk)
				return nil
			}))
		if err != nil {
			return "", err
			//log.Fatal("😡 error with the stream: ", err)
		}
		fmt.Println("")

	} else {

		completion, err := llms.GenerateFromSinglePrompt(ctx, llm, promptText,
			llms.WithTemperature(manifest.Prompt.Settings.Temperature),
			llms.WithStopWords(manifest.Prompt.Settings.StopWords),
		)
		if err != nil {
			return "", err
			//log.Fatal("😡 error with the single prompt: ", err)
		}
		fmt.Println("")
		fmt.Println(completion)
	}

	if outputPath != "" {
		data := []byte(completion + "\n")
		err := os.WriteFile(outputPath, data, 0644)
		if err != nil {
			return "", err
			//log.Fatal("😡 error writing the file: ", err)
		}
	}
	return completion, nil
}

func ChainsCall(ctx context.Context, llm *ollama.LLM, manifest types.Manifest, prompt prompts.ChatPromptTemplate, docs []schema.Document, isStream bool, outputPath string) (string, error) {

	llmChain := chains.NewLLMChain(llm, prompt)
	chain := chains.NewStuffDocuments(llmChain)

	var completion string
	if isStream {

		// ajouter "history" à la map

		_, err := chains.Call(
			context.Background(), chain, //! the conversational memory will be add here ???
			map[string]any{
				"input_documents": docs, // link to context (input_documents?🤔)
			},
			chains.WithTemperature(manifest.Prompt.Settings.Temperature),
			chains.WithStopWords(manifest.Prompt.Settings.StopWords),
			chains.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
				fmt.Print(string(chunk))
				completion += string(chunk)
				return nil
			}),
		)
		if err != nil {
			return "", err
		}
		fmt.Println("")

	} else {
		answer, err := chains.Call(context.Background(), chain, map[string]any{
			"input_documents": docs, // link to context (input_documents?🤔)
		},
			chains.WithTemperature(manifest.Prompt.Settings.Temperature),
			chains.WithStopWords(manifest.Prompt.Settings.StopWords),
		)
		
		if err != nil {
			return "", err
		}
		fmt.Println("")
		completion = answer["text"].(string)
		fmt.Println(completion)
	}
	return completion, nil
}
