-- +goose Up
ALTER TABLE environment
ADD app_id uuid NOT NULL REFERENCES app (id) ON DELETE CASCADE;

-- +goose Down
ALTER TABLE environment
DROP COLUMN app_id;
