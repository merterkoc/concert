-- +goose Up
-- +goose StatementBegin
CREATE TABLE interest_types
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO interest_types (name)
VALUES ('science'),
       ('art'),
       ('music'),
       ('theatre'),
       ('sports'),
       ('tech');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE interest_types;
-- +goose StatementEnd