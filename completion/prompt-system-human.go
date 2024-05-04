package completion

import (
	"context"
	"fmt"
	"seven/model"
	"seven/types"

	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/prompts"
	"go.etcd.io/bbolt"
)

func GenerateWithSystemHumanPrompt(ctx context.Context, llm *ollama.LLM, manifest types.Manifest, history *memory.ChatMessageHistory, memDb *bbolt.DB, outputPath string, showLogs bool) (string, error) {
	if showLogs {
		fmt.Println("🤖 system:", manifest.Prompt.System)
		fmt.Println("🤓 human:", manifest.Prompt.Human)
	}

	systemPromptTemplate := prompts.NewSystemMessagePromptTemplate(
		manifest.Prompt.System,
		nil,
	)

	humanPromptTemplate := prompts.NewHumanMessagePromptTemplate(
		manifest.Prompt.Human,
		nil,
	)

	templateList := []prompts.MessageFormatter{}

	templateList = append(templateList, systemPromptTemplate)

	// add memory here and insert history
	if manifest.Model.Memory {
		templateList = append(templateList, model.GetHistoryPromptTemplate())
	}

	templateList = append(templateList, humanPromptTemplate)

	prompt := prompts.NewChatPromptTemplate(templateList)

	var promptText string

	if manifest.Model.Memory {

		historyMessages, err := model.GetHistoryMessages(ctx, history, showLogs)
		if err != nil {
			//log.Fatal("😡 error when getting history:", err)
			return "", err
		}

		// build prompt with the history of the conversation
		promptText, err = prompt.Format(map[string]interface{}{"historyMessages": historyMessages})
		if err != nil {
			//log.Fatal("😡 error when generating the prompt:", err)
			return "", err
		}

	} else {
		var err error
		promptText, err = prompt.Format(map[string]interface{}{})
		if err != nil {
			//log.Fatal("😡 error when generating the prompt:", err)
			return "", err
		}
	}

	answer, err := model.GenerateFromSinglePrompt(ctx, llm, manifest, promptText, outputPath)
	if err != nil {
		//log.Fatal("😡 ", err)
		return "", err
	}

	// Save the history data to the database
	if manifest.Model.Memory {
		err := model.SaveMessagesToHistory(manifest.Prompt.Human, answer, false, memDb)
		if err != nil {
			//log.Fatal("😡 error when saving the history:", err)
			return "", err
		}
	}
	return answer, nil
}
