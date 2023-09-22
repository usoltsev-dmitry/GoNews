DROP FUNCTION IF EXISTS add_post;

CREATE FUNCTION add_post
(
    p_author_id BIGINT,
    p_title TEXT,
    p_content TEXT
)
RETURNS BIGINT AS $$

DECLARE v_post_id BIGINT;

BEGIN
    INSERT INTO posts
    (
        author_id,
        title,
        content
    )
    SELECT p_author_id,
           p_title,
           p_content
    RETURNING id INTO v_post_id;
	RETURN v_post_id;
END;
$$ LANGUAGE plpgsql;