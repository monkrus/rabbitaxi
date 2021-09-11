package apiserver

//APIserver
type APIserver struct {
	config *Config
}

//New
func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
	}
}

func (s *APIserver) Start() error {
	return nil
}
