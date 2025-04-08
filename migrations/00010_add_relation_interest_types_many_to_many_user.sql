-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_interest_types
(
    user_id          CHAR(36) NOT NULL,
    interest_type_id BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (user_id, interest_type_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (interest_type_id) REFERENCES interest_types (id) ON DELETE CASCADE
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE user_interest_types;
-- +goose StatementEnd