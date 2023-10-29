DROP TABLE IF EXISTS posts;

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    pubdate TIMESTAMPTZ NOT NULL,
    link TEXT NOT NULL
);

CREATE INDEX idx_posts_post_time ON posts (pubdate);