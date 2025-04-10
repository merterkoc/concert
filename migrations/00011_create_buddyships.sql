-- +goose Up
-- +goose StatementBegin
CREATE TABLE buddyships (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    user1_id CHAR(36) NOT NULL,
    user2_id CHAR(36) NOT NULL,
    event_id CHAR(36) NOT NULL,
    matched_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user1 FOREIGN KEY (user1_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_user2 FOREIGN KEY (user2_id) REFERENCES users(id) ON DELETE CASCADE,

    CONSTRAINT no_duplicate_pair UNIQUE (user1_id, user2_id, event_id),
    CHECK (user1_id <> user2_id)
);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS buddyships;
-- +goose StatementEnd