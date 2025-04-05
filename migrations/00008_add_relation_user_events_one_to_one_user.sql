-- +goose Up
-- +goose StatementBegin
ALTER TABLE user_events
    ADD CONSTRAINT fk_user_events_user
        FOREIGN KEY (user_id) REFERENCES users(id)
            ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX idx_user_events_user_id ON user_events(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE user_events
DROP FOREIGN KEY fk_user_events_user;
-- +goose StatementEnd

-- +goose StatementBegin
DROP INDEX idx_user_events_user_id ON user_events;
-- +goose StatementEnd