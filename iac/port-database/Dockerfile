############################
# Build
############################
FROM golang:buster AS builder

WORKDIR /build
COPY . .

RUN make port-database

############################
# Actual image
############################
FROM ubuntu:jammy

COPY --from=builder /build/port-database /bin/port-database

CMD ["/bin/port-database"]
