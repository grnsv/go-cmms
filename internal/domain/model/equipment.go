package model

import (
	"errors"

	"github.com/grnsv/go-cmms/internal/domain/model/b2mml"
)

// Equipment представляет физический объект оборудования в системе.
// Это агрегат, содержащий информацию об оборудовании и его классификацию.
type Equipment struct {
	id              EquipmentID
	data            *b2mml.EquipmentType
	class           *EquipmentClass
	properties      []*EquipmentProperty
	children        []*Equipment
	operatingStatus OperatingStatus
	version         int64
}

// EquipmentID является Value Object идентификатора оборудования
type EquipmentID struct {
	value string
}

// NewEquipmentID создаёт новый идентификатор оборудования
func NewEquipmentID(value string) (EquipmentID, error) {
	if value == "" {
		return EquipmentID{}, ErrEquipmentIDEmpty
	}
	return EquipmentID{value: value}, nil
}

func (id EquipmentID) String() string {
	return id.value
}

// OperatingStatus представляет статус эксплуатации оборудования
type OperatingStatus string

const (
	OperatingStatusActive    OperatingStatus = "active"
	OperatingStatusInactive  OperatingStatus = "inactive"
	OperatingStatusMaintenance OperatingStatus = "maintenance"
)

// EquipmentClass представляет класс (категорию) оборудования.
// Оборудование может принадлежать нескольким классам.
type EquipmentClass struct {
	id         EquipmentClassID
	data       *b2mml.EquipmentClassType
	properties []*EquipmentClassProperty
	children   []*EquipmentClass // иерархия классов
}

// EquipmentClassID является Value Object идентификатора класса оборудования
type EquipmentClassID struct {
	value string
}

// NewEquipmentClassID создаёт новый идентификатор класса оборудования
func NewEquipmentClassID(value string) (EquipmentClassID, error) {
	if value == "" {
		return EquipmentClassID{}, ErrEquipmentClassIDEmpty
	}
	return EquipmentClassID{value: value}, nil
}

func (id EquipmentClassID) String() string {
	return id.value
}

// EquipmentProperty представляет свойство конкретного экземпляра оборудования
type EquipmentProperty struct {
	id    EquipmentPropertyID
	data  *b2mml.EquipmentPropertyType
	value PropertyValue
}

// EquipmentPropertyID является Value Object идентификатора свойства
type EquipmentPropertyID struct {
	value string
}

// NewEquipmentPropertyID создаёт новый идентификатор свойства
func NewEquipmentPropertyID(value string) (EquipmentPropertyID, error) {
	if value == "" {
		return EquipmentPropertyID{}, ErrEquipmentPropertyIDEmpty
	}
	return EquipmentPropertyID{value: value}, nil
}

func (id EquipmentPropertyID) String() string {
	return id.value
}

// PropertyValue представляет значение свойства оборудования
type PropertyValue struct {
	value       string
	dataType    string // например: "string", "integer", "float", "datetime"
	unit        string // единица измерения, если применимо
	description string
}

// EquipmentClassProperty представляет свойство класса оборудования
type EquipmentClassProperty struct {
	id         EquipmentClassPropertyID
	data       *b2mml.EquipmentClassPropertyType
	properties []*EquipmentClassProperty // иерархия свойств
}

// EquipmentClassPropertyID является Value Object идентификатора свойства класса
type EquipmentClassPropertyID struct {
	value string
}

// NewEquipmentClassPropertyID создаёт новый идентификатор свойства класса
func NewEquipmentClassPropertyID(value string) (EquipmentClassPropertyID, error) {
	if value == "" {
		return EquipmentClassPropertyID{}, ErrEquipmentClassPropertyIDEmpty
	}
	return EquipmentClassPropertyID{value: value}, nil
}

func (id EquipmentClassPropertyID) String() string {
	return id.value
}

// EquipmentHierarchyLevel представляет уровень оборудования в иерархии
type EquipmentHierarchyLevel struct {
	scope *b2mml.HierarchyScopeType
	level *b2mml.EquipmentLevelType
}

// Конструкторы агрегата Equipment

// NewEquipment создаёт новый экземпляр оборудования
func NewEquipment(
	id EquipmentID,
	b2mmlData *b2mml.EquipmentType,
	class *EquipmentClass,
) *Equipment {
	return &Equipment{
		id:         id,
		data:       b2mmlData,
		class:      class,
		properties: make([]*EquipmentProperty, 0),
		children:   make([]*Equipment, 0),
		version:    1,
	}
}

// NewEquipmentClass создаёт новый класс оборудования
func NewEquipmentClass(
	id EquipmentClassID,
	b2mmlData *b2mml.EquipmentClassType,
) *EquipmentClass {
	return &EquipmentClass{
		id:         id,
		data:       b2mmlData,
		properties: make([]*EquipmentClassProperty, 0),
		children:   make([]*EquipmentClass, 0),
	}
}

// Методы агрегата Equipment

// ID возвращает идентификатор оборудования
func (e *Equipment) ID() EquipmentID {
	return e.id
}

// Class возвращает класс оборудования
func (e *Equipment) Class() *EquipmentClass {
	return e.class
}

// GetB2MMLData возвращает исходные данные B2MML для интеграции и сохранения
func (e *Equipment) GetB2MMLData() *b2mml.EquipmentType {
	return e.data
}

// Properties возвращает свойства оборудования
func (e *Equipment) Properties() []*EquipmentProperty {
	return e.properties
}

// SetOperatingStatus устанавливает статус эксплуатации
func (e *Equipment) SetOperatingStatus(status OperatingStatus) {
	e.operatingStatus = status
	e.version++
}

// GetOperatingStatus возвращает текущий статус эксплуатации
func (e *Equipment) GetOperatingStatus() OperatingStatus {
	return e.operatingStatus
}

// Version возвращает версию агрегата (для оптимистичной блокировки)
func (e *Equipment) Version() int64 {
	return e.version
}

// Методы EquipmentClass

// ID возвращает идентификатор класса
func (ec *EquipmentClass) ID() EquipmentClassID {
	return ec.id
}

// GetB2MMLData возвращает исходные данные B2MML
func (ec *EquipmentClass) GetB2MMLData() *b2mml.EquipmentClassType {
	return ec.data
}

// AddChild добавляет дочернее оборудование
func (e *Equipment) AddChild(child *Equipment) error {
	if child == nil {
		return ErrEquipmentNotFound
	}
	e.children = append(e.children, child)
	e.version++
	return nil
}

// Children возвращает все дочернее оборудование
func (e *Equipment) Children() []*Equipment {
	return e.children
}

// AddProperty добавляет свойство к оборудованию
func (e *Equipment) AddProperty(prop *EquipmentProperty) error {
	if prop == nil {
		return errors.New("property cannot be nil")
	}
	e.properties = append(e.properties, prop)
	e.version++
	return nil
}

// IsActive проверяет, активно ли оборудование
func (e *Equipment) IsActive() bool {
	return e.operatingStatus == OperatingStatusActive
}

// EffectiveDate возвращает период активности оборудования
func (e *Equipment) EffectiveDate() (startDate, endDate *b2mml.DateTimeType, err error) {
	if e.data == nil || e.data.EffectiveStartDate == nil {
		return nil, nil, errors.New("effective dates not set")
	}
	return e.data.EffectiveStartDate, e.data.EffectiveEndDate, nil
}

// PhysicalAssetRef возвращает ссылку на физический актив
func (e *Equipment) PhysicalAssetRef() *b2mml.IdentifierType {
	if e.data == nil {
		return nil
	}
	return e.data.PhysicalAssetID
}

// HierarchyScope возвращает иерархический уровень оборудования
func (e *Equipment) HierarchyScope() *b2mml.HierarchyScopeType {
	if e.data == nil {
		return nil
	}
	return e.data.HierarchyScope
}

// EquipmentLevel возвращает уровень оборудования в иерархии
func (e *Equipment) EquipmentLevel() *b2mml.EquipmentLevelType {
	if e.data == nil {
		return nil
	}
	return e.data.EquipmentLevel
}

// Методы для EquipmentClass

// AddProperty добавляет свойство к классу оборудования
func (ec *EquipmentClass) AddProperty(prop *EquipmentClassProperty) error {
	if prop == nil {
		return errors.New("property cannot be nil")
	}
	ec.properties = append(ec.properties, prop)
	return nil
}

// Properties возвращает свойства класса
func (ec *EquipmentClass) Properties() []*EquipmentClassProperty {
	return ec.properties
}

// AddChild добавляет дочерний класс оборудования
func (ec *EquipmentClass) AddChild(child *EquipmentClass) error {
	if child == nil {
		return ErrEquipmentClassNotFound
	}
	ec.children = append(ec.children, child)
	return nil
}

// Children возвращает дочерние классы
func (ec *EquipmentClass) Children() []*EquipmentClass {
	return ec.children
}
