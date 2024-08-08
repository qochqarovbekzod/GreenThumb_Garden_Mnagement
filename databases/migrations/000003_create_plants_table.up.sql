CREATE TABLE IF NOT EXISTS plants (
    id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY,
    garden_id UUID NOT NULL REFERENCES gardens(id),
    species VARCHAR(100) NOT NULL,
    quantity INTEGER,
    planting_date DATE,
    status plant_status DEFAULT 'planned',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);