package model

import (
	"context"
	"encoding/json"
	"fmt"
	"seven/db"
	"seven/types"

	"github.com/google/uuid"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/prompts"
	"go.etcd.io/bbolt"
)

func GetConversationalHistory(ctx context.Context, manifest types.Manifest, showLogs bool) (*memory.ChatMessageHistory, *bbolt.DB, error) {
	memDb, err := db.Initialize(manifest.Model.MemoryStore)
	if err != nil {
		return nil, nil, err
		//log.Fatal("😡 error initializing the memory db: ", err)
	}
	history := memory.NewChatMessageHistory()

	// Load the history data
	conversation := db.GetAll(memDb)
	for key, value := range conversation {
		if showLogs {
			fmt.Println("  📕:", key, value)
		}
		/*
		  value {kind: "human", text: "Who is his best friend?"}
		  value {kind: "ai", text: "My best friend is my mom."}
		*/
		record := map[string]string{}
		err := json.Unmarshal([]byte(value), &record)
		if err != nil {
			return nil, nil, err
			//log.Fatal("😡 error unmarshalling the record: ", err)
		}
		if record["kind"] == "human" {
			history.AddUserMessage(ctx, record["text"])
		} else {
			history.AddAIMessage(ctx, record["text"])
		}
	}
	if len(conversation) == 0 {
		history.AddAIMessage(ctx, "This is the start of the conversation.")
		//? should we save it or not?
	}
	return history, memDb, nil
}

func GetHistoryPromptTemplate() prompts.GenericMessagePromptTemplate {
	historyPromptTemplate := prompts.NewGenericMessagePromptTemplate(
		"history",
		"{{range .historyMessages}}{{.GetContent}}\n{{end}}",
		[]string{"history"},
	)
	return historyPromptTemplate
}

func GetHistoryMessages(ctx context.Context, history *memory.ChatMessageHistory, showLogs bool) ([]llms.ChatMessage, error) {
	historyMessages, err := history.Messages(ctx)

	if err != nil {
		return nil, err
		//log.Fatal("😡 error when getting history:", err)
	}
	if showLogs {
		fmt.Println("📙History:")
		for idx, message := range historyMessages {
			fmt.Println("  ", idx, "-", message)
		}
		fmt.Println()
	}
	return historyMessages, nil
}

func SaveMessagesToHistory(question, answer string, saveQuestion bool, memDb *bbolt.DB) error {

	if saveQuestion {
		humanRecord, err := json.Marshal(map[string]string{"kind": "human", "text": question})
		if err != nil {
			return err
		}
		err = db.Save(memDb, uuid.Must(uuid.NewRandom()).String(), string(humanRecord))
		if err != nil {
			return err
		}

	}

	aiRecord, err := json.Marshal(map[string]string{"kind": "ai", "text": answer})
	if err != nil {
		return err
	}

	err = db.Save(memDb, uuid.Must(uuid.NewRandom()).String(), string(aiRecord))
	if err != nil {
		return err
	}

	return nil
}
