package ports

import "github.com/phob0s-pl/exPort/domain"

type PortProcessPublisher interface {
	PublishPortProcess(fileName string) error
}

type PortStorePublisher interface {
	PublishPortStore(port *domain.Port) error
}
