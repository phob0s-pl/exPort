.PHONY: all

all: api-gateway port-database port-scanner

.PHONY: api-gateway
api-gateway:
	go build -o api-gateway ./cmd/apiGateway

.PHONY: port-database
port-database:
	go build -o port-database ./cmd/portDatabase

.PHONY: port-scanner
port-scanner:
	go build -o port-scanner ./cmd/portScanner