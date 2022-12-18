package domain

import "crypto/rand"

type RequestID [32]byte

func NewRequestID() (requestID RequestID) {
	_, _ = rand.Read(requestID[:])
	return requestID
}

type PortResponse struct {
	Port      *Port
	RequestID RequestID
}

type PortRequest struct {
	Key       string
	RequestID RequestID
}
