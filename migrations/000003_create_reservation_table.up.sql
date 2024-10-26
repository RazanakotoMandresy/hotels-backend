CREATE TABLE
    reservation (
        uuid UUID NOT NULL,
        reserved_by_uuid UUID NOT NULL,
        hotels_uuid UUID NOT NULL,
        reservation_date_start TEXT NOT NULL,
        reservation_date_end TEXT NOT NULL,
        PRIMARY KEY (uuid)
    );
