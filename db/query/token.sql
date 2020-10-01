-- name: CreateToken :one
INSERT INTO tokens (
    user_id,
    blacklisted,
    token,
    type,
    expires
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetToken :one 
SELECT * FROM tokens
WHERE token = $1 LIMIT 1;

-- name: ListToken :many
SELECT * FROM tokens
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateToken :one
UPDATE tokens
SET blacklisted = $2, expires = $3
WHERE id = $1
RETURNING *;

-- name: DeleteToken :exec
DELETE FROM tokens
WHERE id = $1;