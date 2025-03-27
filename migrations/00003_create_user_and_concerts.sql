-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_concerts (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    user_id CHAR(36) NOT NULL,
    concert_id CHAR(36) NOT NULL,
    status VARCHAR(20) CHECK (status IN ('interested', 'going')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (concert_id) REFERENCES concerts(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_concerts;
-- +goose StatementEnd
