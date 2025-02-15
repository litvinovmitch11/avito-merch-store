CREATE TABLE IF NOT EXISTS merch_store.storage (
    id TEXT NOT NULL PRIMARY KEY,
    user_id TEXT NOT NULL,
    balance INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT TIMEZONE('UTC'::TEXT, NOW()) NOT NULL,
    updated_at TIMESTAMP DEFAULT TIMEZONE('UTC'::TEXT, NOW()) NOT NULL,

    FOREIGN KEY (user_id) REFERENCES merch_store.users(id)
);

CREATE TRIGGER trigger_storage_updated_at
    BEFORE UPDATE ON merch_store.storage
    FOR EACH ROW
    EXECUTE PROCEDURE update_timestamp();
