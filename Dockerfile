# Base
ARG GO_VERSION=1.21
FROM golang:${GO_VERSION} as base
WORKDIR /app
COPY . .
RUN go mod download

# Development
FROM base as dev
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
CMD ["air"]

# Build
FROM base as build
RUN go build cmd/server.go

# Production
FROM gcr.io/distroless/static-debian12 as prod
COPY --from=build /app/server /
ENTRYPOINT ["/server"]