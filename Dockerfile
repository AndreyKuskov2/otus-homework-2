FROM golang:1.23.3

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/health ./cmd/health/main.go

EXPOSE 8000

ENTRYPOINT ["./bin/health"]