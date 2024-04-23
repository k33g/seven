package cli

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Seven struct {
	Config struct {
		Ollama struct {
			URL                string `yaml:"url"`
		} `yaml:"ollama"`
	} `yaml:"config"`
}

func YamlFileToSeven(yamlFile string) (Seven, error) {
	yamlBytes, err := os.ReadFile(yamlFile)
	if err != nil {
		return Seven{}, err
	}
	return YamlStringToSeven(string(yamlBytes))
}

func YamlStringToSeven(yamlString string) (Seven, error) {
	var sevenConfig Seven
	err := yaml.Unmarshal([]byte(yamlString), &sevenConfig)
	if err != nil {
		return Seven{}, err
	}
	return sevenConfig, nil
}
