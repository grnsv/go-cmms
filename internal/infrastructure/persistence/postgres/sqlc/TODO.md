# SQLC Generated Code - Done ✓

- [x] queries/equipment.sql - основные запросы для Equipment
- [x] migrations/001_create_equipment_tables.sql - схема БД
- [x] Сгенерированный код:
  - db.go - инициализация
  - models.go - структуры данных (Equipment, EquipmentClass, etc.)
  - equipment.sql.go - методы для работы с БД
  - querier.go - интерфейс Querier

## Generated Methods

- CreateEquipment, GetEquipmentByID, GetEquipmentByExternalID
- ListEquipment, ListEquipmentByStatus, ListChildEquipment
- UpdateEquipmentStatus, DeleteEquipment
- CreateEquipmentClass, GetEquipmentClassByID, ListAllEquipmentClasses
- UpdateEquipmentClass, DeleteEquipmentClass
- CreateEquipmentProperty, ListEquipmentProperties, UpdateEquipmentProperty
- CreateEquipmentClassProperty, ListEquipmentClassProperties
- AddEquipmentToClass, RemoveEquipmentFromClass
- ListEquipmentClassesForEquipment, ListEquipmentByClass

## TODO

- [ ] Repository pattern implementation
- [ ] Service layer implementation
- [ ] Tests for repositories
