## base go image
FROM golang:1.23-alpine as builder

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=1 go build -o ggocamping .
RUN chmod +x /app/ggocamping

## build a tiny docker image
FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/ggocamping /app

CMD [ "/app/ggocamping"]