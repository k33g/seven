package types


type Parameters struct {
	ManifestPath string // Seven yaml manifest
	EnvfilePath  string // not used
	ConfigPath   string // Configuration file - by default use the SEVENCONFIG environment variable
	OutputPath   string // Completion output file
	ShowLogs     bool
}

type PromptParameters struct {
	Question    string  // Human message
	Context     string  // Context message
	System      string  // System message
	Temperature float64 // Prompt settings
	StopWords   string  // Prompt settings
}

type ModelParameters struct {
	Name        string // Model settings
	PredictRepeatLastN int // Model settings
}
