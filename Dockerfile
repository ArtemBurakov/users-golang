# Build stage
FROM golang:latest AS builder
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/main.go

# Runtime stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/src/app/server /app/

ARG PORT=8045
ARG GIN_MODE=release

ENV PORT $PORT
ENV GIN_MODE $GIN_MODE

EXPOSE $PORT

CMD ["/app/server"]
