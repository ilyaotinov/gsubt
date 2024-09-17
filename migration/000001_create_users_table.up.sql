BEGIN;
CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language plpgsql;

CREATE TYPE user_role AS ENUM ('admin', 'teacher', 'children');
CREATE TABLE IF NOT EXISTS users
(
    id         BIGSERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL DEFAULT '',
    surname    VARCHAR(255) NOT NULL DEFAULT '',
    login      VARCHAR(255) NOT NULL UNIQUE,
    email      VARCHAR(255) NOT NULL DEFAULT '',
    role       user_role    NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE
    update_updated_at_column();
COMMIT;