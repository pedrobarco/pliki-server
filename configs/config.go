package configs

type Config struct {
	Port           int    `yaml:"port"`
	FileServerRoot string `yaml:"fileServerRoot"`
	MaxFileSizeMb  int64  `yaml:"maxFileSizeMb"`
}

func GetDefaultConfig() *Config {
	return &Config{
		Port:           8080,
		FileServerRoot: "./files/",
		MaxFileSizeMb:  10,
	}
}
