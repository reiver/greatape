FROM golang:1.19 AS builder

WORKDIR /src

COPY . .

RUN go mod download

RUN CGO_ENABLED=1 GOOS=linux go build \
    -ldflags '-linkmode external -extldflags "-static" -X github.com/reiver/greatape/components/core.runningInContainer=true -w -s' \
    -a -o ./bin/greatape .

FROM scratch

# development, staging, production
ENV ENVIRONMENT=development

ENV PROTOCOL=https
ENV FQDN=yourdomain.com
ENV PORT=7080

# jwt_token_key returns the jwt token secret
ENV JWT_TOKEN_KEY=INSERT_YOUR_KEY
# jwt_token_expiration returns the jwt token expiration duration.
# Should be time.ParseDuration string.
# Source: https://golang.org/pkg/time/#ParseDuration
# default: 10h
ENV JWT_TOKEN_EXP=10h

ENV POSTGRES_HOST=127.0.0.1
ENV POSTGRES_PORT=5432
ENV POSTGRES_DATABASE=greatape
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=password

COPY --from=builder /src/bin /app

EXPOSE $PORT

ENTRYPOINT ["/app/greatape"]