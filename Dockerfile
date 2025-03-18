#syntax=docker/dockerfile:1

FROM --platform=$BUILDPLATFORM tonistiigi/xx:1.6.1 AS xx

FROM --platform=$BUILDPLATFORM golang:1.24.0-alpine AS builder
WORKDIR /app

COPY --from=xx / /

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# Set Golang build envs based on Docker platform string
ARG TARGETPLATFORM
RUN CGO_ENABLED=0 xx-go build -ldflags='-w -s' -trimpath -o print-xterm256


FROM scratch
COPY --from=builder /app/print-xterm256 /
ENTRYPOINT ["/print-xterm256"]
