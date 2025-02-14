-- name: CreateUser :one
INSERT INTO users (
    email, hashed_password, name, role
) VALUES (
             $1, $2, $3, $4
         )
    RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- Get user by email (for login)
-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;