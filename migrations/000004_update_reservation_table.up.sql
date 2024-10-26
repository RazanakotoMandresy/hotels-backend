ALTER TABLE reservation
    ALTER COLUMN reserved_by_uuid TYPE TEXT USING reserved_by_uuid::TEXT,
    ALTER COLUMN hotels_uuid TYPE TEXT USING hotels_uuid::TEXT;
