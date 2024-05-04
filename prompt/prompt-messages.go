package prompt

import (
	"log"
	"seven/tools"
	"seven/types"
)

func UpdateMessagesWithParams(promptParams types.PromptParameters, manifest *types.Manifest) {
	//! Why I added \n at the end?
	if promptParams.Question != "" {
		manifest.Prompt.Human = promptParams.Question + "\n"
	}
	if promptParams.Context != "" {
		manifest.Prompt.Context = promptParams.Context + "\n"
	}
	if promptParams.System != "" {
		manifest.Prompt.System = promptParams.System + "\n"
	}

}

func EnvVarSubstMessages(manifest *types.Manifest) {
	if manifest.Prompt.System != "" {
		var err error
		manifest.Prompt.System, err = tools.EnvSubst(manifest.Prompt.System)
		if err != nil {
			log.Fatal("😡 error with envsubst: ", err)
		}
	}
	if manifest.Prompt.Human != "" {
		var err error
		manifest.Prompt.Human, err = tools.EnvSubst(manifest.Prompt.Human)
		if err != nil {
			log.Fatal("😡 error with envsubst: ", err)
		}
	}
	if manifest.Prompt.Context != "" {
		var err error
		manifest.Prompt.Context, err = tools.EnvSubst(manifest.Prompt.Context)
		if err != nil {
			log.Fatal("😡 error with envsubst: ", err)
		}
	}
}
