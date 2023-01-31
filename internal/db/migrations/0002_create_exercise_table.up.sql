CREATE TABLE IF NOT EXISTS exercise (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    muscle TEXT NOT NULL,
    equipment TEXT NOT NULL,
    instructions TEXT NOT NULL
);