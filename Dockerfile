FROM golang:1.21-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY ./ ./
RUN go mod download

EXPOSE 8080
EXPOSE 2345


CMD ["air", "-c", ".air.toml"]
