FROM golang:1.25.5-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY src ./src

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app ./src/cmd/main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 3001

ENTRYPOINT ["./app"]
