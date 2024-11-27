## base go image
FROM golang:1.23-alpine as builder

RUN apk update && \
    apk add build-base


RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    swag init --parseDependency

RUN CGO_ENABLED=0 go build -o ggocamping .

RUN chmod +x /app/ggocamping

## build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/ggocamping /app
COPY --from=builder /app/docs /app/docs

WORKDIR /app

CMD [ "./ggocamping"]