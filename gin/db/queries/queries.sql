-- name: ReadRows :many
SELECT *
FROM products
LIMIT 100;

-- name: WriteRows :exec
INSERT INTO products (name, value)
VALUES ($1, $2);