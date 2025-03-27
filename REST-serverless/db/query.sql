-- -- name: ListServers :many
-- SELECT * FROM servers
-- WHERE;

-- name: CreateUser :one
INSERT INTO users (
    name, username, email, img, description
) VALUES (
             $1, $2, $3, $4, $5
         )
RETURNING *;
