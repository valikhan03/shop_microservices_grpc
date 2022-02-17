CREATE TABLE IF NOT EXISTS category(
    slug TEXT UNIQUE,
    c_path TEXT,
    parent_id TEXT,
    c_name TEXT,
    c_status INTEGER
);