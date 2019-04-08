# Part 1 (Builder)
FROM golang:latest as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main

# Part 2 (Final)
FROM alpine:3.7
WORKDIR /usr/src/app
COPY --from=builder /app/main .
EXPOSE 80
CMD ["./main"]
