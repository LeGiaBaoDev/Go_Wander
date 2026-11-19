-- +goose Up
-- Bootstrap R1: extensions, module schemas and database roles.
-- Login credentials for gowander_owner/gowander_app are provisioned outside
-- this migration (environment/secret manager), never in git.

CREATE EXTENSION IF NOT EXISTS postgis;
CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE EXTENSION IF NOT EXISTS citext;

-- +goose StatementBegin
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'gowander_owner') THEN
        CREATE ROLE gowander_owner NOLOGIN;
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'gowander_app') THEN
        CREATE ROLE gowander_app NOLOGIN;
    END IF;
END
$$;
-- +goose StatementEnd

CREATE SCHEMA IF NOT EXISTS catalog AUTHORIZATION gowander_owner;
CREATE SCHEMA IF NOT EXISTS media AUTHORIZATION gowander_owner;
CREATE SCHEMA IF NOT EXISTS mobility AUTHORIZATION gowander_owner;
CREATE SCHEMA IF NOT EXISTS identity AUTHORIZATION gowander_owner;
CREATE SCHEMA IF NOT EXISTS engagement AUTHORIZATION gowander_owner;
CREATE SCHEMA IF NOT EXISTS backoffice AUTHORIZATION gowander_owner;

GRANT USAGE ON SCHEMA catalog, media, mobility, identity, engagement, backoffice TO gowander_app;

-- +goose Down
DROP SCHEMA IF EXISTS backoffice CASCADE;
DROP SCHEMA IF EXISTS engagement CASCADE;
DROP SCHEMA IF EXISTS identity CASCADE;
DROP SCHEMA IF EXISTS mobility CASCADE;
DROP SCHEMA IF EXISTS media CASCADE;
DROP SCHEMA IF EXISTS catalog CASCADE;

DROP ROLE IF EXISTS gowander_app;
DROP ROLE IF EXISTS gowander_owner;

DROP EXTENSION IF EXISTS postgis CASCADE;
DROP EXTENSION IF EXISTS pg_trgm;
DROP EXTENSION IF EXISTS citext;
