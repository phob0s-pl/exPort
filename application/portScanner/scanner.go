package portScanner

import (
	"github.com/phob0s-pl/exPort/domain"
	"github.com/phob0s-pl/exPort/pkg/portProcessor"
	"github.com/phob0s-pl/exPort/ports"
	log "github.com/sirupsen/logrus"
)

type Application struct {
	publisher ports.PortStorePublisher
}

func NewApplication(publisher ports.PortStorePublisher) *Application {
	return &Application{publisher: publisher}
}

func (a *Application) HandleProcess(msg []byte) {
	log.WithField("service", "scanner").
		WithField("file", string(msg)).
		Debugf("new file scanning request")

	if err := portProcessor.ProcessJSONFile(string(msg), a.OnPort); err != nil {
		log.WithField("service", "scanner").
			WithError(err).
			Errorf("failed to process file")
	}

	log.WithField("service", "scanner").
		WithField("file", string(msg)).
		Debugf("finished processing file")
}

func (a *Application) OnPort(port *domain.Port) {
	if err := a.publisher.PublishPortStore(port); err != nil {
		log.WithField("service", "scanner").
			WithError(err).
			Errorf("failed to publish port store")
	}
}
