FROM golang:latest

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o ./examples/experiment/main ./examples/experiment/main.go

WORKDIR /app/examples/experiment

CMD ["./main"]