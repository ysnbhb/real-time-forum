CREATE TABLE
    IF NOT EXISTS user (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        age INTEGER NOT NULL CHECK (age >= 18),
        firsname TEXT NOT NULL,
        lastname TEXT NOT NULL,
        gender TEXT CHECK (gender IN ('male', 'female')) NOT NULL,
        username TEXT NOT NULL UNIQUE,
        email TEXT NOT NULL UNIQUE,
        passwd TEXT NOT NULL,
        uidlog TEXT NOT NULL UNIQUE,
        create_date DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    IF NOT EXISTS post (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER NOT NULL,
        title TEXT NOT NULL,
        contant TEXT NOT NULL,
        create_date DEFAULT CURRENT_TIMESTAMP,
        img TEXT,
        categories TEXT NOU NULL,
        FOREIGN KEY (user_id) REFERENCES user (id)
    );

CREATE TABLE
    IF NOT EXISTS categories (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name_categorie TEXT UNIQUE NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS categories_post (
        categorie_id INTEGER NOT NULL,
        post_id INTEGER NOT NULL,
        FOREIGN KEY (categorie_id) REFERENCES categories (id),
        FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE UNIQUE (categorie_id, post_id)
    );

CREATE TABLE
    IF NOT EXISTS comment (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        post_id INTEGER NOT NULL,
        user_id INTEGER NOT NULL,
        contant TEXT NOT NULL,
        create_date DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE,
        FOREIGN KEY (user_id) REFERENCES user (id)
    );

CREATE TABLE
    IF NOT EXISTS messages (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        fromid INTEGER NOT NULL,
        toid INTEGER NOT NULL,
        message TEXT NOT NULL,
        create_date DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (fromid) REFERENCES user (id),
        FOREIGN KEY (toid) REFERENCES user (id)
    );

CREATE TABLE
    IF NOT EXISTS reaction (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        post_id INTEGER,
        comment_id INTEGER,
        user_id INTEGER NOT NULL,
        type TEXT CHECK (type IN ('likes', 'dislikes')) NOT NULL,
        FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE,
        FOREIGN KEY (user_id) REFERENCES user (id),
        FOREIGN KEY (comment_id) REFERENCES comment (id),
        CHECK (
            (
                post_id IS NOT NULL
                AND comment_id IS NULL
            )
            OR (
                post_id IS NULL
                AND comment_id IS NOT NULL
            )
        ),
        UNIQUE (user_id, post_id),
        UNIQUE (user_id, comment_id)
    );

CREATE TABLE
    IF NOT EXISTS session (
        id INTEGER PRIMARY KEY,
        uid TEXT UNIQUE NOT NULL,
        user_id INTEGER UNIQUE NOT NULL,
        create_date DATE DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES user (id)
    );

INSERT INTO
    categories (name_categorie)
VALUES
    ('Game'),
    ('Tecnolge'),
    ('Ecomerc'),
    ('Natur'),
    ('Viset') ON CONFLICT (name_categorie) DO NOTHING;