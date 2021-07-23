package controller

import (
	"github.com/iikmaulana/gateway/service"
)

type ServerConfig struct {
	Host      string
	Port      int
	Key       string
	Name      string
	Namespace string
}

func NewServer(cfg ServerConfig, reg *Registry) (*service.Server, error) {
	return service.NewServer(service.Config{
		Host:      cfg.Host,
		Port:      cfg.Port,
		Key:       cfg.Key,
		Name:      cfg.Name,
		Namespace: cfg.Namespace,
	}, reg.writer)
}
