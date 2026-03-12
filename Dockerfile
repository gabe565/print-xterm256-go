#syntax=docker/dockerfile:1

FROM --platform=$BUILDPLATFORM golang:1.26.1-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ARG TARGETOS
ARG TARGETARCH
RUN --mount=type=cache,target=/root/.cache \
  CGO_ENABLED=0 GOOS="$TARGETOS" GOARCH="$TARGETARCH" \
  go build -ldflags='-w -s' -trimpath -o print-xterm256


FROM scratch
COPY --from=builder /app/print-xterm256 /
ENTRYPOINT ["/print-xterm256"]
