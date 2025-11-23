package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/grnsv/go-cmms/internal/domain/model"
)

// EquipmentRepository интерфейс репозитория для Equipment
type EquipmentRepository interface {
	// Create сохраняет новое оборудование
	Create(ctx context.Context, equipment *model.Equipment) error

	// GetByID получает оборудование по UUID
	GetByID(ctx context.Context, id uuid.UUID) (*model.Equipment, error)

	// GetByExternalID получает оборудование по внешнему ID
	GetByExternalID(ctx context.Context, externalID string) (*model.Equipment, error)

	// List получает список оборудования с пагинацией
	List(ctx context.Context, limit int32, offset int32) ([]*model.Equipment, error)

	// ListByStatus получает список оборудования по статусу
	ListByStatus(ctx context.Context, status model.OperatingStatus, limit int32, offset int32) ([]*model.Equipment, error)

	// Update обновляет оборудование
	Update(ctx context.Context, equipment *model.Equipment) error

	// Delete удаляет оборудование (soft delete)
	Delete(ctx context.Context, id uuid.UUID) error
}

// EquipmentClassRepository интерфейс репозитория для EquipmentClass
type EquipmentClassRepository interface {
	// Create сохраняет новый класс оборудования
	Create(ctx context.Context, class *model.EquipmentClass) error

	// GetByID получает класс по UUID
	GetByID(ctx context.Context, id uuid.UUID) (*model.EquipmentClass, error)

	// GetByExternalID получает класс по внешнему ID
	GetByExternalID(ctx context.Context, externalID string) (*model.EquipmentClass, error)

	// List получает список классов с пагинацией
	List(ctx context.Context, limit int32, offset int32) ([]*model.EquipmentClass, error)

	// Update обновляет класс
	Update(ctx context.Context, class *model.EquipmentClass) error

	// Delete удаляет класс (soft delete)
	Delete(ctx context.Context, id uuid.UUID) error
}

// UnitOfWork паттерн для управления транзакциями
type UnitOfWork interface {
	// Equipment возвращает репозиторий Equipment
	Equipment() EquipmentRepository

	// EquipmentClass возвращает репозиторий EquipmentClass
	EquipmentClass() EquipmentClassRepository

	// Begin начинает транзакцию
	Begin(ctx context.Context) (UnitOfWork, error)

	// Commit коммитит транзакцию
	Commit(ctx context.Context) error

	// Rollback откатывает транзакцию
	Rollback(ctx context.Context) error
}
