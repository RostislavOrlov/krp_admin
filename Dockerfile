FROM golang:1.21-alpine AS builder

WORKDIR /go/src/github.com/RostislavOrlov/krp_admin
COPY . .

RUN go build -o ./bin/krp_admin ./cmd/krp_admin

FROM alpine:latest AS runner

COPY --from=builder /go/src/github.com/RostislavOrlov/krp_admin/bin/krp_admin /app/krp_admin

RUN apk -U --no-cache add bash ca-certificates \
    && chmod +x /app/krp_admin

WORKDIR /app
ENTRYPOINT ["/app/krp_admin"]
