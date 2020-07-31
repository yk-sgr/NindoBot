FROM golang:1.14-alpine as builder

RUN apk add --no-cache gcc libc-dev
WORKDIR /opt/server
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/bot/main.go

FROM alpine

WORKDIR /opt/bot
COPY --from=builder /opt/bot/bot .
CMD [ "./bot" ]