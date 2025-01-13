# Menggunakan image Go resmi sebagai base image
FROM golang:1.23-alpine AS builder

# Set working directory di dalam container
WORKDIR /app

# Menyalin file Go modules (go.mod & go.sum) dan mengunduh dependensi
COPY go.mod go.sum ./
RUN go mod tidy

# Menyalin seluruh file proyek ke dalam container
COPY . .

# Membuat aplikasi Go
RUN go build -o app .

# Menyusun image final untuk aplikasi Go
FROM alpine:latest

# Menyalin file aplikasi dari stage sebelumnya
WORKDIR /root/
COPY --from=builder /app/app .

# Set perintah untuk menjalankan aplikasi
CMD ["./app"]