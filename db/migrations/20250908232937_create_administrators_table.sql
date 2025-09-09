-- +goose Up
-- +goose StatementBegin
CREATE TABLE administrators(
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phone VARCHAR(10) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE administrators;
-- +goose StatementEnd
