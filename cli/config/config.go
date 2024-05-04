package config

import (
	"log"
	"os"
	"seven/types"
)

// GetManifestAndConfig reads the yaml config file and yaml manifest file, and returns the parsed manifest and seven objects.
//
// It takes in a parameters object of type types.Parameters, which contains the paths to the config and manifest files.
// The function returns a tuple of types.Manifest and types.Seven, representing the parsed manifest and seven objects respectively.
func GetManifestAndConfig(params types.Parameters) (types.Manifest, types.Seven) {
	// Read the yaml config file
	seven, err := YamlFileToSeven(params.ConfigPath)
	if err != nil {
		log.Fatal("😡 error loading config file: ", err)
	}

	// Ollama Settings
	ollamaURL := seven.Config.Ollama.URL
	if ollamaURL == "" {
		ollamaURL = "http://localhost:11434"
		//ollamaURL = "http://host.docker.internal:11434"
	}
	// Predefined variable:
	os.Setenv("SEVEN_OLLAMA_URL", ollamaURL)

	// Read the yaml manifest file
	manifest, err := YamlFileToManifest(params.ManifestPath)
	if err != nil {
		log.Fatal("😡 error loading .yaml file: ", err)
	}
	// Predefined variable:
	os.Setenv("SEVEN_MODEL_NAME", manifest.Model.Name)

	return manifest, seven
}
