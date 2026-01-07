CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE CHECK (
        username != "" AND LENGTH(username) BETWEEN 5 AND 25
    ),
    password BLOB NOT NULL CHECK (LENGTH(password) = 64),
    created_at TEXT NOT NULL DEFAULT (DATETIME('now')),
    updated_at TEXT NOT NULL GENERATED ALWAYS AS (DATETIME('now')) STORED
);
