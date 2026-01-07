CREATE TABLE translations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    english TEXT NOT NULL CHECK (english != ""),
    russian TEXT NOT NULL CHECK (russian != "")
);
