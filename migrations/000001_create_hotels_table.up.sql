CREATE TABLE
    hotels (
        uuid UUID NOT NULL,
        name TEXT NOT NULL,
        place TEXT NOT NULL,
        descriptions TEXT NOT NULL,
        created_by TEXT NOT NULL,
        services TEXT[],
        images TEXT[],
        reservation_list TEXT[],
        prix BIGINT,
        status BOOLEAN,
        created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
        updated_at TIMESTAMP(0) WITHOUT TIME ZONE,
        deleted_at TIMESTAMP(0) WITHOUT TIME ZONE,
        PRIMARY KEY (uuid)
    );