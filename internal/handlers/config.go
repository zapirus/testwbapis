package handlers

type Config struct {
	HTTPAddr string `toml:"bind_addr"`
}

func NewConfig() *Config {
	return &Config{
		HTTPAddr: ":8080",
	}

}
