CREATE TABLE
    hotels (
        uuid UUID NOT NULL,
        name TEXT NOT NULL,
        place TEXT NOT NULL,
        description TEXT NOT NULL,
        avalaible_on TEXT,
        created_by TEXT NOT NULL,
        prix BIGINT,
        -- status will be an boolean is open or not 
        status BOOLEAN,
        services TEXT[],
        images TEXT[],
        created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
        updated_at TIMESTAMP(0) WITHOUT TIME ZONE,
        deleted_at TIMESTAMP(0) WITHOUT TIME ZONE,
        PRIMARY KEY (uuid)
    );