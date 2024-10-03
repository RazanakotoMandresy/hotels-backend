CREATE TABLE
    hotels (
        uuid UUID NOT NULL,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        services TEXT[],
        images TEXT[],
        prix BIGINT,
        status SMALLINT,
        ouverture TEXT NOT NULL,
        created_by TEXT NOT NULL,
        created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
        updated_at TIMESTAMP(0) WITHOUT TIME ZONE,
        deleted_at TIMESTAMP(0) WITHOUT TIME ZONE,
        PRIMARY KEY (uuid)
    );