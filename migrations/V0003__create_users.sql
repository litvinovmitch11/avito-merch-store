CREATE TABLE IF NOT EXISTS merch_store.users (
    id TEXT NOT NULL PRIMARY KEY,
    username TEXT NOT NULL,
    balance INTEGER NOT NULL CHECK (balance > 0),
    created_at TIMESTAMP DEFAULT TIMEZONE('UTC'::TEXT, NOW()) NOT NULL,
    updated_at TIMESTAMP DEFAULT TIMEZONE('UTC'::TEXT, NOW()) NOT NULL
);

CREATE TABLE IF NOT EXISTS merch_store.personal_data (
    id TEXT NOT NULL PRIMARY KEY,
    user_id TEXT NOT NULL,
    hashed_password TEXT NOT NULL,
    salt TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT TIMEZONE('UTC'::TEXT, NOW()) NOT NULL,
    updated_at TIMESTAMP DEFAULT TIMEZONE('UTC'::TEXT, NOW()) NOT NULL,
    
    FOREIGN KEY (user_id) REFERENCES merch_store.users(id) 
);

CREATE TRIGGER trigger_users_updated_at
    BEFORE UPDATE ON merch_store.users
    FOR EACH ROW
    EXECUTE PROCEDURE update_timestamp();

CREATE TRIGGER trigger_personal_data_updated_at
    BEFORE UPDATE ON merch_store.personal_data
    FOR EACH ROW
    EXECUTE PROCEDURE update_timestamp();
