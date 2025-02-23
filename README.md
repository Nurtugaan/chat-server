# Chat Server

## Описание
Этот проект представляет собой WebSocket-чат-сервер на языке Go с возможностью сохранения истории сообщений в PostgreSQL. Сервер поддерживает WebSocket-подключения для общения пользователей в реальном времени, а также предоставляет REST API для получения истории сообщений.

## Стек технологий
- **Backend:** Go (Gin, Gorilla WebSocket)
- **База данных:** PostgreSQL
- **Контейнеризация:** Docker, Docker Compose
- **Конфигурация:** Viper

## Возможности
- Поддержка WebSocket-соединений для обмена сообщениями в реальном времени
- Сохранение сообщений в базе данных PostgreSQL
- REST API для получения истории сообщений
- Запуск через Docker и Docker Compose

## Структура проекта
```
chat-server/
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── main.go
├── config/
│   └── config.go
├── internal/
│   ├── handler/
│   │   └── handler.go
│   ├── model/
│   │   └── message.go
│   ├── repository/
│   │   └── postgres.go
│   ├── service/
│   │   └── service.go
│   └── websocket/
│       └── client.go
└── migrations/
    └── 001_create_messages_table.sql
```

## Установка и запуск

### 1. Клонирование репозитория
```sh
git clone https://github.com/yourusername/chat-server.git
cd chat-server
```

### 2. Запуск с Docker
```sh
docker-compose up --build
```
Сервер будет доступен по адресу `http://localhost:8080`.

### 3. Запуск без Docker
#### Установка зависимостей
```sh
go mod download
```

#### Запуск PostgreSQL (если не используется Docker)
```sh
docker run --name chat-db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=chatapp -p 5432:5432 -d postgres:14
```

#### Применение миграций
```sh
go run migrations/001_create_messages_table.sql
```

#### Запуск сервера
```sh
go run main.go
```

## API

### WebSocket API
- **Подключение:** `ws://localhost:8080/ws?username=YOUR_NAME`
- **Отправка сообщений:** Текстовые сообщения передаются в JSON-формате:
  ```json
  {
    "username": "Alice",
    "content": "Hello, world!"
  }
  ```
- **Получение сообщений:** Клиент получает JSON-объекты с сообщениями в реальном времени.

### REST API
- **Получение истории сообщений:**
  ```sh
  GET /messages
  ```
  **Ответ:**
  ```json
  [
    {
      "id": 1,
      "username": "Alice",
      "content": "Hello, world!",
      "created_at": "2025-02-23T12:00:00Z"
    }
  ]
  ```

## Переменные окружения
| Переменная       | Описание              | Значение по умолчанию |
|------------------|----------------------|----------------------|
| DB_HOST         | Хост базы данных      | db                   |
| DB_PORT         | Порт базы данных      | 5432                 |
| DB_USER         | Пользователь БД       | postgres             |
| DB_PASSWORD     | Пароль БД             | password             |
| DB_NAME         | Название БД           | chatapp              |
