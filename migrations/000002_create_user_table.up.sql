CREATE TABLE
    user (
        id SERIAL,
        name UNIQUE TEXT NOT NULL,
        mail UNIQUE TEXT NOT NULL,
        list_hotels TEXT[],
        created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
        updated_at TIMESTAMP(0) WITHOUT TIME ZONE,
        deleted_at TIMESTAMP(0) WITHOUT TIME ZONE,
        PRIMARY KEY (id)
    )