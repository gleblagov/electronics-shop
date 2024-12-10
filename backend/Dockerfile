FROM golang:1.22-alpine AS builder

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o backend ./cmd/main.go

FROM alpine AS runner
COPY --from=builder /build/backend .
ENTRYPOINT ["./backend"]
