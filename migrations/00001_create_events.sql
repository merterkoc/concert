-- +goose Up
-- +goose StatementBegin
CREATE TABLE events
(
    id         CHAR(36) PRIMARY KEY,
    title      VARCHAR(255) NOT NULL,
    artist     VARCHAR(255) NOT NULL,
    venue      VARCHAR(255) NOT NULL,
    date_time  DATETIME     NOT NULL,
    genre      VARCHAR(100),
    location   POINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events;
-- +goose StatementEnd
