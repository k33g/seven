package config

import (
	"os"
	"seven/types"

	"gopkg.in/yaml.v3"
)

/*
Tips: yaml to structure: https://zhwt.github.io/yaml-to-go/
*/

func YamlFileToManifest(yamlFile string) (types.Manifest, error) {
	yamlBytes, err := os.ReadFile(yamlFile)
	if err != nil {
		return types.Manifest{}, err
	}
	return YamlStringToManifest(string(yamlBytes))
}

func YamlStringToManifest(yamlString string) (types.Manifest, error) {
	var manifest types.Manifest
	err := yaml.Unmarshal([]byte(yamlString), &manifest)
	if err != nil {
		return types.Manifest{}, err
	}
	return manifest, nil
}

func YamlFileToSeven(yamlFile string) (types.Seven, error) {
	yamlBytes, err := os.ReadFile(yamlFile)
	if err != nil {
		return types.Seven{}, err
	}
	return YamlStringToSeven(string(yamlBytes))
}

func YamlStringToSeven(yamlString string) (types.Seven, error) {
	var sevenConfig types.Seven
	err := yaml.Unmarshal([]byte(yamlString), &sevenConfig)
	if err != nil {
		return types.Seven{}, err
	}
	return sevenConfig, nil
}
