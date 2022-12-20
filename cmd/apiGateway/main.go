package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/phob0s-pl/exPort/application/apiGateway"
	"github.com/phob0s-pl/exPort/pkg/broker"
	"github.com/phob0s-pl/exPort/pkg/logger"
	log "github.com/sirupsen/logrus"
)

func main() {
	var (
		signalC   = make(chan os.Signal)
		listenerC = make(chan error)
	)

	logger.Configure()
	log.WithField("service", "api-gateway").
		Infof("service starting")

	signal.Notify(signalC, syscall.SIGTERM, syscall.SIGINT)

	publisher, err := broker.NewPublisher()
	if err != nil {
		log.WithField("service", "api-gateway").
			WithError(err).Fatalf("failed to create publisher")
	}

	application := apiGateway.NewApplication(publisher)

	go func() {
		listenerC <- application.Listen()
	}()

	select {
	case sig := <-signalC:
		log.WithField("service", "api-gateway").
			WithField("signal", sig.String()).
			Infof("stopping service")
		application.Stop()
		publisher.Stop()
		os.Exit(0)
	case err := <-listenerC:
		log.WithField("service", "api-gateway").
			WithError(err).
			Errorf("lister failed")
		os.Exit(1)
	}
}
