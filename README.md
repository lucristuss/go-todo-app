
# Go ToDo App

Это веб-приложение для управления задачами (ToDo List), написанное на языке Go с использованием фреймворка Gin. Приложение использует базу данных PostgreSQL для хранения данных о задачах.

## Функциональность

Приложение предоставляет следующий REST API для управления задачами:

- **POST /tasks** — создание новой задачи.
- **GET /tasks** — получение списка всех задач.
- **GET /tasks/{id}** — получение задачи по ID.
- **PUT /tasks/{id}** — обновление задачи по ID.
- **DELETE /tasks/{id}** — удаление задачи по ID.

## Разработка

### Требования

- Go версии 1.18 или выше
- Docker и Docker Compose (для локального развертывания)
- PostgreSQL (используется как база данных)

### Структура проекта

```
/go-todo-app
│
├── main.go                # Основной файл приложения
├── models.go              # Структуры для работы с базой данных
├── handlers.go            # Функции-обработчики для API
├── db.go                  # Инициализация соединения с базой данных
├── Dockerfile             # Конфигурация для Docker
├── docker-compose.yml     # Конфигурация для Docker Compose
├── main_test.go           # Тесты для API
└── README.md              # Документация проекта
```

### Установка и запуск

1. **Клонируй репозиторий:**

   ```sh
   git clone https://github.com/lucristuss/go-todo-app.git
   cd go-todo-app
   ```

2. **Запусти приложение с помощью Docker Compose:**

   Для развертывания приложения с Docker и PostgreSQL, используйте команду:

   ```sh
   docker-compose up --build
   ```

   Это создаст и запустит два контейнера:
   - Контейнер с веб-приложением на Go.
   - Контейнер с базой данных PostgreSQL.

3. **Тестирование API:**

   После успешного запуска приложения, API будет доступно на порту `8081`.

   Пример запросов через `curl`:

   - **Получить список задач:**
     ```sh
     curl http://localhost:8081/tasks
     ```

   - **Добавить задачу:**
     ```sh
     curl -X POST http://localhost:8081/tasks -d '{"name": "Новая задача"}' -H "Content-Type: application/json"
     ```

   - **Обновить задачу:**
     ```sh
     curl -X PUT http://localhost:8081/tasks/{id} -d '{"name": "Обновленная задача"}' -H "Content-Type: application/json"
     ```

   - **Удалить задачу:**
     ```sh
     curl -X DELETE http://localhost:8081/tasks/{id}
     ```
## Ошибки и обработка ошибок

В приложении предусмотрена обработка ошибок, которая позволяет:
- При получении некорректных данных от пользователя (например, при добавлении задачи с неполными данными), сервер возвращает ошибку с кодом `400` и сообщением об ошибке.
- При попытке работы с несуществующей задачей сервер вернет ошибку с кодом `404` (Not Found).
- В случае серверных ошибок будет возвращен статус `500` с описанием проблемы.

Ошибки обрабатываются через middleware, которые логируют все сообщения об ошибках и позволяют приложению корректно реагировать на разные типы ошибок.

Пример обработки ошибки в коде:

```go
func handleError(err error) {
    if err != nil {
        log.Error("Ошибка при обработке запроса: ", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка сервера"})
    }
}
```

4. **Запуск тестов:**

   Для запуска тестов используйте команду:

   ```sh
   go test -v
   ```

### Развертывание в облаке

Для развертывания приложения в облаке, используйте Docker и настройте контейнеризацию в облачной платформе (например, AWS, GCP, Azure или DigitalOcean).
