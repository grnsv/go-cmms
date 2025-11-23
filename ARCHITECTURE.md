# GO-CMMS Application

## Структура приложения

```
├── cmd/server/main.go          # Точка входа приложения
├── internal/
│   ├── config/                 # Конфигурация приложения
│   ├── infrastructure/         # Инфраструктура (БД, репозитории)
│   ├── domain/                 # Доменная модель
│   ├── app/                    # Use cases приложения
│   └── api/                    # API layer (handlers, OpenAPI spec)
└── go.mod
```

## Компоненты приложения

### 1. Config Layer (`internal/config/config.go`)
Загружает конфигурацию из переменных окружения:
- `SERVER_HOST` - адрес сервера (по умолчанию 0.0.0.0)
- `SERVER_PORT` - порт (по умолчанию 8080)
- `DATABASE_URL` - URL подключения к PostgreSQL
- `LOG_LEVEL` - уровень логирования (debug, info, warn, error)

### 2. Infrastructure Layer
**Database** (`internal/infrastructure/database.go`):
- Инициализирует подключение к PostgreSQL
- Настраивает пул соединений
- Возвращает `*sqlc.Queries` для работы с БД

**Repository** (`internal/infrastructure/postgres/repository/equipment.go`):
- Реализация репозиториев Equipment и EquipmentClass
- Использует сгенерированный sqlc код
- Реализует интерфейсы из `internal/domain/repository`

### 3. Domain Layer
**Model** (`internal/domain/model/equipment.go`):
- Агрегаты: Equipment, EquipmentClass
- Value Objects: EquipmentID, PropertyValue
- Доменные события и ошибки

**Repository Interfaces** (`internal/domain/repository/equipment.go`):
- EquipmentRepository
- EquipmentClassRepository
- UnitOfWork

### 4. Application Layer (`internal/app/equipment.go`)
Use Cases:
- `ListEquipmentUseCase` - получение списка оборудования
- `GetEquipmentByIDUseCase` - получение по ID
- `CreateEquipmentUseCase` - создание нового

### 5. API Layer
**Handler** (`internal/api/handler/handler.go`):
- Реализует HTTP endpoints
- Преобразует HTTP запросы в use case вызовы
- Возвращает HTTP ответы

## Поток данных

```
HTTP Request
    ↓
Handler (API Layer)
    ↓
Use Case (Application Layer)
    ↓
Repository (Domain ↔ Infrastructure)
    ↓
SQLC Queries (sgenerated code)
    ↓
PostgreSQL Database
```

## Сборка и запуск

### Требования
- Go 1.21+
- PostgreSQL 12+
- sqlc v1.30.0+
- ogen (для генерации OpenAPI handlers)

### Установка зависимостей
```bash
go mod download
```

### Компиляция
```bash
go build ./cmd/server
```

### Запуск
```bash
# С конфигурацией по умолчанию
./server

# С переменными окружения
DATABASE_URL="postgres://user:pass@localhost/cmms" SERVER_PORT=8081 ./server
```

### Docker Compose (для БД)
```yaml
version: '3'
services:
  postgres:
    image: postgres:15
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
      POSTGRES_DB: go_cmms
    ports:
      - "5432:5432"
```

## API Endpoints

### Health Check
```
GET /health
```

### Equipment
```
GET  /api/v1/equipment           - Получить список
POST /api/v1/equipment           - Создать
GET  /api/v1/equipment/{id}      - Получить по ID
```

## Миграции БД

Миграции находятся в `internal/infrastructure/postgres/sqlc/migrations/`:
- `001_create_equipment_tables.sql` - создание таблиц Equipment

Для применения миграций используйте инструменты вроде:
- migrate
- Flyway
- Liquibase

## Регенерация кода

### sqlc (для слоя доступа к данным)
```bash
cd internal/infrastructure/postgres/sqlc
go generate ./...
```

### ogen (для API handlers)
```bash
cd internal/api
go generate ./...
```

## Конфигурация логирования

Текущий вывод - стандартный `log` пакет. Для production рекомендуется:
- zap
- logrus
- slog (Go 1.21+)

## TODO

- [ ] Полная реализация Repository методов (сейчас заглушки)
- [ ] Интеграция ogen для автоматического API
- [ ] Логирование (slog или zap)
- [ ] Миграции (migrate или Flyway)
- [ ] Tests (unit, integration)
- [ ] Docker image
- [ ] Kubernetes manifests
- [ ] CI/CD pipeline
