FROM golang:1.25-alpine AS builder
LABEL authors="Григорий"

WORKDIR /src

COPY go.mod .

RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY ../.. .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /src/app ./app

FROM alpine:3.19 AS curl-layer
RUN apk update && apk add --no-cache curl


FROM alpine/curl AS runtime

WORKDIR /app
COPY --from=builder /src/app ./app

USER 10001
CMD ["./app/app"]


