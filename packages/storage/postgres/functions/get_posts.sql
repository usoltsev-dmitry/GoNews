DROP FUNCTION IF EXISTS get_posts;

CREATE FUNCTION get_posts(p_limit INT)

RETURNS TABLE
(
    id INT,
    title TEXT,
    content TEXT,
    pubdate TIMESTAMPTZ,
    link TEXT
) AS $$

BEGIN
    RETURN QUERY
    SELECT p.id,
           p.title,
           p.content,
           p.pubdate,
           p.link
    FROM posts p
    ORDER BY p.pubdate DESC
    LIMIT p_limit;
END;
$$ LANGUAGE plpgsql;