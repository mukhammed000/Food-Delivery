DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS cart_items;

DROP TABLE IF EXISTS carts;

DROP TABLE IF EXISTS orders;

DROP TABLE IF EXISTS offices;

DROP TABLE IF EXISTS products;

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_enum') THEN
        DROP TYPE role_enum;
    END IF;
END$$;

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender_enum') THEN
        DROP TYPE gender_enum;
    END IF;
END$$;
