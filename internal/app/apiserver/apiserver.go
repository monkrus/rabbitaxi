package apiserver

import (
	"github.com/sirupsen/logrus"
)

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

// Start ...
func (s *APIserver) Start() error {
	if err := s.configLogger(); err != nil {
		return err
	}
	s.logger.Info("starting api server")
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
