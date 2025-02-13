CREATE TABLE IF NOT EXISTS merch_store.products (
    id TEXT NOT NULL PRIMARY KEY,
    title TEXT NOT NULL,
    price INTEGER NOT NULL CHECK (price > 0),
    created_at TIMESTAMP DEFAULT TIMEZONE('UTC'::TEXT, NOW()) NOT NULL,
    updated_at TIMESTAMP DEFAULT TIMEZONE('UTC'::TEXT, NOW()) NOT NULL
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW() AT TIME ZONE 'UTC';
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_products_updated_at
    BEFORE UPDATE ON merch_store.products
    FOR EACH ROW
    EXECUTE PROCEDURE update_timestamp();
