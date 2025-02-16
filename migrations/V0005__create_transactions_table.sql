CREATE TABLE IF NOT EXISTS merch_store.transactions (
    id TEXT NOT NULL PRIMARY KEY,
    from_id TEXT NOT NULL,
    to_id TEXT,
    amount INTEGER CHECK (amount > 0) NOT NULL,
    created_at TIMESTAMP DEFAULT TIMEZONE('UTC'::TEXT, NOW()) NOT NULL,
    updated_at TIMESTAMP DEFAULT TIMEZONE('UTC'::TEXT, NOW()) NOT NULL,

    FOREIGN KEY (from_id) REFERENCES merch_store.users(id),
    FOREIGN KEY (to_id) REFERENCES merch_store.users(id)
);

CREATE TRIGGER trigger_transactions_updated_at
    BEFORE UPDATE ON merch_store.transactions
    FOR EACH ROW
    EXECUTE PROCEDURE update_timestamp();
