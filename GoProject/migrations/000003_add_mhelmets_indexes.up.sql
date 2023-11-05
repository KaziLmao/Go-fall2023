CREATE INDEX IF NOT EXISTS mhelmets_name_idx ON mhelmets USING GIN (to_tsvector('simple', name));
CREATE INDEX IF NOT EXISTS mhelmets_material_idx ON mhelmets USING GIN (to_tsvector('simple', material));
CREATE INDEX IF NOT EXISTS mhelmets_protection_idx ON mhelmets USING GIN (to_tsvector('simple', protection));