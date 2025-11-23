-- Equipment Classes queries

-- name: CreateEquipmentClass :one
INSERT INTO equipment_classes (
    external_id,
    version,
    description,
    published_date,
    effective_start_date,
    effective_end_date,
    hierarchy_scope_id,
    equipment_level,
    parent_class_id,
    b2mml_data
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING
    id,
    external_id,
    version,
    description,
    published_date,
    effective_start_date,
    effective_end_date,
    hierarchy_scope_id,
    equipment_level,
    parent_class_id,
    b2mml_data,
    created_at,
    updated_at,
    deleted_at;

-- name: GetEquipmentClassByID :one
SELECT
    id,
    external_id,
    version,
    description,
    published_date,
    effective_start_date,
    effective_end_date,
    hierarchy_scope_id,
    equipment_level,
    parent_class_id,
    b2mml_data,
    created_at,
    updated_at,
    deleted_at
FROM equipment_classes
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetEquipmentClassByExternalID :one
SELECT
    id,
    external_id,
    version,
    description,
    published_date,
    effective_start_date,
    effective_end_date,
    hierarchy_scope_id,
    equipment_level,
    parent_class_id,
    b2mml_data,
    created_at,
    updated_at,
    deleted_at
FROM equipment_classes
WHERE external_id = $1 AND deleted_at IS NULL;

-- name: ListAllEquipmentClasses :many
SELECT
    id,
    external_id,
    version,
    description,
    published_date,
    effective_start_date,
    effective_end_date,
    hierarchy_scope_id,
    equipment_level,
    parent_class_id,
    b2mml_data,
    created_at,
    updated_at,
    deleted_at
FROM equipment_classes
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdateEquipmentClass :one
UPDATE equipment_classes
SET
    version = $2,
    description = $3,
    published_date = $4,
    effective_start_date = $5,
    effective_end_date = $6,
    hierarchy_scope_id = $7,
    equipment_level = $8,
    b2mml_data = $9,
    updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING
    id,
    external_id,
    version,
    description,
    published_date,
    effective_start_date,
    effective_end_date,
    hierarchy_scope_id,
    equipment_level,
    parent_class_id,
    b2mml_data,
    created_at,
    updated_at,
    deleted_at;

-- name: DeleteEquipmentClass :exec
UPDATE equipment_classes
SET deleted_at = NOW()
WHERE id = $1;

-- Equipment Class Properties queries

-- name: CreateEquipmentClassProperty :one
INSERT INTO equipment_class_properties (
    equipment_class_id,
    external_id,
    description,
    property_type,
    parent_property_id,
    b2mml_data
) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
    id,
    equipment_class_id,
    external_id,
    description,
    property_type,
    parent_property_id,
    b2mml_data,
    created_at,
    updated_at;

-- name: ListEquipmentClassProperties :many
SELECT
    id,
    equipment_class_id,
    external_id,
    description,
    property_type,
    parent_property_id,
    b2mml_data,
    created_at,
    updated_at
FROM equipment_class_properties
WHERE equipment_class_id = $1
ORDER BY created_at;

-- Equipment queries

-- name: CreateEquipment :one
INSERT INTO equipment (
    external_id,
    version,
    description,
    published_date,
    effective_start_date,
    effective_end_date,
    hierarchy_scope_id,
    equipment_level,
    operating_status,
    physical_asset_id,
    operational_location_id,
    parent_equipment_id,
    b2mml_data
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
)
RETURNING
    id,
    external_id,
    version,
    description,
    published_date,
    effective_start_date,
    effective_end_date,
    hierarchy_scope_id,
    equipment_level,
    operating_status,
    physical_asset_id,
    operational_location_id,
    parent_equipment_id,
    b2mml_data,
    created_at,
    updated_at,
    deleted_at,
    record_version;

-- name: GetEquipmentByID :one
SELECT
    id,
    external_id,
    version,
    description,
    published_date,
    effective_start_date,
    effective_end_date,
    hierarchy_scope_id,
    equipment_level,
    operating_status,
    physical_asset_id,
    operational_location_id,
    parent_equipment_id,
    b2mml_data,
    created_at,
    updated_at,
    deleted_at,
    record_version
FROM equipment
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetEquipmentByExternalID :one
SELECT
    id,
    external_id,
    version,
    description,
    published_date,
    effective_start_date,
    effective_end_date,
    hierarchy_scope_id,
    equipment_level,
    operating_status,
    physical_asset_id,
    operational_location_id,
    parent_equipment_id,
    b2mml_data,
    created_at,
    updated_at,
    deleted_at,
    record_version
FROM equipment
WHERE external_id = $1 AND deleted_at IS NULL;

-- name: ListEquipment :many
SELECT
    id,
    external_id,
    version,
    description,
    published_date,
    effective_start_date,
    effective_end_date,
    hierarchy_scope_id,
    equipment_level,
    operating_status,
    physical_asset_id,
    operational_location_id,
    parent_equipment_id,
    b2mml_data,
    created_at,
    updated_at,
    deleted_at,
    record_version
FROM equipment
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListEquipmentByStatus :many
SELECT
    id,
    external_id,
    version,
    description,
    published_date,
    effective_start_date,
    effective_end_date,
    hierarchy_scope_id,
    equipment_level,
    operating_status,
    physical_asset_id,
    operational_location_id,
    parent_equipment_id,
    b2mml_data,
    created_at,
    updated_at,
    deleted_at,
    record_version
FROM equipment
WHERE operating_status = $1 AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListChildEquipment :many
SELECT
    id,
    external_id,
    version,
    description,
    published_date,
    effective_start_date,
    effective_end_date,
    hierarchy_scope_id,
    equipment_level,
    operating_status,
    physical_asset_id,
    operational_location_id,
    parent_equipment_id,
    b2mml_data,
    created_at,
    updated_at,
    deleted_at,
    record_version
FROM equipment
WHERE parent_equipment_id = $1 AND deleted_at IS NULL
ORDER BY created_at;

-- name: UpdateEquipmentStatus :one
UPDATE equipment
SET
    operating_status = $2,
    updated_at = NOW(),
    record_version = record_version + 1
WHERE id = $1 AND deleted_at IS NULL
RETURNING
    id,
    external_id,
    version,
    description,
    published_date,
    effective_start_date,
    effective_end_date,
    hierarchy_scope_id,
    equipment_level,
    operating_status,
    physical_asset_id,
    operational_location_id,
    parent_equipment_id,
    b2mml_data,
    created_at,
    updated_at,
    deleted_at,
    record_version;

-- name: DeleteEquipment :exec
UPDATE equipment
SET deleted_at = NOW()
WHERE id = $1;

-- Equipment Properties queries

-- name: CreateEquipmentProperty :one
INSERT INTO equipment_properties (
    equipment_id,
    external_id,
    class_property_id,
    property_value,
    property_data_type,
    property_unit,
    description,
    b2mml_data
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING
    id,
    equipment_id,
    external_id,
    class_property_id,
    property_value,
    property_data_type,
    property_unit,
    description,
    b2mml_data,
    created_at,
    updated_at;

-- name: ListEquipmentProperties :many
SELECT
    id,
    equipment_id,
    external_id,
    class_property_id,
    property_value,
    property_data_type,
    property_unit,
    description,
    b2mml_data,
    created_at,
    updated_at
FROM equipment_properties
WHERE equipment_id = $1
ORDER BY created_at;

-- name: UpdateEquipmentProperty :one
UPDATE equipment_properties
SET
    property_value = $2,
    property_data_type = $3,
    property_unit = $4,
    description = $5,
    b2mml_data = $6,
    updated_at = NOW()
WHERE id = $1
RETURNING
    id,
    equipment_id,
    external_id,
    class_property_id,
    property_value,
    property_data_type,
    property_unit,
    description,
    b2mml_data,
    created_at,
    updated_at;

-- name: DeleteEquipmentProperty :exec
DELETE FROM equipment_properties
WHERE id = $1;

-- Equipment Class Mappings queries

-- name: AddEquipmentToClass :one
INSERT INTO equipment_class_mappings (
    equipment_id,
    equipment_class_id
) VALUES ($1, $2)
RETURNING
    id,
    equipment_id,
    equipment_class_id,
    created_at;

-- name: RemoveEquipmentFromClass :exec
DELETE FROM equipment_class_mappings
WHERE equipment_id = $1 AND equipment_class_id = $2;

-- name: ListEquipmentClassesForEquipment :many
SELECT
    ec.id,
    ec.external_id,
    ec.version,
    ec.description,
    ec.published_date,
    ec.effective_start_date,
    ec.effective_end_date,
    ec.hierarchy_scope_id,
    ec.equipment_level,
    ec.parent_class_id,
    ec.b2mml_data,
    ec.created_at,
    ec.updated_at,
    ec.deleted_at
FROM equipment_class_mappings ecm
JOIN equipment_classes ec ON ecm.equipment_class_id = ec.id
WHERE ecm.equipment_id = $1 AND ec.deleted_at IS NULL
ORDER BY ec.created_at;

-- name: ListEquipmentByClass :many
SELECT
    e.id,
    e.external_id,
    e.version,
    e.description,
    e.published_date,
    e.effective_start_date,
    e.effective_end_date,
    e.hierarchy_scope_id,
    e.equipment_level,
    e.operating_status,
    e.physical_asset_id,
    e.operational_location_id,
    e.parent_equipment_id,
    e.b2mml_data,
    e.created_at,
    e.updated_at,
    e.deleted_at,
    e.record_version
FROM equipment e
JOIN equipment_class_mappings ecm ON e.id = ecm.equipment_id
WHERE ecm.equipment_class_id = $1 AND e.deleted_at IS NULL
ORDER BY e.created_at DESC;
