# Menggunakan image Go resmi sebagai base image
FROM golang:1.23-bookworm

# Install dependencies dan air (alat untuk hot-reloading)
RUN apt-get update && \
    apt-get install -y --no-install-recommends git curl && \
    rm -rf /var/lib/apt/lists/*

# Set working directory di dalam container
WORKDIR /app

# Menyalin file Go modules (go.mod & go.sum) dan mengunduh dependensi
COPY go.mod go.sum ./
RUN go mod tidy

# Menyalin seluruh file proyek ke dalam container
COPY . .

# Perintah untuk menjalankan air yang akan memantau perubahan file
CMD ["air"]