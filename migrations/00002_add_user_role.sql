-- +goose Up
-- +goose StatementBegin
CREATE TYPE user_role AS ENUM('admin', 'seller', 'client');
ALTER TABLE users
ADD COLUMN role user_role NOT NULL DEFAULT 'client';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN role;
DROP TYPE user_role;
-- +goose StatementEnd
