DROP PROCEDURE IF EXISTS add_posts;

CREATE PROCEDURE add_posts(posts_json JSONB)
AS $$
BEGIN
    INSERT INTO posts
    (
        title,
        content,
        post_time,
        link
    )
    SELECT (p->>'Title')::TEXT,
           (p->>'Content')::TEXT,
           (p->>'PostTime')::BIGINT,
           (p->>'Link')::TEXT
    FROM jsonb_array_elements(posts_json) p;
END;
$$ LANGUAGE plpgsql;