CREATE SCHEMA IF NOT EXISTS audit;
CREATE SCHEMA IF NOT EXISTS auth;
CREATE SCHEMA IF NOT EXISTS app;

CREATE TABLE IF NOT EXISTS auth.app_user_roles (
                                                   id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                                   role_name VARCHAR(255) NOT NULL UNIQUE,
                                                   description TEXT NOT NULL ,
                                                   created_at BIGINT NOT NULL DEFAULT ((extract(epoch from now()) * 1000)::bigint),
                                                   updated_at BIGINT NOT NULL DEFAULT ((extract(epoch from now()) * 1000)::bigint),
                                                   created_by BIGINT NOT NULL DEFAULT 1,
                                                   updated_by BIGINT NOT NULL DEFAULT 1
);


CREATE TABLE IF NOT EXISTS auth.app_users (
                                              id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                              first_name VARCHAR(255) NOT NULL ,
                                              last_name VARCHAR(255) NULL,
                                              mobile_no VARCHAR(255) NOT NULL ,
                                              preferred_user_name VARCHAR(1000) NOT NULL ,
                                              email VARCHAR(1000) NOT NULL UNIQUE,
                                              role_id BIGINT NOT NULL REFERENCES auth.app_user_roles(id),
                                              is_active BOOLEAN NOT NULL DEFAULT TRUE,
                                              is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
                                              fcm_token TEXT NULL,
                                              created_at BIGINT NOT NULL DEFAULT ((extract(epoch from now()) * 1000)::bigint),
                                              updated_at BIGINT NOT NULL DEFAULT ((extract(epoch from now()) * 1000)::bigint),
                                              created_by BIGINT NOT NULL DEFAULT 1,
                                              updated_by BIGINT NOT NULL DEFAULT 1,
                                              last_active_date BIGINT NULL
);

-- Helpful indexes for app_users
CREATE INDEX IF NOT EXISTS idx_app_users_mobile ON auth.app_users(mobile_no);
CREATE INDEX IF NOT EXISTS idx_app_users_role_id ON auth.app_users(role_id);
CREATE INDEX IF NOT EXISTS idx_app_users_email ON auth.app_users(email);
CREATE INDEX IF NOT EXISTS idx_app_users_preferred_user_name ON auth.app_users(preferred_user_name);

CREATE OR REPLACE FUNCTION auth.update_unix_milli_generic()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = (extract(epoch from now()) * 1000)::bigint;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger for app_users
CREATE TRIGGER trg_update_app_users_timestamp
    BEFORE UPDATE ON auth.app_users
    FOR EACH ROW
EXECUTE FUNCTION auth.update_unix_milli_generic();

-- Trigger for app_user_roles
CREATE TRIGGER trg_update_app_user_roles_timestamp
    BEFORE UPDATE ON auth.app_user_roles
    FOR EACH ROW
EXECUTE FUNCTION auth.update_unix_milli_generic();
