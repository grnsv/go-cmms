package model

import (
	"errors"

	"github.com/grnsv/go-cmms/internal/domain/model/b2mml"
)

// B2MMLAdapter помогает адаптировать данные между доменной моделью и B2MML
type B2MMLAdapter struct{}

// NewEquipmentFromB2MML создаёт Equipment из B2MML данных
func NewEquipmentFromB2MML(
	id EquipmentID,
	b2mmlData *b2mml.EquipmentType,
	class *EquipmentClass,
) (*Equipment, error) {
	if b2mmlData == nil {
		return nil, errors.New("b2mml data cannot be nil")
	}

	eq := NewEquipment(id, b2mmlData, class)

	// Преобразовать свойства из B2MML
	if b2mmlData.EquipmentProperty != nil {
		for _, prop := range b2mmlData.EquipmentProperty {
			if prop != nil && prop.ID != nil {
				propID, err := NewEquipmentPropertyID(prop.ID.Value)
				if err != nil {
					continue
				}

				value := PropertyValue{
					value:       prop.ID.Value,
					dataType:    "string",
				}

				equipmentProp := NewEquipmentProperty(propID, prop, value)
				_ = eq.AddProperty(equipmentProp)
			}
		}
	}

	// Преобразовать дочернее оборудование
	if b2mmlData.EquipmentChild != nil {
		for _, childB2MML := range b2mmlData.EquipmentChild {
			if childB2MML != nil && childB2MML.ID != nil {
				childID, err := NewEquipmentID(childB2MML.ID.Value)
				if err != nil {
					continue
				}

				child := NewEquipment(childID, childB2MML, class)
				_ = eq.AddChild(child)
			}
		}
	}

	return eq, nil
}

// NewEquipmentClassFromB2MML создаёт EquipmentClass из B2MML данных
func NewEquipmentClassFromB2MML(
	id EquipmentClassID,
	b2mmlData *b2mml.EquipmentClassType,
) (*EquipmentClass, error) {
	if b2mmlData == nil {
		return nil, errors.New("b2mml data cannot be nil")
	}

	ec := NewEquipmentClass(id, b2mmlData)

	// Преобразовать свойства класса
	if b2mmlData.EquipmentClassProperty != nil {
		for _, prop := range b2mmlData.EquipmentClassProperty {
			if prop != nil && prop.ID != nil {
				propID, err := NewEquipmentClassPropertyID(prop.ID.Value)
				if err != nil {
					continue
				}

				classProperty := NewEquipmentClassProperty(propID, prop)
				_ = ec.AddProperty(classProperty)
			}
		}
	}

	// Преобразовать дочерние классы
	if b2mmlData.EquipmentClassChild != nil {
		for _, childB2MML := range b2mmlData.EquipmentClassChild {
			if childB2MML != nil && childB2MML.ID != nil {
				childID, err := NewEquipmentClassID(childB2MML.ID.Value)
				if err != nil {
					continue
				}

				child, err := NewEquipmentClassFromB2MML(childID, childB2MML)
				if err != nil {
					continue
				}

				_ = ec.AddChild(child)
			}
		}
	}

	return ec, nil
}

// ToB2MML преобразует Equipment обратно в B2MML структуру
func (e *Equipment) ToB2MML() *b2mml.EquipmentType {
	if e == nil || e.data == nil {
		return nil
	}
	// Возвращаем оригинальные данные B2MML
	// которые синхронизируются через репозиторий
	return e.data
}

// ToB2MML преобразует EquipmentClass обратно в B2MML структуру
func (ec *EquipmentClass) ToB2MML() *b2mml.EquipmentClassType {
	if ec == nil || ec.data == nil {
		return nil
	}
	return ec.data
}
