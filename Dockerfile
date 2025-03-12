# Используем официальный образ Golang
FROM golang:1.22.2-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Копируем файлы проекта
COPY . .

# Загружаем зависимости Go
RUN go mod tidy

# Собираем приложение
ENV CGO_ENABLED=1
RUN apk add --no-cache gcc musl-dev

RUN go build -o main .

# Указываем порт, который будет использоваться
EXPOSE 8080

# Запускаем приложение
CMD ["/app/main"]
