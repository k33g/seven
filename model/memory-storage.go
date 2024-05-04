package model

import (
	"context"
	"seven/types"

	"github.com/tmc/langchaingo/memory"
	"go.etcd.io/bbolt"
)

// InitializeCoversationalHistory initializes the conversational history for the given manifest and parameters.
//
// Parameters:
// - ctx: The context.Context for the function.
// - params: The types.Parameters for the function.
// - manifest: The types.Manifest for the function.
//
// Returns:
// - *memory.ChatMessageHistory: The initialized conversational history.
// - *bbolt.DB: The BoltDB database used for storing the conversational history.
// - error: An error if there was a problem initializing the conversational history.
func InitializeConversationalHistory(ctx context.Context, params types.Parameters, manifest types.Manifest) (*memory.ChatMessageHistory, *bbolt.DB, error) {
	if manifest.Model.Memory {
		history, memDb, err := GetConversationalHistory(ctx, manifest, params.ShowLogs)
		if err != nil {
			return nil, nil, err
			//log.Fatal("😡 error getting the conversation history: ", err)
		}
		return history, memDb, nil
	}
	return nil, nil, nil

}
