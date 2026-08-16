-- +goose Up
CREATE TABLE app (
    id         uuid         PRIMARY KEY DEFAULT uuidv7(),
    name       varchar(128) NOT NULL,
    created_at timestamptz  NOT NULL DEFAULT now(),
    updated_at timestamptz  NOT NULL DEFAULT now(),
    deleted_at timestamptz
);
CREATE UNIQUE INDEX ux_app_name ON app (name);

CREATE TABLE environment (
    id            uuid        PRIMARY KEY DEFAULT uuidv7(),
    name          varchar(64) NOT NULL,
    created_at    timestamptz NOT NULL DEFAULT now(),
    updated_at    timestamptz NOT NULL DEFAULT now(),
    deleted_at    timestamptz
);

CREATE TABLE app_version (
    id           uuid        PRIMARY KEY DEFAULT uuidv7(),
    app_id       uuid        NOT NULL REFERENCES app (id) ON DELETE CASCADE,
    version      varchar(12) NOT NULL,
    released_at  timestamptz,
    created_at   timestamptz NOT NULL DEFAULT now(),
    updated_at   timestamptz NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX ux_app_version_version ON app_version (app_id, version);
CREATE UNIQUE INDEX ux_app_version ON app_version (version);

CREATE TABLE feature_flag (
    id          uuid         PRIMARY KEY DEFAULT uuidv7(),
    app_id      uuid         NOT NULL REFERENCES app (id) ON DELETE CASCADE,
    name        varchar(256) NOT NULL,
    description text         NULL,
    tags        jsonb        NULL,
    created_at  timestamptz  NOT NULL DEFAULT now(),
    updated_at  timestamptz  NOT NULL DEFAULT now(),
    deleted_at  timestamptz
);
CREATE UNIQUE INDEX ux_feature_flag_app_key ON feature_flag (app_id, name) WHERE deleted_at IS NULL;
CREATE INDEX ix_feature_flag_tags ON feature_flag USING gin (tags);

CREATE TABLE feature_flag_setting (
    id                 uuid        PRIMARY KEY DEFAULT uuidv7(),
    feature_flag_id    uuid        NOT NULL REFERENCES feature_flag (id) ON DELETE CASCADE,
    environment_id     uuid        NOT NULL REFERENCES environment (id) ON DELETE RESTRICT,
    active             boolean     NOT NULL DEFAULT false,
    min_app_version_id uuid REFERENCES app_version (id) ON DELETE SET NULL,
    max_app_version_id uuid REFERENCES app_version (id) ON DELETE SET NULL,
    created_at         timestamptz NOT NULL DEFAULT now(),
    updated_at         timestamptz NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX ux_ff_setting_flag_env ON feature_flag_setting (feature_flag_id, environment_id);
CREATE INDEX ix_ff_setting_env_active ON feature_flag_setting (environment_id, active);

CREATE TABLE feature_flag_override (
    id                      uuid         PRIMARY KEY DEFAULT uuidv7(),
    feature_flag_setting_id uuid         NOT NULL REFERENCES feature_flag_setting (id) ON DELETE CASCADE,
    subject_value           varchar(256) NOT NULL,
    note                    text         NULL,
    expires_at              timestamptz,
    created_at              timestamptz  NOT NULL DEFAULT now(),
    updated_at              timestamptz  NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX ux_ff_override_subject
    ON feature_flag_override (subject_value);
CREATE INDEX ix_ff_override_lookup ON feature_flag_override (subject_value);

CREATE TABLE audit_log (
    id          uuid         PRIMARY KEY DEFAULT uuidv7(),
    request_id  varchar(64)  NULL,
    reason      text         NULL,
    created_at  timestamptz  NOT NULL DEFAULT now()
);
CREATE INDEX ix_audit_log_created_at ON audit_log (created_at DESC);
CREATE INDEX ix_audit_log_request_id ON audit_log (request_id);
CREATE INDEX ix_audit_log_reason ON audit_log (reason);

-- +goose Down
DROP TABLE IF EXISTS audit_log;
DROP TABLE IF EXISTS feature_flag_override;
DROP TABLE IF EXISTS feature_flag_setting;
DROP TABLE IF EXISTS feature_flag;
DROP TABLE IF EXISTS app_version;
DROP TABLE IF EXISTS environment;
DROP TABLE IF EXISTS app;
