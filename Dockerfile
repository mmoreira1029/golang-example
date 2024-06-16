FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o golang-example .

FROM scratch
COPY --from=builder /app/golang-example .
CMD ["./golang-example"]