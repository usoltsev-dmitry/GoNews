DROP PROCEDURE IF EXISTS add_posts;

CREATE PROCEDURE add_posts(posts_json JSONB)
AS $$
BEGIN
    INSERT INTO posts
    (
        title,
        content,
        pubdate,
        link
    )
    SELECT (p->>'Title')::TEXT,
           (p->>'Content')::TEXT,
           (p->>'PubDate')::TIMESTAMPTZ,
           (p->>'Link')::TEXT
    FROM jsonb_array_elements(posts_json) p
    ON CONFLICT (link) DO NOTHING;
END;
$$ LANGUAGE plpgsql;