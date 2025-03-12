# Используем официальный образ Golang
FROM golang:1.21-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта
COPY . .

# Загружаем зависимости Go
RUN go mod init go-todo-app && go mod tidy

# Собираем приложение
RUN go build -o main .

# Указываем порт, который будет использоваться
EXPOSE 8080

# Запускаем приложение
CMD ["/app/main"]
