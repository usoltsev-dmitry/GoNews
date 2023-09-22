DROP FUNCTION IF EXISTS update_post;

CREATE FUNCTION update_post
(
    p_id BIGINT,
    p_title TEXT,
    p_content TEXT
)
RETURNS BOOLEAN AS $$
BEGIN
    UPDATE posts p
    SET title = p_title,
        content = p_content,
        updated_at = extract(epoch from now())
    WHERE p.id = p_id
    AND (
            p.title != p_title
            OR p_content != p_content
        );

	RETURN FOUND;
END;
$$ LANGUAGE plpgsql;