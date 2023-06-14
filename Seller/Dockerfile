FROM golang:1.17
WORKDIR /app
COPY seller_service.go .
RUN go build -o seller_service .
EXPOSE 8080
CMD ["./seller_service"]