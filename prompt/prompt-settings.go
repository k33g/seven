package prompt

import (
	"seven/types"
	"strings"
)

// UpdatePromptSettingsWithParams updates the prompt settings in the manifest with the values from the prompt parameters.
//
// Parameters:
// - promptParams: The prompt parameters containing the new values for the prompt settings.
// - manifest: A pointer to the manifest struct where the prompt settings will be updated.
func UpdateSettingsWithParams(promptParams types.PromptParameters, manifest *types.Manifest) {
	// Prompt Settings

	if promptParams.Temperature != 0 {
		manifest.Prompt.Settings.Temperature = promptParams.Temperature
	}

	if promptParams.StopWords != "" {
		manifest.Prompt.Settings.StopWords = strings.Split(promptParams.StopWords, ",")
	}
}
