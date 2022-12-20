package memDB

import (
	"sync"

	"github.com/phob0s-pl/exPort/domain"
)

type Database struct {
	sync.RWMutex
	store map[string]*domain.Port
}

func NewDatabase() *Database {
	return &Database{
		store: map[string]*domain.Port{},
	}
}

func (d *Database) PortAdd(port *domain.Port) error {
	d.Lock()
	defer d.Unlock()

	d.store[port.Key] = port

	return nil
}
