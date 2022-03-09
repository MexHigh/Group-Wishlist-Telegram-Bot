# Build Go backend
FROM golang:1.17 AS go-builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"' -tags timetzdata


# Collect builds in scratch image
FROM scratch
LABEL maintainer="Leon Schmidt"

# Copy CA certs and timezone info
COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=go-builder /usr/share/zoneinfo /usr/share/zoneinfo
# Copy compiled binary
COPY --from=go-builder /go/bin/group-wishlist-telegram-bot /app/group-wishlist-telegram-bot
# Copy example config
COPY ./config.example.json /app/config.json

WORKDIR /app
ENTRYPOINT ["/app/group-wishlist-telegram-bot"]
CMD []