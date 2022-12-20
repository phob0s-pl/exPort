package ports

import "github.com/phob0s-pl/exPort/domain"

// PortStore is interface for underlying database implementation
type PortStore interface {
	PortAdd(port *domain.Port) error
}
