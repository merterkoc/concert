-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_events (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    user_id CHAR(36) NOT NULL,
    event_id VARCHAR(255) NOT NULL,
    status ENUM('interested', 'going') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_events;
-- +goose StatementEnd
