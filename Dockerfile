# Sử dụng Golang base image
FROM golang:1.23.2-alpine

# Cài đặt các package cần thiết
RUN apk update && apk add --no-cache git

# Đặt thư mục làm việc bên trong container
WORKDIR /app

# Copy tất cả file từ thư mục hiện tại vào thư mục làm việc trong container
COPY . .

# Download module dependencies
RUN go mod tidy

# Build app thành một binary executable
RUN go build -o ginjwt1-app .

# Expose cổng mà ứng dụng sử dụng
EXPOSE 5000

# Command to run the executable
CMD ["./ginjwt1-app"]
