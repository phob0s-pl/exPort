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

.PHONY: docker-api-gateway
docker-api-gateway:
	docker build -t export-api-gateway:latest --file iac/api-gateway/Dockerfile .

.PHONY: docker-port-database
docker-port-database:
	docker build -t export-port-database:latest --file iac/port-database/Dockerfile .

.PHONY: docker-port-scanner
docker-port-scanner:
	docker build -t export-port-scanner:latest --file iac/port-scanner/Dockerfile .
