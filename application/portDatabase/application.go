package portDatabase

import (
	"bytes"
	"encoding/gob"
	"github.com/phob0s-pl/exPort/domain"
	"github.com/phob0s-pl/exPort/ports"
	log "github.com/sirupsen/logrus"
)

type Application struct {
	database  ports.PortStore
	publisher ports.PortResponsePublisher
}

func NewApplication(database ports.PortStore, publisher ports.PortResponsePublisher) *Application {
	return &Application{
		database:  database,
		publisher: publisher,
	}
}

func (a *Application) HandleStoreMessage(data []byte) {
	var (
		port domain.Port
	)

	log.WithField("service", "database").
		Debugf("new store request received")

	if err := gob.NewDecoder(bytes.NewBuffer(data)).Decode(&port); err != nil {
		log.WithField("service", "database").
			WithError(err).
			Warnf("failed to decode port")
		return
	}

	if err := a.database.PortAdd(&port); err != nil {
		log.WithField("service", "database").
			WithError(err).
			Warnf("failed to add port")
		return
	}

	log.WithField("service", "database").
		WithField("key", port.Key).
		Debugf("port successfully stored")
}

func (a *Application) HandleGetMessage(data []byte) {
	var (
		request domain.PortRequest
	)

	if err := gob.NewDecoder(bytes.NewBuffer(data)).Decode(&request); err != nil {
		log.WithField("service", "database").
			WithError(err).
			Warnf("failed to decode request")
		return
	}

	port, err := a.database.PortGet(request.Key)
	if err != nil {
		if err := a.publisher.PublishPortGetResponse(request.RequestID, nil); err != nil {
			log.WithField("service", "database").
				WithError(err).
				Warnf("failed to publish response")
		}
		return
	}

	if err := a.publisher.PublishPortGetResponse(request.RequestID, port); err != nil {
		log.WithField("service", "database").
			WithError(err).
			Warnf("failed to publish response")
	}
}
