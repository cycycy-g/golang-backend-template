CREATE TABLE IF NOT EXISTS users
(
    id
    bigserial
    PRIMARY
    KEY,
    username
    varchar
    NOT
    NULL
    UNIQUE,
    hashed_password
    varchar
    NOT
    NULL,
    email
    varchar
    NOT
    NULL
    UNIQUE,
    full_name
    varchar
    NOT
    NULL,
    created_at
    timestamptz
    NOT
    NULL
    DEFAULT
    CURRENT_TIMESTAMP,
    updated_at
    timestamptz
    NOT
    NULL
    DEFAULT
    CURRENT_TIMESTAMP
);

CREATE INDEX ON users (username);
CREATE INDEX ON users (email);