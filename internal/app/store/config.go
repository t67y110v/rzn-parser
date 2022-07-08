package store

type Config struct {
	DataBaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}
