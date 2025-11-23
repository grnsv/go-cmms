package model

import (
	"errors"

	"github.com/grnsv/go-cmms/internal/domain/model/b2mml"
)

// NewPropertyValue создаёт новое значение свойства
func NewPropertyValue(value, dataType string) PropertyValue {
	return PropertyValue{
		value:    value,
		dataType: dataType,
	}
}

// NewPropertyValueWithUnit создаёт значение свойства с единицей измерения
func NewPropertyValueWithUnit(value, dataType, unit string) PropertyValue {
	return PropertyValue{
		value:    value,
		dataType: dataType,
		unit:     unit,
	}
}

// Value возвращает строковое значение свойства
func (pv PropertyValue) Value() string {
	return pv.value
}

// DataType возвращает тип данных свойства
func (pv PropertyValue) DataType() string {
	return pv.dataType
}

// Unit возвращает единицу измерения
func (pv PropertyValue) Unit() string {
	return pv.unit
}

// Description возвращает описание свойства
func (pv PropertyValue) Description() string {
	return pv.description
}

// SetDescription устанавливает описание свойства
func (pv *PropertyValue) SetDescription(desc string) {
	pv.description = desc
}

// NewEquipmentProperty создаёт новое свойство оборудования
func NewEquipmentProperty(
	id EquipmentPropertyID,
	b2mmlData *b2mml.EquipmentPropertyType,
	value PropertyValue,
) *EquipmentProperty {
	return &EquipmentProperty{
		id:    id,
		data:  b2mmlData,
		value: value,
	}
}

// ID возвращает идентификатор свойства
func (ep *EquipmentProperty) ID() EquipmentPropertyID {
	return ep.id
}

// GetB2MMLData возвращает исходные данные B2MML
func (ep *EquipmentProperty) GetB2MMLData() *b2mml.EquipmentPropertyType {
	return ep.data
}

// Value возвращает значение свойства
func (ep *EquipmentProperty) Value() PropertyValue {
	return ep.value
}

// NewEquipmentClassProperty создаёт новое свойство класса оборудования
func NewEquipmentClassProperty(
	id EquipmentClassPropertyID,
	b2mmlData *b2mml.EquipmentClassPropertyType,
) *EquipmentClassProperty {
	return &EquipmentClassProperty{
		id:         id,
		data:       b2mmlData,
		properties: make([]*EquipmentClassProperty, 0),
	}
}

// ID возвращает идентификатор свойства класса
func (ecp *EquipmentClassProperty) ID() EquipmentClassPropertyID {
	return ecp.id
}

// GetB2MMLData возвращает исходные данные B2MML
func (ecp *EquipmentClassProperty) GetB2MMLData() *b2mml.EquipmentClassPropertyType {
	return ecp.data
}

// AddChild добавляет дочернее свойство
func (ecp *EquipmentClassProperty) AddChild(child *EquipmentClassProperty) error {
	if child == nil {
		return errors.New("child property cannot be nil")
	}
	ecp.properties = append(ecp.properties, child)
	return nil
}

// Children возвращает дочерние свойства
func (ecp *EquipmentClassProperty) Children() []*EquipmentClassProperty {
	return ecp.properties
}
