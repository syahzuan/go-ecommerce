CREATE TABLE IF NOT EXISTS orders (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  userId INTEGER NOT NULL,
  total REAL NOT NULL,
  status TEXT CHECK( status IN ('pending', 'completed', 'cancelled') ) NOT NULL DEFAULT 'pending',
  address TEXT NOT NULL,
  createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (userId) REFERENCES users(id)
);
