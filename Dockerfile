#syntax=docker/dockerfile:1

FROM --platform=$BUILDPLATFORM golang:1.24.0-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# Set Golang build envs based on Docker platform string
ARG TARGETPLATFORM
RUN <<EOT
  set -eux
  case "$TARGETPLATFORM" in
    'linux/amd64') export GOARCH=amd64 ;;
    'linux/arm/v6') export GOARCH=arm GOARM=6 ;;
    'linux/arm/v7') export GOARCH=arm GOARM=7 ;;
    'linux/arm64') export GOARCH=arm64 ;;
    *) echo "Unsupported target: $TARGETPLATFORM" && exit 1 ;;
  esac
  go build -ldflags='-w -s' -trimpath -o print-xterm256
EOT

FROM scratch
COPY --from=builder /app/print-xterm256 /
ENTRYPOINT ["/print-xterm256"]
