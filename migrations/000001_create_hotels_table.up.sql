CREATE TABLE
    hotels (
        uuid UUID NOT NULL,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        services TEXT[],
        prix BIGINT,
        status SMALLINT,
        ouverture TEXT NOT NULL,
        -- pour l'instant mbola tsy atao not null ilay created by fa rehefa mahavita jwt sy mahita moyen update en temp reel anle izy 
        created_by TEXT,
        created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
        updated_at TIMESTAMP(0) WITHOUT TIME ZONE,
        deleted_at TIMESTAMP(0) WITHOUT TIME ZONE,
        PRIMARY KEY (uuid)
    );