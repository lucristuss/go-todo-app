# Используем официальный образ Golang
FROM golang:1.23-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

RUN apk add --no-cache bash

# Копируем файлы go.mod и go.sum
COPY go.mod ./ 
COPY go.sum ./ 

# Загружаем зависимости
RUN go mod download

# Копируем остальные файлы проекта
COPY . .

# Копируем скрипт wait-for-it в контейнер
COPY wait-for-it.sh /wait-for-it.sh

# Даем права на выполнение скрипта
RUN chmod +x /wait-for-it.sh

# Загружаем зависимости Go
RUN go mod tidy

# Устанавливаем нужные пакеты для компиляции
ENV CGO_ENABLED=1
RUN apk add --no-cache gcc musl-dev

# Собираем приложение
RUN go build -o main .

# Указываем порт, который будет использоваться
EXPOSE 8081

# Запускаем приложение с ожиданием готовности базы данных
CMD ["/wait-for-it.sh", "todo-db:5432", "--", "/app/main"]
