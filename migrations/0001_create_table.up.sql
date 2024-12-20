CREATE TABLE IF NOT EXISTS refresh_token (
    id SERIAL PRIMARY KEY,
    refresh_token VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(400) NOT NULL,
    refresh_token_id INT,
    CONSTRAINT fk_refresh_token
        FOREIGN KEY (refresh_token_id)
        REFERENCES refresh_token(id)
        ON DELETE CASCADE
);


CREATE UNIQUE INDEX idx_users_email ON users(email);

