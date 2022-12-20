package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/phob0s-pl/exPort/application/portDatabase"
	"github.com/phob0s-pl/exPort/domain"
	"github.com/phob0s-pl/exPort/pkg/broker"
	"github.com/phob0s-pl/exPort/pkg/logger"
	"github.com/phob0s-pl/exPort/pkg/memDB"
	log "github.com/sirupsen/logrus"
)

func main() {
	signalC := make(chan os.Signal, 2)

	logger.Configure()
	log.WithField("service", "database").
		Infof("service starting")

	signal.Notify(signalC, syscall.SIGTERM, syscall.SIGINT)

	application := portDatabase.NewApplication(memDB.NewDatabase())

	storeConsumer, err := broker.NewConsumer(domain.TopicPortStore, application.HandleStoreMessage)
	if err != nil {
		log.WithField("service", "database").
			WithError(err).Fatalf("failed to create consumer")
	}

	sig := <-signalC
	log.WithField("service", "database").
		WithField("signal", sig.String()).
		Infof("stopping service")
	storeConsumer.Stop()
	os.Exit(0)
}
