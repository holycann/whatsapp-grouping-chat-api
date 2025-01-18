# Menggunakan image Go resmi sebagai base image
FROM golang:1.23-bookworm

RUN apt-get update && \
    apt-get install -y --no-install-recommends git curl && \
    rm -rf /var/lib/apt/lists/*

# Install dependencies dan air (alat untuk hot-reloading)
RUN curl -fLo /usr/local/bin/air https://github.com/air-verse/air/releases/download/v1.61.5/air_1.61.5_linux_amd64.tar.gz && \
    chmod +x /usr/local/bin/air

# Set working directory di dalam container
WORKDIR /app

# Menyalin file Go modules (go.mod & go.sum) dan mengunduh dependensi
COPY go.mod go.sum ./
RUN go mod tidy

# Menyalin seluruh file proyek ke dalam container
COPY . .

# Perintah untuk menjalankan air yang akan memantau perubahan file
CMD ["air"]