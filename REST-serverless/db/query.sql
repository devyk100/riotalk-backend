-- -- name: ListServers :many
-- SELECT * FROM servers
-- WHERE;

-- -- name: CreateUser :one
-- INSERT INTO users (
--     name, username, email, img, description
-- ) VALUES (
--              $1, $2, $3, $4, $5
--          )
-- RETURNING *;

-- name: CreateUserOrDoNothing :exec
INSERT INTO users (name, username, email, img, description, provider)
VALUES ($1, $2, $3, $4, $5, $6)
    ON CONFLICT (email) DO NOTHING;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;
