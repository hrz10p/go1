PRAGMA foreign_keys = ON;


CREATE TABLE IF NOT EXISTS users (
                       id VARCHAR PRIMARY KEY,
                       username VARCHAR UNIQUE,
                       email VARCHAR UNIQUE,
                       password VARCHAR(60)
);

CREATE TABLE posts (
                       id INTEGER PRIMARY KEY AUTOINCREMENT,
                       title VARCHAR,
                       content VARCHAR,
                       UID VARCHAR,
                       FOREIGN KEY (UID) REFERENCES users(id) ON DELETE CASCADE
);


CREATE TABLE categories (
                            id INTEGER PRIMARY KEY AUTOINCREMENT,
                            name VARCHAR
);

CREATE TABLE post_cats (
                           post_id INTEGER,
                           category_id INTEGER,
                           PRIMARY KEY (post_id, category_id),
                           FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
                           FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);

CREATE TABLE sessions (
                          id VARCHAR,
                          uid VARCHAR,
                          expireTime DATE,
                          PRIMARY KEY (id, uid),
                          FOREIGN KEY (uid) REFERENCES users(id) ON DELETE CASCADE
);


INSERT INTO categories (name) VALUES
                                  ('Applicants'),
                                  ('Staff'),
                                  ('Researches');

