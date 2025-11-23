package handler

import (
	"context"

	"github.com/grnsv/go-cmms/internal/app"
)

// Handler адаптирует ogen-сгенерированные интерфейсы к use cases.
// Полные сигнатуры методов, DTO и маршруты генерирует ogen из spec.yaml
type Handler struct {
	// Use cases
	listEquipmentUC    *app.ListEquipmentUseCase
	getEquipmentByIDUC *app.GetEquipmentByIDUseCase
	createEquipmentUC  *app.CreateEquipmentUseCase
}

// NewHandler создаёт новый handler с инъекцией use cases
func NewHandler(
	listEquipmentUC *app.ListEquipmentUseCase,
	getEquipmentByIDUC *app.GetEquipmentByIDUseCase,
	createEquipmentUC *app.CreateEquipmentUseCase,
) *Handler {
	return &Handler{
		listEquipmentUC:    listEquipmentUC,
		getEquipmentByIDUC: getEquipmentByIDUC,
		createEquipmentUC:  createEquipmentUC,
	}
}

// Методы адаптеров - должны соответствовать интерфейсам ogen.
// Точные сигнатуры, DTO и параметры маршрутов генерирует ogen из spec.yaml

// ListEquipment адаптирует запрос ogen к ListEquipmentUseCase
func (h *Handler) ListEquipment(ctx context.Context, limit *int, offset *int) (interface{}, error) {
	l := int32(10)
	if limit != nil && *limit > 0 {
		l = int32(*limit)
	}

	o := int32(0)
	if offset != nil && *offset >= 0 {
		o = int32(*offset)
	}

	result, err := h.listEquipmentUC.Execute(ctx, app.ListEquipmentInput{
		Limit:  l,
		Offset: o,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetEquipment адаптирует запрос ogen к GetEquipmentByIDUseCase
func (h *Handler) GetEquipment(ctx context.Context, id string) (interface{}, error) {
	result, err := h.getEquipmentByIDUC.Execute(ctx, app.GetEquipmentByIDInput{
		ExternalID: id,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreateEquipment адаптирует запрос ogen к CreateEquipmentUseCase
func (h *Handler) CreateEquipment(ctx context.Context, req interface{}) (interface{}, error) {
	// Ogen десериализует JSON в сгенерированный тип
	// Здесь нужно преобразовать к нашему CreateEquipmentInput
	// При реальной генерации ogen это будет типизировано
	result, err := h.createEquipmentUC.Execute(ctx, app.CreateEquipmentInput{
		// Поля будут заполнены из десериализованного req
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
