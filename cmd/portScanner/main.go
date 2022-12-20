package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/phob0s-pl/exPort/application/portScanner"
	"github.com/phob0s-pl/exPort/domain"
	"github.com/phob0s-pl/exPort/pkg/broker"
	"github.com/phob0s-pl/exPort/pkg/logger"
	log "github.com/sirupsen/logrus"
)

func main() {
	signalC := make(chan os.Signal, 2)

	logger.Configure()
	log.WithField("service", "scanner").
		Infof("service starting")

	signal.Notify(signalC, syscall.SIGTERM, syscall.SIGINT)

	publisher, err := broker.NewPublisher()
	if err != nil {
		log.WithField("service", "scanner").
			WithError(err).Fatalf("failed to create publisher")
	}

	application := portScanner.NewApplication(publisher)

	scannerConsumer, err := broker.NewConsumer(domain.TopicPortProcess, application.HandleProcess)
	if err != nil {
		log.WithField("service", "scanner").
			WithError(err).Fatalf("failed to create consumer")
	}

	sig := <-signalC
	log.WithField("service", "scanner").
		WithField("signal", sig.String()).
		Infof("stopping service")
	scannerConsumer.Stop()
	publisher.Stop()
	os.Exit(0)
}
