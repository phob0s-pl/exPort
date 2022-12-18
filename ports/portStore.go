package ports

import "github.com/phob0s-pl/exPort/domain"

type PortStore interface {
	PortAdd(port *domain.Port) error
	PortGet(key string) (*domain.Port, error)
}
