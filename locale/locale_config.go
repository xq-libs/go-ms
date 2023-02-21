package locale

type LocaleConfig struct {
	Default string   `toml:"default"`
	BaseDir string   `toml:"dir"`
	Files   []string `toml:"files"`
}
