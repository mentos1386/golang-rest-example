# Base
ARG GO_VERSION=1.21
FROM golang:${GO_VERSION} as build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /server cmd/server.go

# Get curl
FROM tarampampam/curl AS curl

# Production
FROM gcr.io/distroless/static-debian12 as release

ENV PORT=1234
EXPOSE 1234

COPY --from=curl /bin/curl /bin/curl
HEALTHCHECK --interval=5s --timeout=2s --retries=2 --start-period=2s CMD [ \
    "curl", "--fail", "http://127.0.0.1:1234/healthz" \
]

COPY --from=build /server /server

ENV DATABASE_MIGRATIONS=file:///app/migrations
COPY --from=build /app/migrations /app/migrations

USER nonroot:nonroot
ENTRYPOINT ["/server"]
