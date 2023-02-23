package locale

type Config struct {
	Default string   `ini:"default"`
	BaseDir string   `ini:"dir"`
	Files   []string `ini:"files"`
}
