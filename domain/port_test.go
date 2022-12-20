package domain

import "testing"

func TestPortValid(t *testing.T) {
	var (
		invalidPort Port
		validPort   = Port{Key: "123"}
	)

	if invalidPort.IsValid() {
		t.Errorf("expected port with no key to be invalid")
	}

	if !validPort.IsValid() {
		t.Errorf("expected port with key to be valid")
	}
}
