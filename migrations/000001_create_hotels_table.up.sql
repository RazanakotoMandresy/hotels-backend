CREATE TABLE
    hotels (
        uuid UUID NOT NULL,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        services TEXT[],
        prix BIGINT,
        created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
        updated_at TIMESTAMP(0) WITHOUT TIME ZONE,
        deleted_at TIMESTAMP(0) WITHOUT TIME ZONE,
        PRIMARY KEY (uuid)
    );