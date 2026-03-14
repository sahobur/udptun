FROM golang:1.26 AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
COPY ./vendor /
RUN CGO_ENABLED=0 GOOS=linux go build -o udptun ./main.go

FROM alpine 
COPY --from=builder /app/udptun /app/

