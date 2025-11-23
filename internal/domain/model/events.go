package model

import "time"

// DomainEvent является базовым интерфейсом для всех доменных событий
type DomainEvent interface {
	AggregateID() string
	OccurredAt() time.Time
}

// EquipmentCreatedEvent возникает при создании нового оборудования
type EquipmentCreatedEvent struct {
	aggregateID string
	occurredAt  time.Time
	equipmentID EquipmentID
	classID     EquipmentClassID
}

func NewEquipmentCreatedEvent(equipmentID EquipmentID, classID EquipmentClassID) *EquipmentCreatedEvent {
	return &EquipmentCreatedEvent{
		aggregateID: equipmentID.String(),
		occurredAt:  time.Now(),
		equipmentID: equipmentID,
		classID:     classID,
	}
}

func (e *EquipmentCreatedEvent) AggregateID() string {
	return e.aggregateID
}

func (e *EquipmentCreatedEvent) OccurredAt() time.Time {
	return e.occurredAt
}

// EquipmentStatusChangedEvent возникает при изменении статуса оборудования
type EquipmentStatusChangedEvent struct {
	aggregateID string
	occurredAt  time.Time
	equipmentID EquipmentID
	oldStatus   OperatingStatus
	newStatus   OperatingStatus
}

func NewEquipmentStatusChangedEvent(equipmentID EquipmentID, oldStatus, newStatus OperatingStatus) *EquipmentStatusChangedEvent {
	return &EquipmentStatusChangedEvent{
		aggregateID: equipmentID.String(),
		occurredAt:  time.Now(),
		equipmentID: equipmentID,
		oldStatus:   oldStatus,
		newStatus:   newStatus,
	}
}

func (e *EquipmentStatusChangedEvent) AggregateID() string {
	return e.aggregateID
}

func (e *EquipmentStatusChangedEvent) OccurredAt() time.Time {
	return e.occurredAt
}

// EquipmentClassCreatedEvent возникает при создании нового класса оборудования
type EquipmentClassCreatedEvent struct {
	aggregateID string
	occurredAt  time.Time
	classID     EquipmentClassID
}

func NewEquipmentClassCreatedEvent(classID EquipmentClassID) *EquipmentClassCreatedEvent {
	return &EquipmentClassCreatedEvent{
		aggregateID: classID.String(),
		occurredAt:  time.Now(),
		classID:     classID,
	}
}

func (e *EquipmentClassCreatedEvent) AggregateID() string {
	return e.aggregateID
}

func (e *EquipmentClassCreatedEvent) OccurredAt() time.Time {
	return e.occurredAt
}
