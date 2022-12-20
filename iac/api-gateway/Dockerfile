############################
# Build
############################
FROM golang:buster AS builder

WORKDIR /build
COPY . .

RUN make api-gateway

############################
# Actual image
############################
FROM ubuntu:jammy

COPY --from=builder /build/api-gateway /bin/api-gateway

CMD ["/bin/api-gateway"]
