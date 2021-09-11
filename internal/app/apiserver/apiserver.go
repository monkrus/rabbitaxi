package apiserver

import "github.com/sirupsen/logrus"

//APIserver
type APIserver struct {
	config *Config
	logger *logrus.Logger
}

//New
func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
	}
}

func (s *APIserver) Start() error {
	return nil
}

func (s *APIserver) configLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}
