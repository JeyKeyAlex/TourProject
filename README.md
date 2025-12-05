# TourProject

Тестовый проект для работы с пользователями и тестирования различных технологий: RabbitMQ, Kafka, PostgreSQL, Redis, gRPC и других.

## Описание

Проект представляет собой микросервис для управления пользователями, построенный на чистой архитектуре с использованием Go. Проект демонстрирует интеграцию различных технологий и паттернов проектирования.

## Архитектура

Проект следует принципам чистой архитектуры и разделен на следующие слои:

- **Transport Layer** (`internal/transport/http`) - HTTP обработчики и маршрутизация
- **Endpoint Layer** (`internal/endpoint`) - точки входа для бизнес-логики (go-kit endpoints)
- **Service Layer** (`internal/service`) - бизнес-логика приложения
- **Database Layer** (`internal/database`) - работа с базой данных
- **Entities** (`internal/entities`) - доменные модели

## Технологии

### Реализовано

- **Go 1.24+** - основной язык программирования
- **PostgreSQL** - основная база данных (pgx/v5)
- **Chi Router** - HTTP роутер
- **Go-Kit** - набор инструментов для микросервисов
- **Zerolog** - структурированное логирование
- **Migrations** - SQL миграции для схемы БД

### Планируется к реализации

- **RabbitMQ** - брокер сообщений
- **Kafka** - потоковая обработка данных
- **Redis** - кэширование и сессии
- **gRPC** - RPC коммуникация
- **SMTP** - отправка email уведомлений
- **Workers** - фоновые задачи

## API Endpoints

### User Service

- `GET /UserService/users` - получить список пользователей
- `POST /UserService/users` - создать нового пользователя
- `GET /UserService/users/{id}` - получить пользователя по ID
- `GET /ping` - health check

## Установка и запуск

### Требования

- Go 1.24+
- PostgreSQL 12+

### Настройка

1. Клонируйте репозиторий:
```bash
git clone <repository-url>
cd TestProject
```

2. Установите зависимости:
```bash
go mod download
```

3. Примените миграции:
```bash
# Применить миграции
go run cmd/migrate/up/main.go

# Откатить миграции (если нужно)
go run cmd/migrate/down/main.go
```

4. Запустите приложение:
```bash
go run cmd/app/main.go
```

## Конфигурация

Конфигурация приложения загружается из переменных окружения. Основные параметры:

- `RWDB_CONNECTION_STRING` - строка подключения к PostgreSQL (обязательно)
- `HTTP_ADDRESS` - адрес HTTP сервера (по умолчанию: `:8080`)
- `LOG_LEVEL` - уровень логирования (по умолчанию: `info`)
- `RUNTIME_USE_CPUS` - количество используемых CPU (0 = все доступные)
- `RUNTIME_MAX_THREADS` - максимальное количество потоков (0 = 10000)

Полный список параметров конфигурации можно найти в `internal/config/confing.go`.

## Разработка

Проект использует стандартные Go практики и паттерны:

- Clean Architecture
- Dependency Injection
- Repository Pattern
- Middleware для HTTP запросов
- Структурированное логирование

