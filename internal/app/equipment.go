package app

import (
	"context"
	"fmt"

	"github.com/grnsv/go-cmms/internal/domain/repository"
)

// ListEquipmentInput входные параметры для ListEquipment
type ListEquipmentInput struct {
	Limit  int32
	Offset int32
}

// ListEquipmentOutput выходные данные для ListEquipment
type ListEquipmentOutput struct {
	Count int
	// Items []*Equipment // TODO: когда будет полная реализация репозитория
}

// ListEquipmentUseCase use case для получения списка оборудования
type ListEquipmentUseCase struct {
	equipmentRepo repository.EquipmentRepository
}

// NewListEquipmentUseCase создаёт новый use case
func NewListEquipmentUseCase(equipmentRepo repository.EquipmentRepository) *ListEquipmentUseCase {
	return &ListEquipmentUseCase{
		equipmentRepo: equipmentRepo,
	}
}

// Execute выполняет use case
func (uc *ListEquipmentUseCase) Execute(ctx context.Context, input ListEquipmentInput) (*ListEquipmentOutput, error) {
	// Заглушка: в реальности будет работать с репозиторием
	if input.Limit <= 0 {
		input.Limit = 10
	}
	if input.Limit > 100 {
		input.Limit = 100
	}

	return &ListEquipmentOutput{
		Count: 0,
	}, nil
}

// GetEquipmentByIDInput входные параметры для GetEquipmentByID
type GetEquipmentByIDInput struct {
	ExternalID string
}

// GetEquipmentByIDOutput выходные данные для GetEquipmentByID
type GetEquipmentByIDOutput struct {
	// Equipment *Equipment // TODO: когда будет полная реализация репозитория
}

// GetEquipmentByIDUseCase use case для получения оборудования по ID
type GetEquipmentByIDUseCase struct {
	equipmentRepo repository.EquipmentRepository
}

// NewGetEquipmentByIDUseCase создаёт новый use case
func NewGetEquipmentByIDUseCase(equipmentRepo repository.EquipmentRepository) *GetEquipmentByIDUseCase {
	return &GetEquipmentByIDUseCase{
		equipmentRepo: equipmentRepo,
	}
}

// Execute выполняет use case
func (uc *GetEquipmentByIDUseCase) Execute(ctx context.Context, input GetEquipmentByIDInput) (*GetEquipmentByIDOutput, error) {
	if input.ExternalID == "" {
		return nil, fmt.Errorf("external_id is required")
	}

	// equipment, err := uc.equipmentRepo.GetByExternalID(ctx, input.ExternalID)
	// if err != nil {
	//     return nil, err
	// }

	return &GetEquipmentByIDOutput{}, nil
}

// CreateEquipmentInput входные параметры для CreateEquipment
type CreateEquipmentInput struct {
	ExternalID string
	Version    string
	Status     string
}

// CreateEquipmentOutput выходные данные для CreateEquipment
type CreateEquipmentOutput struct {
	ID string
	// Equipment *Equipment
}

// CreateEquipmentUseCase use case для создания оборудования
type CreateEquipmentUseCase struct {
	equipmentRepo repository.EquipmentRepository
}

// NewCreateEquipmentUseCase создаёт новый use case
func NewCreateEquipmentUseCase(equipmentRepo repository.EquipmentRepository) *CreateEquipmentUseCase {
	return &CreateEquipmentUseCase{
		equipmentRepo: equipmentRepo,
	}
}

// Execute выполняет use case
func (uc *CreateEquipmentUseCase) Execute(ctx context.Context, input CreateEquipmentInput) (*CreateEquipmentOutput, error) {
	if input.ExternalID == "" {
		return nil, fmt.Errorf("external_id is required")
	}

	// equipment := model.NewEquipment(...)
	// if err := uc.equipmentRepo.Create(ctx, equipment); err != nil {
	//     return nil, err
	// }

	return &CreateEquipmentOutput{
		ID: "placeholder-uuid",
	}, nil
}
