# Доменная модель Equipment

## Обзор

Доменная модель `Equipment` описывает физические объекты оборудования в системе CMMS и их классификацию согласно спецификации B2MML (ANSI/ISA-95).

## Структура модели

### Агрегаты

#### Equipment (Оборудование)
Основной агрегат, представляющий конкретный экземпляр оборудования.

**Поля:**
- `id` - уникальный идентификатор оборудования (Value Object)
- `data` - исходные данные B2MML типа `EquipmentType`
- `class` - класс (категория) оборудования
- `properties` - свойства конкретного оборудования
- `children` - иерархия дочернего оборудования
- `operatingStatus` - статус эксплуатации (active, inactive, maintenance)
- `version` - версия для оптимистичной блокировки

**Методы:**
- `ID()` - получить идентификатор
- `Class()` - получить класс оборудования
- `GetB2MMLData()` - получить исходные B2MML данные
- `Properties()` - получить все свойства
- `SetOperatingStatus(status)` - изменить статус
- `GetOperatingStatus()` - получить текущий статус
- `IsActive()` - проверить, активно ли оборудование
- `AddChild(child)` - добавить дочернее оборудование
- `Children()` - получить иерархию
- `AddProperty(prop)` - добавить свойство
- `EffectiveDate()` - получить период активности
- `PhysicalAssetRef()` - получить ссылку на физический актив
- `HierarchyScope()` - получить уровень иерархии
- `EquipmentLevel()` - получить уровень в классификации
- `Version()` - получить версию

#### EquipmentClass (Класс оборудования)
Представляет группу оборудования с похожими характеристиками.

**Поля:**
- `id` - уникальный идентификатор класса
- `data` - исходные данные B2MML типа `EquipmentClassType`
- `properties` - свойства класса
- `children` - иерархия дочерних классов

**Методы:**
- `ID()` - получить идентификатор
- `GetB2MMLData()` - получить исходные B2MML данные
- `Properties()` - получить все свойства класса
- `AddProperty(prop)` - добавить свойство класса
- `AddChild(child)` - добавить дочерний класс
- `Children()` - получить иерархию классов

### Value Objects

#### EquipmentID
Идентификатор оборудования. Гарантирует, что ID не пустой и валидный.

```go
id, err := NewEquipmentID("PUMP-001")
if err != nil {
    log.Fatal(err)
}
```

#### EquipmentClassID
Идентификатор класса оборудования.

```go
classID, err := NewEquipmentClassID("CENTRIFUGAL_PUMP")
```

#### EquipmentPropertyID
Идентификатор свойства оборудования.

```go
propID, err := NewEquipmentPropertyID("POWER_RATING")
```

#### PropertyValue
Значение свойства с метаинформацией.

```go
value := NewPropertyValueWithUnit("7.5", "float", "kW")
value.SetDescription("Номинальная мощность")
```

### Доменные события

#### EquipmentCreatedEvent
Событие, возникающее при создании нового оборудования.

#### EquipmentStatusChangedEvent
Событие изменения статуса оборудования.

#### EquipmentClassCreatedEvent
Событие создания нового класса оборудования.

## Интеграция с B2MML

Модель использует сгенерированные структуры из B2MML XSD:
- `b2mml.EquipmentType` - основные данные оборудования
- `b2mml.EquipmentClassType` - данные класса оборудования
- `b2mml.EquipmentPropertyType` - свойства оборудования
- `b2mml.EquipmentClassPropertyType` - свойства класса
- `b2mml.IdentifierType` - идентификаторы
- `b2mml.DateTimeType` - даты и времена
- `b2mml.HierarchyScopeType` - уровень иерархии
- `b2mml.EquipmentLevelType` - уровень классификации

Структуры хранят оригинальные B2MML данные для интеграции и сохранения, но предоставляют удобный доменный интерфейс.

## Примеры использования

### Создание оборудования

```go
// Создать ID
id, _ := NewEquipmentID("PUMP-001")
classID, _ := NewEquipmentClassID("CENTRIFUGAL")

// Создать класс
class := NewEquipmentClass(classID, b2mmlClassData)

// Создать оборудование
equipment := NewEquipment(id, b2mmlData, class)

// Установить статус
equipment.SetOperatingStatus(OperatingStatusActive)
```

### Работа с иерархией

```go
// Добавить дочернее оборудование
childID, _ := NewEquipmentID("PUMP-001-MOTOR")
child := NewEquipment(childID, childB2MMLData, class)
equipment.AddChild(child)

// Получить все дочернее оборудование
children := equipment.Children()
```

### Работа со свойствами

```go
// Создать свойство
propID, _ := NewEquipmentPropertyID("POWER_RATING")
value := NewPropertyValueWithUnit("7.5", "float", "kW")
property := NewEquipmentProperty(propID, b2mmlPropData, value)

// Добавить к оборудованию
equipment.AddProperty(property)
```

## Доменные ошибки

- `ErrEquipmentIDEmpty` - пустой идентификатор оборудования
- `ErrEquipmentClassIDEmpty` - пустой ID класса
- `ErrEquipmentPropertyIDEmpty` - пустой ID свойства
- `ErrEquipmentClassPropertyIDEmpty` - пустой ID свойства класса
- `ErrEquipmentNotFound` - оборудование не найдено
- `ErrEquipmentAlreadyExists` - оборудование уже существует
- `ErrEquipmentInvalidStatus` - некорректный статус
- `ErrEquipmentClassNotFound` - класс оборудования не найден
- `ErrEquipmentClassAlreadyExists` - класс уже существует

## Архитектурные решения

1. **Инкапсуляция**: Все поля структур приватные, доступ через методы
2. **Value Objects**: ID используют Pattern Value Object для безопасности
3. **Версионирование**: Поле `version` для оптимистичной блокировки при обновлении
4. **B2MML интеграция**: Хранение оригинальных B2MML данных для сохранения и обмена
5. **Доменные события**: Для асинхронного взаимодействия с другими агрегатами
6. **Иерархия**: Поддержка древовидной структуры как для оборудования, так и для классов
