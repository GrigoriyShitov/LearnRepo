FROM golang:1.25-alpine AS builder
WORKDIR /src

RUN apk add --no-cache make git curl

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY .. .

ARG MAKE_TARGET=build
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    make ${MAKE_TARGET}

FROM alpine:3.19
RUN apk add --no-cache curl

WORKDIR /internal
COPY --from=builder /src/bin/awesome-backend ./awesome-backend
#COPY --from=builder /src/migrations ./migrations
COPY --from=builder /src/docs ./docs

EXPOSE 8080

CMD ["./awesome-backend"]
