-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD user_image VARCHAR(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN user_image;
-- +goose StatementEnd
