package model

import "errors"

var (
	// Equipment errors
	ErrEquipmentIDEmpty           = errors.New("equipment id cannot be empty")
	ErrEquipmentClassIDEmpty      = errors.New("equipment class id cannot be empty")
	ErrEquipmentPropertyIDEmpty   = errors.New("equipment property id cannot be empty")
	ErrEquipmentClassPropertyIDEmpty = errors.New("equipment class property id cannot be empty")

	// Equipment validation errors
	ErrEquipmentNotFound       = errors.New("equipment not found")
	ErrEquipmentAlreadyExists  = errors.New("equipment already exists")
	ErrEquipmentInvalidStatus  = errors.New("invalid equipment status")

	// EquipmentClass errors
	ErrEquipmentClassNotFound      = errors.New("equipment class not found")
	ErrEquipmentClassAlreadyExists = errors.New("equipment class already exists")
)
