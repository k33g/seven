package cli

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
