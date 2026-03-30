FROM golang:1.26 AS builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/nutrition-service

FROM alpine:3
RUN apk add --no-cache ca-certificates
COPY --from=builder /bin/nutrition-service nutrition-service
CMD ["/nutrition-service"]
