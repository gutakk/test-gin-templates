FROM golang:1.16-buster

WORKDIR /app

# Install Go dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go get github.com/pressly/goose/cmd/goose

EXPOSE 8080
