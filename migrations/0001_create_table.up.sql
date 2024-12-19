CREATE TABLE IF NOT EXISTS refresh_token (
    id SERIAL PRIMARY KEY,
    refresh_token VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    refresh_token_id INT NOT NULL,
    CONSTRAINT fk_refresh_token
        FOREIGN KEY (refresh_token_id)
        REFERENCES refresh_token(id)
        ON DELETE CASCADE
);

