DROP FUNCTION IF EXISTS get_posts;

CREATE FUNCTION get_posts(p_limit INT)

RETURNS TABLE
(
    id INT,
    title TEXT,
    content TEXT,
    post_time BIGINT,
    link TEXT
) AS $$

BEGIN
    RETURN QUERY
    SELECT p.id,
           p.title,
           p.content,
           p.post_time,
           p.link
    FROM posts p
    ORDER BY p.post_time DESC
    LIMIT p_limit;
END;
$$ LANGUAGE plpgsql;