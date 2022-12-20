package apiGateway

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phob0s-pl/exPort/ports"
	log "github.com/sirupsen/logrus"
)

type Application struct {
	srv       *http.Server
	publisher ports.PortProcessPublisher
}

func NewApplication(publisher ports.PortProcessPublisher) *Application {
	var (
		app = &Application{
			publisher: publisher,
		}
		router = mux.NewRouter()
	)

	router.HandleFunc("/port/process-file/{filename}", app.portProcess)

	app.srv = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return app
}

func (a *Application) Stop() {
	if err := a.srv.Close(); err != nil {
		log.WithField("service", "api-gateway").
			WithError(err).
			Warnf("failed to stop listener")
	}
}

func (a *Application) Listen() error {
	return a.srv.ListenAndServe()
}

func (a *Application) portProcess(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	filename, ok := vars["filename"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	log.WithField("service", "api-gateway").
		WithField("url", request.URL).
		Debugf("new request")

	if err := a.publisher.PublishPortProcess(filename); err != nil {
		log.WithField("service", "api-gateway").
			WithField("url", request.URL).
			WithError(err).
			Warnf("failed to publish processing request")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
