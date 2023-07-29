ARG GO_VERSION=1.18
FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-w -s" -o main ./cmd/api/

FROM alpine

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 1323

CMD ["./main"]