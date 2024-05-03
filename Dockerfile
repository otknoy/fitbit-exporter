FROM golang:1.22.2 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o fitbit-exporter main.go

FROM scratch
COPY --from=builder /app/fitbit-exporter /bin/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/bin/fitbit-exporter"]
