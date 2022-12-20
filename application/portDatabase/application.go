package portDatabase

import (
	"bytes"
	"encoding/gob"

	"github.com/phob0s-pl/exPort/domain"
	"github.com/phob0s-pl/exPort/ports"
	log "github.com/sirupsen/logrus"
)

// Application is database application which stores port to underlying database when message handler is invoked
type Application struct {
	database ports.PortStore
}

func NewApplication(database ports.PortStore) *Application {
	return &Application{
		database: database,
	}
}

func (a *Application) HandleStoreMessage(data []byte) {
	var port domain.Port

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
