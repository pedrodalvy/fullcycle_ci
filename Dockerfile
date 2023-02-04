FROM golang:1.20-alpine3.17

WORKDIR /app

COPY . .
RUN go build -o math math.go

CMD ["./math"]
