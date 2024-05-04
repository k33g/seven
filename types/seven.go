package types

type Seven struct {
	Config struct {
		Ollama struct {
			URL string `yaml:"url"`
		} `yaml:"ollama"`
	} `yaml:"config"`
}
