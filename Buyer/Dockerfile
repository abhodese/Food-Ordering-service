FROM golang:1.17
WORKDIR /app
COPY buyer_service.go .
RUN go build -o buyer_service .
EXPOSE 8080
CMD ["./buyer_service"]