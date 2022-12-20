package portDatabase

import (
	"bytes"
	"encoding/gob"
	"testing"

	"github.com/phob0s-pl/exPort/domain"
)

type DbMock struct {
	counter int
}

func (d *DbMock) PortAdd(*domain.Port) error {
	d.counter++
	return nil
}

func addDummyPort(t *testing.T, port *domain.Port, fn func([]byte)) {
	var (
		buff    bytes.Buffer
		encoder = gob.NewEncoder(&buff)
	)

	if err := encoder.Encode(port); err != nil {
		t.Fatalf("failed to encode port: %s", err)
	}

	fn(buff.Bytes())
}

func TestApplication(t *testing.T) {
	var (
		dbMock      = &DbMock{}
		application = NewApplication(dbMock)
	)
	addDummyPort(t, &domain.Port{Key: "test"}, application.HandleStoreMessage)

	if dbMock.counter != 1 {
		t.Errorf("expected database mock to be invoked on HandleStoreMessage")
	}
}
