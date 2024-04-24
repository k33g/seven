package types

type PromptComposition int64

const (
	SystemOnly PromptComposition = iota
	ContextOnly
	HumanOnly

	SystemAndContext
	SystemAndHuman
	ContextAndHuman

	SystemAndContextAndHuman

	Functions

	NoComposition
)
// https://zhwt.github.io/yaml-to-go/
type Manifest struct {
	Title         string `yaml:"title"`
	Documentation string `yaml:"documentation"`
	BeforeScript  string `yaml:"before-script"`
	AfterScript   string `yaml:"after-script"`
	Model struct {
		Name        string `yaml:"name"`
		Memory      bool   `yaml:"memory"`
		MemoryStore string `yaml:"memory-store"`
		Stream      bool   `yaml:"stream"`
		Settings struct {
			PredictRepeatLastN int `yaml:"predictRepeatLastN"`
		} `yaml:"settings"`
	} `yaml:"model"`
	Prompt struct {
		// https://github.com/tmc/langchaingo/blob/main/llms/ollama/internal/ollamaclient/types.go#L151
		Settings struct {
			Temperature   float64 `yaml:"temperature"`
			StopWords     []string `yaml:"stopWords"` // does it works with Ollama?
		} `yaml:"settings"`
		System  string `yaml:"system"`
		Context string `yaml:"context"`
		Human   string `yaml:"human"`
	} `yaml:"prompt"`
	Documents []struct { // only used with context
		Title   string `yaml:"title"`
		Content string `yaml:"content"`
	} `yaml:"documents"`	
	Functions []struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Parameters  string `yaml:"parameters"`
		Cmd         string `yaml:"cmd"`
	} `yaml:"functions"`
}


func (m Manifest) PromptComposedBy() PromptComposition {

	if m.Functions != nil {
		return Functions
	}

	if m.Prompt.System != "" && m.Prompt.Context == "" && m.Prompt.Human == "" {
		return SystemOnly
	}
	if m.Prompt.System == "" && m.Prompt.Context != "" && m.Prompt.Human == "" {
		return ContextOnly
	}
	if m.Prompt.System == "" && m.Prompt.Context == "" && m.Prompt.Human != "" {
		return HumanOnly
	}

	if m.Prompt.System != "" && m.Prompt.Context != "" && m.Prompt.Human == "" {
		return SystemAndContext
	}
	if m.Prompt.System != "" && m.Prompt.Context == "" && m.Prompt.Human != "" {
		return SystemAndHuman
	}
	if m.Prompt.System == "" && m.Prompt.Context != "" && m.Prompt.Human != "" {
		return ContextAndHuman
	}

	if m.Prompt.System != "" && m.Prompt.Context != "" && m.Prompt.Human != "" {
		return SystemAndContextAndHuman
	}

	return NoComposition
}
