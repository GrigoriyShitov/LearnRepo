# Используем официальный образ Golang в качестве базового
FROM golang:alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum и загружаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальные файлы в контейнер
COPY . .

# Собираем приложение
RUN go build -o main .

# Используем другой образ для запуска
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем скомпилированный бинарный файл из предыдущего образа
COPY --from=builder /app/main .

# Открываем порт 8080
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]