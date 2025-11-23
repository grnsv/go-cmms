# GO-CMMS - Computerized Maintenance Management System

Приложение для управления производственным оборудованием на основе схем B2MML/ANSI-ISA-95.

## 🏗️ Архитектура

Приложение организовано в соответствии с принципами **Clean Architecture**:

```
Presentation (API) → Application (Use Cases) → Domain (Business Logic) → Infrastructure (DB)
```

### Слои

1. **cmd/server** - точка входа приложения
2. **internal/api** - HTTP API (handlers, OpenAPI spec с ogen)
3. **internal/app** - use cases приложения
4. **internal/domain** - бизнес-логика, модели, правила
5. **internal/infrastructure** - работа с БД, репозитории
6. **internal/config** - конфигурация приложения

## 📦 Компоненты

### Доменная модель (internal/domain/model/)
- **Equipment** - агрегат оборудования
- **EquipmentClass** - категория оборудования
- **Value Objects** - EquipmentID, PropertyValue, и др.
- **Events** - доменные события
- **Errors** - доменные ошибки

### Слой приложения (internal/app/)
- **ListEquipmentUseCase** - получение списка оборудования
- **GetEquipmentByIDUseCase** - поиск по ID
- **CreateEquipmentUseCase** - создание оборудования

### Слой доступа к данным
- **sqlc** - автоматическая генерация SQL методов
- **Repository** - интерфейсы и реализация доступа к данным
- **PostgreSQL** - основная БД

### API
- **ogen** - генерация HTTP handlers из OpenAPI спецификации
- **spec.yaml** - OpenAPI 3.1 спецификация
- **handler** - адаптеры для преобразования HTTP в use cases

## 🛠️ Технологический стек

- **Go 1.21+**
- **PostgreSQL 12+**
- **sqlc** - type-safe SQL queries
- **ogen** - OpenAPI code generation
- **uuid** - уникальные идентификаторы
- **lib/pq** - PostgreSQL driver

## 🚀 Быстрый старт

### Предварительно

1. Убедитесь, что установлены:
   - Go 1.21+
   - PostgreSQL 12+

2. Клонируйте репозиторий:
```bash
git clone https://github.com/grnsv/go-cmms.git
cd go-cmms
```

### Установка

1. Скопируйте конфигурацию:
```bash
cp .env.example .env
```

2. Отредактируйте переменные окружения:
```bash
vim .env
```

3. Создайте БД:
```bash
createdb -U user go_cmms
```

4. Установите зависимости:
```bash
go mod download
```

5. Соберите приложение:
```bash
go build -o server ./cmd/server
```

6. Запустите:
```bash
./server
```

Server запустится на `http://0.0.0.0:8080`

## 📚 API Endpoints

### Health Check
```
GET /health
```

### Equipment Management
```
GET    /api/v1/equipment              # Список оборудования
POST   /api/v1/equipment              # Создать оборудование
GET    /api/v1/equipment/{id}         # Получить по ID
```

## 🗂️ Структура проекта

```
go-cmms/
├── cmd/
│   └── server/
│       └── main.go                   # Точка входа
├── internal/
│   ├── api/
│   │   ├── spec.yaml                 # OpenAPI спецификация
│   │   ├── handler/                  # HTTP handlers
│   │   └── ogen/                     # Сгенерировано ogen
│   ├── app/
│   │   └── equipment.go              # Use cases
│   ├── config/
│   │   └── config.go                 # Конфигурация
│   ├── domain/
│   │   ├── model/                    # Доменные модели
│   │   └── repository/               # Repository interfaces
│   └── infrastructure/
│       ├── database.go               # DB инициализация
│       └── postgres/
│           ├── repository/           # Repository implementations
│           └── sqlc/                 # Сгенерировано sqlc
├── .env.example                      # Пример конфигурации
├── go.mod
├── go.sum
├── README.md
└── ARCHITECTURE.md
```

## 🔧 Разработка

### Регенерация кода

#### SQLC (доступ к БД)
```bash
go tool sqlc generate -f internal/infrastructure/postgres/sqlc/sqlc.yaml
```

#### OGEN (API handlers)
```bash
cd internal/api && go generate ./...
```

## 📊 База данных

### Таблицы

- `equipment_classes` - классы оборудования
- `equipment_class_properties` - свойства классов
- `equipment` - экземпляры оборудования
- `equipment_properties` - свойства оборудования
- `equipment_class_mappings` - связь M-N между equipment и classes

### Особенности

- Soft delete через `deleted_at`
- Оптимистичная блокировка (`record_version`)
- JSONB для B2MML данных
- Полная индексация для производительности

## 🧪 Тестирование

```bash
go test ./...
```

С покрытием:
```bash
go test -cover ./...
```

## 📝 Лицензия

MIT
