FROM golang:1.18-alpine as builder

WORKDIR /app

# Fetch dependencies.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the binary.
RUN cd cmd/web-server-cache \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags='-w -s -extldflags "-static"' -a \
      -o /go/bin/web-server-cache .

FROM alpine:3.16

ARG ROOT_FOLDER="app"
ARG BUILD_ID=0
ARG VERSION=0.0.1

ENV ROOT_FOLDER=${ROOT_FOLDER} \
    BUILD_ID=${BUILD_ID} \
    VERSION=${VERSION}

# Copy static executable
COPY --from=builder /go/bin/web-server-cache app/web-server-cache

COPY public public

EXPOSE 9090

# Run binary.
ENTRYPOINT ["app/web-server-cache"]
