PRAGMA foreign_keys = ON;


CREATE TABLE IF NOT EXISTS users (
                       id VARCHAR PRIMARY KEY,
                       username VARCHAR UNIQUE,
                       email VARCHAR UNIQUE,
                       rolestring VARCHAR,
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

CREATE TABLE deps ( id INTEGER PRIMARY KEY AUTOINCREMENT, dep_name VARCHAR, staff_quantity INTEGER);

INSERT INTO users (id, username, email, rolestring, password) VALUES
('11fdc4ea-22ae-4656-b244-972630cc36f8', 'admin', 'admin@mail.ru', 'admin', '$2a$12$sqJoZPBnOiwWpMe2mczT6.B1JfCwerYA2fX7VF/RZrTTEDM6hyfqW');
