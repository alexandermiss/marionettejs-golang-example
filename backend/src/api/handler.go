package api 

import (
	"github.com/sirupsen/logrus"
	"linguanski/src/service"
)

var internalErrorMessage = NewErrorMessage("internal server error")

type Server struct {
    providerManager  service.ProviderManager
	clientManager    service.ClientManager
	logger           *logrus.Logger
}

func NewServer(providerManager service.ProviderManager, clientManager service.ClientManager, logger *logrus.Logger) *Server {
	return &Server{
        providerManager: providerManager,
		clientManager: clientManager,
		logger:        logger,
	}
}
