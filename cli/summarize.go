package cli

import (
	"context"
	"fmt"
	"os"
	"seven/types"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
)

func GenerateSummary(ctx context.Context, llm *ollama.LLM, manifest types.Manifest, showLogs bool, outputPath string) (string, error) {

	if showLogs {
		fmt.Println("📝 splitter:", manifest.Summary.Splitter)
		fmt.Println("🔪 chunk:", manifest.Summary.Chunk)
	}

	if manifest.Summary.Chunk.Overlap == 0 {
		manifest.Summary.Chunk.Overlap = 128
	}
	if manifest.Summary.Chunk.Size == 0 {
		manifest.Summary.Chunk.Size = 1536
	}

	var splitter textsplitter.TextSplitter
	switch manifest.Summary.Splitter {
	case "markdown":
		splitter = textsplitter.NewMarkdownTextSplitter(
			textsplitter.WithChunkSize(manifest.Summary.Chunk.Size),
			textsplitter.WithChunkOverlap(manifest.Summary.Chunk.Overlap),
		)
	case "text":
		splitter = textsplitter.NewRecursiveCharacter(
			textsplitter.WithChunkSize(manifest.Summary.Chunk.Size),
			textsplitter.WithChunkOverlap(manifest.Summary.Chunk.Overlap),
		)
	default:
		splitter = textsplitter.NewRecursiveCharacter(
			textsplitter.WithChunkSize(manifest.Summary.Chunk.Size),
			textsplitter.WithChunkOverlap(manifest.Summary.Chunk.Overlap),
		)
	}

	fmt.Println(splitter)

	chunks, err := splitter.SplitText(manifest.Summary.Document)
	if err != nil {
		return "", err
	}
	var docs []schema.Document
	for _, chunk := range chunks {
		docs = append(docs, schema.Document{
			PageContent: chunk, // 🤔 should I add <doc></doc>?
			Metadata:    map[string]any{},
			Score:       0.0,
		})
	}

	if showLogs {
		fmt.Println("📝 documents:", docs)
	}

	llmSummarizationChain := chains.LoadRefineSummarization(llm)

	var completion string
	if manifest.Model.Stream {
		//var completion string
		fmt.Println("--------------------------------------------------")

		outputValues, err := chains.Call(
			ctx,
			llmSummarizationChain,
			map[string]any{"input_documents": docs},
			chains.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
				//completion += string(chunk)
				fmt.Print(string(chunk))
				if len(chunk) == 0 {
					fmt.Println("")
					fmt.Println("--------------------------------------------------")
				}

				return nil
			}),
		)
		if err != nil {
			return "", err
		}
		completion = outputValues["text"].(string)
		fmt.Println("📝 Last summary:")
		fmt.Println(completion)
	} else {

		outputValues, err := chains.Call(
			ctx,
			llmSummarizationChain,
			map[string]any{"input_documents": docs},
		)
		if err != nil {
			return "", err
		}
		completion = outputValues["text"].(string)
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
