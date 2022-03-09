# Build Go backend
FROM golang:1.17 AS go-builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=1 GOOS=linux go install -a -ldflags '-linkmode external -extldflags "-static"' .

# Collect builds in scratch image
FROM scratch
LABEL maintainer="Leon Schmidt"

# Copy CA certs and timezone info
COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=go-builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=go-builder /go/bin/group-wishlist-telegram-bot /app/group-wishlist-telegram-bot
# Copy example config
COPY ./config.example.json /app/config.json

#VOLUME /app/db

WORKDIR /app
ENTRYPOINT ["/app/group-wishlist-telegram-bot"]
CMD []