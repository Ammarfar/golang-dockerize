# Builder
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY . ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o mezink cmd/api.go
RUN CGO_ENABLED=0 GOOS=linux go build -o seeder pkg/database/gorm/seeder/run.go

# Deploy the application binary into a lean image
FROM alpine:edge

RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Jakarta

WORKDIR /

COPY --from=builder /app/mezink mezink
COPY --from=builder /app/seeder seeder

ENTRYPOINT ["/mezink"]