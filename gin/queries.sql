-- name: Read :many
SELECT * FROM product
LIMIT $1;

-- name: Write :exec
INSERT INTO
    product (name, value)
VALUES
    ($1, $2);
