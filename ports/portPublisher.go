package ports

import "github.com/phob0s-pl/exPort/domain"

type PortTaskPublisher interface {
	PublishPortProcess(fileName string) error
	PublishPortGetRequest(key string, requestID domain.RequestID) error
}

type PortResponsePublisher interface {
	PublishPortGetResponse(requestID domain.RequestID, port *domain.Port) error
}

type PortStorePublisher interface {
	PublishPortStore(port *domain.Port) error
}
