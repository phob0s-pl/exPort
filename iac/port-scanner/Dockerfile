############################
# Build
############################
FROM golang:buster AS builder

WORKDIR /build
COPY . .

RUN make port-scanner

############################
# Actual image
############################
FROM ubuntu:jammy

COPY --from=builder /build/port-scanner /bin/port-scanner
COPY iac/port-scanner/ports.json /ports.json

WORKDIR /

CMD ["/bin/port-scanner"]
