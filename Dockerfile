FROM golang:1.21.0-bullseye
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest
COPY . /app
RUN swag init --dir cmd
EXPOSE 8001
CMD go run cmd/main.go
