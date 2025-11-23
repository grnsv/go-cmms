# SQLC Generated Persistence Layer

Этот пакет содержит автоматически сгенерированный код доступа к БД с использованием **sqlc**.

## Структура

### Миграции
- `migrations/001_create_equipment_tables.sql` - схема БД Equipment

**Таблицы:**
- `equipment_classes` - классы оборудования
- `equipment_class_properties` - свойства классов
- `equipment` - экземпляры оборудования
- `equipment_properties` - свойства оборудования
- `equipment_class_mappings` - связь многие-ко-многим между equipment и equipment_classes

### Запросы
- `queries/equipment.sql` - 27 запросов для работы с Equipment

**Категории запросов:**
1. Equipment Classes (CRUD операции)
2. Equipment Class Properties
3. Equipment (CRUD, статусы, иерархия)
4. Equipment Properties
5. Equipment Class Mappings (связи)

### Сгенерированный код

#### `db.go` (31 строка)
Инициализация БД и Querier:
```go
type DB struct {
    conn *sql.DB
}

func New(conn *sql.DB) *Queries
```

#### `models.go` (84 строки)
Структуры данных, соответствующие таблицам:
- `Equipment`
- `EquipmentClass`
- `EquipmentClassMapping`
- `EquipmentClassProperty`
- `EquipmentProperty`

#### `equipment.sql.go` (1276 строк)
24 метода на `*Queries`:
- `CreateEquipment`, `GetEquipmentByID`, `GetEquipmentByExternalID`
- `ListEquipment`, `ListEquipmentByStatus`, `ListChildEquipment`
- `UpdateEquipmentStatus`, `DeleteEquipment`
- `CreateEquipmentClass`, `GetEquipmentClassByID`, `ListAllEquipmentClasses`
- `UpdateEquipmentClass`, `DeleteEquipmentClass`
- `CreateEquipmentProperty`, `ListEquipmentProperties`, `UpdateEquipmentProperty`, `DeleteEquipmentProperty`
- `CreateEquipmentClassProperty`, `ListEquipmentClassProperties`
- `AddEquipmentToClass`, `RemoveEquipmentFromClass`
- `ListEquipmentClassesForEquipment`, `ListEquipmentByClass`

#### `querier.go` (45 строк)
Интерфейс `Querier` для зависимости injection:
```go
type Querier interface {
    // все методы выше
}
```

## Использование

### Инициализация
```go
import "github.com/grnsv/go-cmms/internal/infrastructure/persistence/sqlc"

db := sql.Open("postgres", dsn)
queries := sqlc.New(db)
```

### Примеры

```go
// Создание оборудования
eq, err := queries.CreateEquipment(ctx, &sqlc.CreateEquipmentParams{
    ExternalID: "PUMP-001",
    Version: sql.NullString{String: "1.0", Valid: true},
    OperatingStatus: sql.NullString{String: "active", Valid: true},
    // ...другие параметры
})

// Получение по ID
equipment, err := queries.GetEquipmentByID(ctx, equipmentID)

// Список с пагинацией
equipments, err := queries.ListEquipment(ctx, &sqlc.ListEquipmentParams{
    Limit:  10,
    Offset: 0,
})

// Обновление статуса
updated, err := queries.UpdateEquipmentStatus(ctx, &sqlc.UpdateEquipmentStatusParams{
    ID:           equipmentID,
    OperatingStatus: "maintenance",
})
```

## Конфигурация

`sqlc.yaml` определяет:
- **queries**: путь к SQL запросам
- **schema**: путь к миграциям
- **package**: `postgres` (имя пакета)
- **overrides**: типы Go для специальных типов PostgreSQL
  - `uuid` → `github.com/google/uuid.UUID`
  - `timestamptz` → `time.Time`
  - `jsonb` → `encoding/json.RawMessage`

## Безопасность

- Все запросы используют параметризованные запросы (защита от SQL инъекций)
- sqlc компилирует SQL в compile-time
- Типизированные параметры и результаты

## Дополнительно

Для регенерации кода после изменения SQL запросов:
```bash
cd internal/infrastructure/persistence/sqlc
go generate ./...
```

Или напрямую:
```bash
sqlc generate -f internal/infrastructure/persistence/sqlc/sqlc.yaml
```
