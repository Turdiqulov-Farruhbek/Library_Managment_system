FROM golang:1.22.1 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/myapp .
EXPOSE 9090
CMD ["./myapp"]





# # Base image sifatida Go'ni tanlab olish
# FROM golang:1.22.1 AS builder

# COPY  . /app

# # Ishchi katalog orqali ishlash uchun work directoryni belgilash
# WORKDIR /app

# # Modullarni va dependency-larni yuklash
# COPY go.mod go.sum ./
# RUN go mod download

# # Barcha loyihani katalogga ko'chirish
# COPY . .

# # Loyiha kodini build qilish
# RUN go build -o library cmd/main.go

# # Minimalistic run-time image yaratish
# FROM alpine:latest

# # Ishchi katalog orqali ishlash uchun work directoryni belgilash
# WORKDIR /root/

# # Builder layihadan olingan binary faylini kopiya qilish
# COPY --from=builder /app/library .
# COPY .env .env

# # Tarmoq portini ochish
# EXPOSE 9090

# # Ilovani ishga tushirish
# CMD ["./library"]
