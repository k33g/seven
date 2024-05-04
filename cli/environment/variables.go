package environment

import (
	"os"
	"seven/types"
	"strings"
)

func ReadVariablesSection(manifest types.Manifest) {
	// Read the variables section of the yaml manifest file
	for _, varItem := range manifest.Variables {
		nameAndValue := strings.SplitN(varItem, "=", 2)
		os.Setenv(nameAndValue[0], nameAndValue[1])
	}
}
