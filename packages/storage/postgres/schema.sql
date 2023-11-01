DROP TABLE IF EXISTS posts;

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    pubdate TIMESTAMPTZ NOT NULL,
    link TEXT NOT NULL
);
--Добавим индекс на дату публикции для быстрой сортировки данных по дате
CREATE INDEX idx_posts_post_time ON posts (pubdate);
--Добавим уникальный индекс на ссылку для предотвращения записи дубликатов публикации
CREATE UNIQUE INDEX link_unique_idx ON posts (link);