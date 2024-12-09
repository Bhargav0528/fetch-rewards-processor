FROM golang:1.23.4-alpine3.21

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

WORKDIR ./rewards-processor

RUN go build -o ./fetch-rewards-api

EXPOSE 8080

CMD ["./fetch-rewards-api"]


