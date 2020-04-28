package server

type Config struct {
	BindAddr string
	LogLvl   string
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLvl:   "trace",
	}
}
