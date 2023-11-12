FROM golang:1.17 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o exporter .

FROM alpine:latest
WORKDIR /
COPY --from=builder /app/exporter .
CMD ["./exporter"]