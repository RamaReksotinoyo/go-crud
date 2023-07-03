# Stage pertama: Build aplikasi
FROM golang:1.19-alpine AS build

# Membuat direktori kerja
WORKDIR /app

# Menyalin file main.go dan go.mod ke dalam direktori kerja
COPY main.go /app
COPY go.mod /app

# Mengambil dependensi Go
RUN go mod download

# Menambahkan entri go.sum untuk dependensi github.com/labstack/echo/v4
RUN go get github.com/labstack/echo/v4@v4.5.0

# Mengkompilasi program Go menjadi binary
RUN go build -o main .

# Stage kedua: Menjalankan aplikasi
FROM alpine:latest

# Membuat direktori kerja
WORKDIR /app

# Menyalin binary hasil kompilasi ke dalam image
COPY --from=build /app/main /app/main

# Menjalankan program ketika container dijalankan
CMD ["./main"]