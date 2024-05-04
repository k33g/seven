package completion

import (
	"context"
	"fmt"
	"seven/model"
	"seven/types"

	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/schema"
)

func GenerateWithStuffDocuments(ctx context.Context, llm *ollama.LLM, manifest types.Manifest, showLogs bool, outputPath string) (string, error) {
	// ! conversational memory not implemented for this part
	if showLogs {
		fmt.Println("🤖 system:", manifest.Prompt.System)
		fmt.Println("📝 context:", manifest.Prompt.Context)
		fmt.Println("🤓 human:", manifest.Prompt.Human)
	}

	/*
		type Document struct {
			PageContent string
			Metadata    map[string]any
			Score       float32
		}
	*/
	var docs []schema.Document
	for _, item := range manifest.Documents {

		docs = append(docs, schema.Document{
			PageContent: item.Content,
			Metadata:    map[string]any{},
			Score:       0.0,
		})
	}

	if showLogs {
		fmt.Println("📝📝📝", docs)
		fmt.Println("")
	}

	// TODO: try something like this:
	/*
		prompts.NewGenericMessagePromptTemplate(
			"context",
			"{{ .context }}",
			[]string{"context"},
		),
	*/

	systemPromptTemplate := prompts.NewSystemMessagePromptTemplate(
		manifest.Prompt.System,
		nil,
	)
	contextPromptTemplate := prompts.NewSystemMessagePromptTemplate(
		manifest.Prompt.Context,
		[]string{"context"}, // link because chains.NewStuffDocuments
	)
	/*
		StuffDocuments is a chain that combines documents with a separator and uses
		the stuffed documents in an LLMChain. The input values to the llm chain
		contains all input values given to this chain, and the stuffed document as
		a string in the key specified by the "DocumentVariableName" field that is
		by default set to "context".

		See https://github.com/tmc/langchaingo/blob/main/chains/stuff_documents.go#L18
	*/
	humanPromptTemplate := prompts.NewHumanMessagePromptTemplate(
		manifest.Prompt.Human,
		nil,
	)
	templateList := []prompts.MessageFormatter{}

	templateList = append(templateList, systemPromptTemplate)
	templateList = append(templateList, contextPromptTemplate)
	templateList = append(templateList, humanPromptTemplate)

	prompt := prompts.NewChatPromptTemplate(templateList)

	completion, err := model.ChainsCall(ctx, llm, manifest, prompt, docs, manifest.Model.Stream, outputPath)
	if err != nil {
		return "", err
	}
	return completion, nil

}
