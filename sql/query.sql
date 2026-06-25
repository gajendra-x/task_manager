-- name: GetAuthor :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO users (
  name, email
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateAuthor :exec
UPDATE users
  set name = $2,
  email = $3
WHERE id = $1;
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM users
WHERE id = $1;