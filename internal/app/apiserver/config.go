package apiserver

type Config struct {
	BindAddr       string `toml:"bind_addr"`
	LogLevel       string `toml:"log_level"`
	DatabaseURL    string `toml:"database_url"`
	EmailSender    string `toml:"email_sender"`
	PasswordSender string `toml:"password_sender"`
	SmptPort       int    `toml:"smpt_port"`
}

// Default config values
func NewConfig() *Config {
	return &Config{
		BindAddr:       ":8080",
		LogLevel:       "debug",
		EmailSender:    "",
		PasswordSender: "",
		SmptPort:       456,
	}
}
