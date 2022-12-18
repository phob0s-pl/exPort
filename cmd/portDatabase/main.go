package main

import (
	"github.com/phob0s-pl/exPort/application/portDatabase"
	"github.com/phob0s-pl/exPort/domain"
	"github.com/phob0s-pl/exPort/pkg/broker"
	"github.com/phob0s-pl/exPort/pkg/logger"
	"github.com/phob0s-pl/exPort/pkg/memDB"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		signalC = make(chan os.Signal)
	)

	logger.Configure()
	log.WithField("service", "database").
		Infof("service starting")

	signal.Notify(signalC, syscall.SIGTERM, syscall.SIGINT)

	publisher, err := broker.NewPublisher()
	if err != nil {
		log.WithField("service", "database").
			WithError(err).Fatalf("failed to create publisher")
	}

	application := portDatabase.NewApplication(memDB.NewDatabase(), publisher)

	getConsumer, err := broker.NewConsumer(domain.TopicPortGetRequest, application.HandleGetMessage)
	if err != nil {
		log.WithField("service", "database").
			WithError(err).Fatalf("failed to create consumer")
	}

	storeConsumer, err := broker.NewConsumer(domain.TopicPortStore, application.HandleStoreMessage)
	if err != nil {
		log.WithField("service", "database").
			WithError(err).Fatalf("failed to create consumer")
	}

	select {
	case sig := <-signalC:
		log.WithField("service", "database").
			WithField("signal", sig.String()).
			Infof("stopping service")
		storeConsumer.Stop()
		getConsumer.Stop()
		publisher.Stop()
		os.Exit(0)
	}
}
