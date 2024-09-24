CREATE TABLE users (
        id SERIAL,
        name  TEXT NOT NULL UNIQUE,
        mail  TEXT NOT NULL UNIQUE,
        list_hotels TEXT[],
        created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
        updated_at TIMESTAMP(0) WITHOUT TIME ZONE,
        deleted_at TIMESTAMP(0) WITHOUT TIME ZONE,
        PRIMARY KEY (id)
    ) 