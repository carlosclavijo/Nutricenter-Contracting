-- +goose Up
-- +goose StatementBegin
CREATE TABLE administrator
(
    id            UUID PRIMARY KEY,
    first_name    VARCHAR(100) NOT NULL,
    last_name     VARCHAR(100) NOT NULL,
    email         VARCHAR(200) NOT NULL UNIQUE,
    password      VARCHAR(255) NOT NULL,
    gender        CHAR(1)      NOT NULL DEFAULT 'U' CHECK (gender IN ('U', 'M', 'F')),
    birth         DATE,
    phone         VARCHAR(10),
    last_login_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    created_at    TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP    NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMP             DEFAULT NULL
);
-- Type Gender U = Unknown, M = Male, F = Female

CREATE TABLE patient
(
    id            UUID PRIMARY KEY,
    first_name    VARCHAR(100) NOT NULL,
    last_name     VARCHAR(100) NOT NULL,
    email         VARCHAR(200) NOT NULL UNIQUE,
    password      VARCHAR(255) NOT NULL,
    gender        CHAR(1)      NOT NULL DEFAULT 'U' CHECK (gender IN ('U', 'M', 'F')),
    birth         DATE,
    phone         VARCHAR(10),
    last_login_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    created_at    TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP    NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMP             DEFAULT NULL
);
-- Type Gender U = Unknown, M = Male, F = Female

CREATE TABLE contract
(
    id               UUID PRIMARY KEY,
    administrator_id UUID      NOT NULL REFERENCES administrator (id),
    patient_id       UUID      NOT NULL REFERENCES patient (id),
    type             CHAR(1)   NOT NULL CHECK (type IN ('H', 'M')),
    status           CHAR(1)   NOT NULL DEFAULT 'C' CHECK (status IN ('C', 'A', 'F')),
    creation         TIMESTAMP NOT NULL DEFAULT NOW(),
    start            TIMESTAMP NOT NULL,
    finalized        TIMESTAMP NOT NULL,
    cost             INT       NOT NULL,
    created_at       TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at       TIMESTAMP          DEFAULT NULL
);
-- Type H = Home, M = Monthly
-- Status C = Created, A = Active, F = Finalized

CREATE TABLE delivery
(
    id          UUID PRIMARY KEY,
    contract_id UUID             NOT NULL REFERENCES contract (id),
    date        TIMESTAMP        NOT NULL,
    street      VARCHAR(50)      NOT NULL,
    number      INT              NOT NULL,
    latitude    DOUBLE PRECISION NOT NULL,
    longitude   DOUBLE PRECISION NOT NULL,
    status      CHAR(1)          NOT NULL DEFAULT 'P' CHECK (status IN ('P', 'D', 'C')),
    created_at  TIMESTAMP        NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP        NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMP                 DEFAULT NULL
);
-- Type Status P = Pending, D = Delivered, C = Canceled

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS delivery;
DROP TABLE IF EXISTS contract;
DROP TABLE IF EXISTS administrator;
DROP TABLE IF EXISTS patient;
-- +goose StatementEnd

-- goose postgres "user=postgres password=abc12345 dbname=nutricenter-contracting sslmode=disable" down
-- goose postgres "user=postgres password=abc12345 dbname=nutricenter-contracting sslmode=disable" up