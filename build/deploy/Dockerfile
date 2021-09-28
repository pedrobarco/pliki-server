FROM golang:alpine as builder

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o bin/pliki-server ./cmd/pliki-server

FROM alpine

WORKDIR /app

COPY --from=builder /app/bin/pliki-server .

CMD ["./pliki-server"]
