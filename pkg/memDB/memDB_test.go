package memDB

import (
	"testing"

	"github.com/phob0s-pl/exPort/domain"
)

func TestPortAdd(t *testing.T) {
	db := NewDatabase()
	if err := db.PortAdd(&domain.Port{
		Key: "test",
	}); err != nil {
		t.Errorf("failed to add port to database: %s", err)
	}

	if l := len(db.store); l != 1 {
		t.Errorf("after adding port expected to have 1 port, got %d", l)
	}
}
