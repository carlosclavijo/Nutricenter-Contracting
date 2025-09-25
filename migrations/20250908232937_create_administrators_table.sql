-- +goose Up
-- +goose StatementBegin
CREATE TABLE administrator
(
    id            UUID PRIMARY KEY,
    first_name    VARCHAR(100) NOT NULL,
    last_name     VARCHAR(100) NOT NULL,
    email         VARCHAR(200) NOT NULL UNIQUE,
    password      VARCHAR(255) NOT NULL,
    gender        CHAR(1)      NOT NULL DEFAULT 'U',
    birth         DATE,
    phone         VARCHAR(10),
    last_login_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    created_at    TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP    NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMP             DEFAULT NULL
);

CREATE TABLE patient
(
    id            UUID PRIMARY KEY,
    first_name    VARCHAR(100) NOT NULL,
    last_name     VARCHAR(100) NOT NULL,
    email         VARCHAR(200) NOT NULL UNIQUE,
    password      VARCHAR(255) NOT NULL,
    gender        CHAR(1)      NOT NULL DEFAULT 'U',
    birth         DATE,
    phone         VARCHAR(10),
    last_login_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    created_at    TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP    NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMP             DEFAULT NULL
);

CREATE TABLE contract
(
    id               UUID PRIMARY KEY,
    administrator_id UUID        NOT NULL REFERENCES administrator (id),
    patient_id       UUID        NOT NULL REFERENCES patient (id),
    contract_type    VARCHAR(10) NOT NULL,
    contract_status  VARCHAR(12) NOT NULL DEFAULT 'CREATED',
    creation_date    TIMESTAMP   NOT NULL DEFAULT NOW(),
    start_date       TIMESTAMP   NOT NULL,
    end_date         TIMESTAMP   NOT NULL,
    cost_value       INT         NOT NULL,
    created_at       TIMESTAMP   NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMP   NOT NULL DEFAULT NOW(),
    deleted_at       TIMESTAMP            DEFAULT NOW()
);

CREATE TABLE delivery
(
    id          UUID PRIMARY KEY,
    contract_id UUID      NOT NULL REFERENCES contract (id),
    date        TIMESTAMP NOT NULL,
    street      VARCHAR(50),
    number      INT,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMP          DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS delivery;
DROP TABLE IF EXISTS contract;
DROP TABLE IF EXISTS administrator;
DROP TABLE IF EXISTS patient;
-- +goose StatementEnd
