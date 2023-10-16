CREATE TABLE IF NOT EXISTS mhelmets (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    year integer NOT NULL,
    material text NOT NULL,
    ventilation boolean NOT NULL,
    protection text NOT NULL,
    weight integer NOT NULL,
    sun_protection boolean NOT NULL,
    lining text NOT NULL,
    fastening text NOT NULL
)