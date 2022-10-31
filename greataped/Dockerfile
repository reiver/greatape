FROM golang:1.19 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' -o ./bin/greatape .

FROM scratch

ENV PROTOCOL="http"
ENV DOMAIN="localhost"
ENV PORT=80

COPY --from=builder /app/bin /app

EXPOSE $PORT

ENTRYPOINT ["/app/greatape"]