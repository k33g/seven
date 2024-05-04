package model

import (
	"seven/types"

	"github.com/tmc/langchaingo/llms/ollama"
)

// GetLLMForFunctions returns an instance of ollama.LLM and an error.
//
// It takes a manifest of type types.Manifest and seven of type types.Seven as parameters.
// The manifest parameter contains information about the model name and settings.
// The seven parameter contains the configuration for the Ollama server URL.
// The function uses the ollama.New function to create a new instance of ollama.LLM.
// It sets the model name using the manifest parameter, the server URL using the seven parameter,
// and the predict repeat last N value using the manifest settings.
// The function also sets the format to "json".
// It returns the created ollama.LLM instance and any error that occurred during the creation.
func GetLLMForFunctions(manifest types.Manifest, seven types.Seven) (*ollama.LLM, error) {
	llm, err := ollama.New(
		ollama.WithModel(manifest.Model.Name),
		ollama.WithServerURL(seven.Config.Ollama.URL),
		ollama.WithPredictRepeatLastN(manifest.Model.Settings.PredictRepeatLastN),
		ollama.WithFormat("json"),
	)
	return llm, err
}

// GetLLMForCompletion returns an instance of ollama.LLM and an error.
//
// It takes a manifest of type types.Manifest and seven of type types.Seven as parameters.
// The manifest parameter contains information about the model name and settings.
// The seven parameter contains the configuration for the Ollama server URL.
// The function uses the ollama.New function to create a new instance of ollama.LLM.
// It sets the model name using the manifest parameter, the server URL using the seven parameter,
// and the predict repeat last N value using the manifest settings.
// It returns the created ollama.LLM instance and any error that occurred during the creation.
func GetLLMForCompletion(manifest types.Manifest, seven types.Seven) (*ollama.LLM, error) {
	llm, err := ollama.New(
		ollama.WithModel(manifest.Model.Name),
		ollama.WithServerURL(seven.Config.Ollama.URL),
		ollama.WithPredictRepeatLastN(manifest.Model.Settings.PredictRepeatLastN),
	)
	return llm, err
}

// GetLLM returns an instance of ollama.LLM and an error.
//
// It takes a manifest of type types.Manifest and seven of type types.Seven as parameters.
// The manifest parameter contains information about the model name and settings.
// The seven parameter contains the configuration for the Ollama server URL.
// The function returns a boolean value indicating whether the LLM is for functions or completion,
// the LLM instance, and any error that occurred during the creation.
func GetLLM(manifest types.Manifest, seven types.Seven) (bool, *ollama.LLM, error) {
	if manifest.Functions != nil {
		llm, err := GetLLMForFunctions(manifest, seven)
		return true, llm, err
	}
	llm, err := GetLLMForCompletion(manifest, seven)
	return false, llm, err
}
