DROP FUNCTION IF EXISTS delete_post;

CREATE FUNCTION delete_post
(
    p_id BIGINT
)
RETURNS BOOLEAN AS $$

BEGIN
    DELETE
    FROM posts p
    WHERE p.id = p_id;

	RETURN FOUND;
END;
$$ LANGUAGE plpgsql;