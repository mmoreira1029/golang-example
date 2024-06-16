FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o otel-go .

FROM scratch
COPY --from=builder /app/otel-go .
CMD ["./otel-go"]