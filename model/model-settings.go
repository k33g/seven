package model

import "seven/types"

// UpdateModelSettingsWithParams updates the model settings in the manifest with the values from the model parameters.
//
// Parameters:
// - modelParams: The model parameters containing the new values for the model settings.
// - manifest: A pointer to the manifest struct where the model settings will be updated.
func UpdateModelSettingsWithParams(modelParams types.ModelParameters, manifest *types.Manifest) {
	// Model Settings - Ollama Settings
	if modelParams.PredictRepeatLastN != 64 {
		manifest.Model.Settings.PredictRepeatLastN = modelParams.PredictRepeatLastN
	}

	// Use the value of the manifest except if the option llm is set
	if modelParams.Name != "" {
		manifest.Model.Name = modelParams.Name
	}
}
