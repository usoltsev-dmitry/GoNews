DROP FUNCTION IF EXISTS add_author;

CREATE FUNCTION add_author
(
    p_name TEXT
)
RETURNS BIGINT AS $$

DECLARE v_author_id BIGINT;

BEGIN
    INSERT INTO authors
    (
        name
    )
    SELECT p_name
    RETURNING id INTO v_author_id;
	RETURN v_author_id;
END;
$$ LANGUAGE plpgsql;