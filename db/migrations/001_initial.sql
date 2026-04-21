--user table
CREATE TABLE IF NOT EXISTS users (
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(50) UNIQUE NOT NULL,
    email      VARCHAR(255) UNIQUE NOT NULL,
    password   TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

--games table
CREATE TABLE IF NOT EXISTS games (
    id           SERIAL PRIMARY KEY,
    steam_app_id INT UNIQUE NOT NULL,
    name         VARCHAR(255) NOT NULL,
    created_at   TIMESTAMP DEFAULT NOW()
);

--user_games table
CREATE TABLE IF NOT EXISTS user_games (
    id         SERIAL PRIMARY KEY,
    user_id    INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    game_id    INT NOT NULL REFERENCES games(id) ON DELETE CASCADE,
    status     VARCHAR(20) NOT NULL DEFAULT 'unplayed',
    rating     INT CHECK (rating >= 1 AND rating <= 10),
    review     TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (user_id, game_id)
);