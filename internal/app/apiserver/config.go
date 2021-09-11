package apiserver

//Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
}

//NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
