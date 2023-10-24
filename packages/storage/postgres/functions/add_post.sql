DROP FUNCTION IF EXISTS add_post;

CREATE FUNCTION add_post
(
    p_title TEXT,
    p_content TEXT,
    p_post_time BIGINT,
    p_link TEXT
)
RETURNS BIGINT AS $$

DECLARE v_id BIGINT;

BEGIN
    INSERT INTO posts
    (
        title,
        content,
        post_time,
        link
    )
    SELECT p_title,
           p_content,
           p_post_time,
           p_link
    RETURNING id INTO v_id;

    RETURN v_id;
END;
$$ LANGUAGE plpgsql;