FROM golang:1.20.5-bullseye

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /app/main

CMD ["/app/main"]
