CREATE TABLE snippets (
    snippet_id serial PRIMARY KEY,
    title varchar(100) NOT NULL,
    snippet_content text NOT NULL,
    created timestamp NOT NULL,
    expires timestamp NOT NULL
);

CREATE TABLE users (
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password char(60) NOT NULL,
    created timestamp NOT NULL
);

ALTER TABLE
    users
ADD
    CONSTRAINT users_uc_email UNIQUE (email);

INSERT INTO
    users (name, email, password, created)
VALUES
    (
        'Alice Jones',
        'alice@example.com',
        '$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG',
        '2018-12-23 17:25:22'
    );
