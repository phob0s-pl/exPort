package memDB

import (
	"fmt"
	"github.com/phob0s-pl/exPort/domain"
	"sync"
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

func (d *Database) PortGet(key string) (*domain.Port, error) {
	d.RLock()
	defer d.RUnlock()

	port, exist := d.store[key]
	if !exist {
		return nil, fmt.Errorf("not found")
	}

	return port, nil
}
