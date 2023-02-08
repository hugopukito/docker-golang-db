FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

FROM alpine:latest
RUN apk update && apk add netcat-openbsd
WORKDIR /app
COPY --from=builder /app/main .
COPY wait-for.sh .

CMD [ "/app/main" ]