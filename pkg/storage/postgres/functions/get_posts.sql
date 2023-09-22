DROP FUNCTION IF EXISTS get_posts;

CREATE FUNCTION get_posts()

RETURNS TABLE
(
    id BIGINT,
    author_id BIGINT,
    author TEXT,
    title TEXT,
    content TEXT,
    created_at BIGINT,
    updated_at BIGINT
) AS $$

BEGIN
    RETURN QUERY
    SELECT p.id,
           a.id AS author_id,
           a.name AS author,
           p.title,
           p.content,
           p.created_at,
           p.updated_at
    FROM posts p
    INNER JOIN authors a
        ON a.id = p.author_id;
END;
$$ LANGUAGE plpgsql;