-- Create equipment_classes table
CREATE TABLE equipment_classes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    external_id VARCHAR(255) NOT NULL UNIQUE,
    version VARCHAR(50),
    description TEXT,
    published_date TIMESTAMPTZ,
    effective_start_date TIMESTAMPTZ,
    effective_end_date TIMESTAMPTZ,
    hierarchy_scope_id VARCHAR(255),
    equipment_level VARCHAR(255),
    parent_class_id UUID REFERENCES equipment_classes(id) ON DELETE SET NULL,
    b2mml_data JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_equipment_classes_external_id ON equipment_classes(external_id);
CREATE INDEX idx_equipment_classes_parent_class_id ON equipment_classes(parent_class_id);
CREATE INDEX idx_equipment_classes_deleted_at ON equipment_classes(deleted_at);

-- Create equipment_class_properties table
CREATE TABLE equipment_class_properties (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    equipment_class_id UUID NOT NULL REFERENCES equipment_classes(id) ON DELETE CASCADE,
    external_id VARCHAR(255) NOT NULL,
    description TEXT,
    property_type VARCHAR(255),
    parent_property_id UUID REFERENCES equipment_class_properties(id) ON DELETE SET NULL,
    b2mml_data JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(equipment_class_id, external_id)
);

CREATE INDEX idx_equipment_class_properties_class_id ON equipment_class_properties(equipment_class_id);
CREATE INDEX idx_equipment_class_properties_parent_id ON equipment_class_properties(parent_property_id);

-- Create equipment table
CREATE TABLE equipment (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    external_id VARCHAR(255) NOT NULL UNIQUE,
    version VARCHAR(50),
    description TEXT,
    published_date TIMESTAMPTZ,
    effective_start_date TIMESTAMPTZ,
    effective_end_date TIMESTAMPTZ,
    hierarchy_scope_id VARCHAR(255),
    equipment_level VARCHAR(255),
    operating_status VARCHAR(50) DEFAULT 'inactive',
    physical_asset_id VARCHAR(255),
    operational_location_id VARCHAR(255),
    parent_equipment_id UUID REFERENCES equipment(id) ON DELETE SET NULL,
    b2mml_data JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    record_version BIGINT NOT NULL DEFAULT 1
);

CREATE INDEX idx_equipment_external_id ON equipment(external_id);
CREATE INDEX idx_equipment_parent_id ON equipment(parent_equipment_id);
CREATE INDEX idx_equipment_status ON equipment(operating_status);
CREATE INDEX idx_equipment_deleted_at ON equipment(deleted_at);
CREATE INDEX idx_equipment_created_at ON equipment(created_at);

-- Create equipment_properties table
CREATE TABLE equipment_properties (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    equipment_id UUID NOT NULL REFERENCES equipment(id) ON DELETE CASCADE,
    external_id VARCHAR(255) NOT NULL,
    class_property_id UUID REFERENCES equipment_class_properties(id),
    property_value TEXT,
    property_data_type VARCHAR(50),
    property_unit VARCHAR(50),
    description TEXT,
    b2mml_data JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(equipment_id, external_id)
);

CREATE INDEX idx_equipment_properties_equipment_id ON equipment_properties(equipment_id);
CREATE INDEX idx_equipment_properties_class_property_id ON equipment_properties(class_property_id);

-- Create equipment_class_mappings table (связь много-ко-многим между equipment и equipment_classes)
CREATE TABLE equipment_class_mappings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    equipment_id UUID NOT NULL REFERENCES equipment(id) ON DELETE CASCADE,
    equipment_class_id UUID NOT NULL REFERENCES equipment_classes(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(equipment_id, equipment_class_id)
);

CREATE INDEX idx_equipment_class_mappings_equipment_id ON equipment_class_mappings(equipment_id);
CREATE INDEX idx_equipment_class_mappings_class_id ON equipment_class_mappings(equipment_class_id);
