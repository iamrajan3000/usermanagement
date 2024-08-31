# Build stage
FROM golang:1.21.6 AS build

WORKDIR /app

COPY . .

RUN go mod tidy
# Build the Go binary as a static executable
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/main .

RUN chmod +x /root/main

EXPOSE 8080

CMD ["./main"]
