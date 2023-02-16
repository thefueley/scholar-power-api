CREATE TABLE IF NOT EXISTS workout (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    plan_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    sets TEXT NOT NULL,
    reps TEXT NOT NULL,
    load TEXT NOT NULL,
    created_at TEXT DEFAULT (strftime('%m-%d-%Y %H:%M', CURRENT_TIMESTAMP)),
    edited_at TEXT DEFAULT (strftime('%m-%d-%Y %H:%M', CURRENT_TIMESTAMP)),
    creator_id INTEGER NOT NULL,
    exercise_id INTEGER NOT NULL,
    instructions_id INTEGER NOT NULL,
    FOREIGN KEY (creator_id) REFERENCES user(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (exercise_id) REFERENCES exercise(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (instructions_id) REFERENCES exercise(instructions) ON DELETE CASCADE ON UPDATE CASCADE
);