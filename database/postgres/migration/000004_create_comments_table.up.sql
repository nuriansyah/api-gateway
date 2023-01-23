CREATE TABLE IF NOT EXISTS comments (
    id serial PRIMARY KEY,
    post_id integer NOT NULL,
    author_id integer NOT NULL,
    content text NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (post_id) REFERENCES posts(id),
    FOREIGN KEY (author_id) REFERENCES users(id)
);