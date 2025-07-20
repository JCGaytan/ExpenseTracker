CREATE TABLE IF NOT EXISTS transactions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    amount REAL NOT NULL,
    date DATETIME NOT NULL,
    category_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    description TEXT,
    FOREIGN KEY(category_id) REFERENCES categories(id),
    FOREIGN KEY(user_id) REFERENCES users(id)
);
