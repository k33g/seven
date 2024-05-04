package logs

import (
	"fmt"
	"seven/types"
	
)

func ShowModelPromptInfo(params types.Parameters, manifest types.Manifest, seven types.Seven) {
	if params.ShowLogs {
		fmt.Println("🧠 model:", manifest.Model.Name)
		fmt.Println("🌍 URL:", seven.Config.Ollama.URL)
		fmt.Println("🤭 predictRepeatLastN:", manifest.Model.Settings.PredictRepeatLastN)
		fmt.Println("🌡️  temperature:", manifest.Prompt.Settings.Temperature)
		fmt.Println("🚫 stop words:", manifest.Prompt.Settings.StopWords)
		fmt.Println()
	}
}

func ShowMemoryUsageInfo(params types.Parameters, manifest types.Manifest) {
	if params.ShowLogs {
		fmt.Println("🧠 model:", manifest.Model.Name)
		if manifest.Model.Memory {
			fmt.Println("🧠 memory:", manifest.Model.MemoryStore)
		}
	}
}

func ShowIfWeAreUsingFunctions(params types.Parameters, isItForFunctions bool) {
	if params.ShowLogs {
		if isItForFunctions {
			fmt.Println("✋ using functions")
		}
	}
}
