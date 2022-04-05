CREATE TABLE IF NOT EXISTS products(
      slug TEXT NOT NULL UNIQUE,
      title TEXT NOT NULL,
      price INT NOT NULL,
      description TEXT,
      category TEXT
);

