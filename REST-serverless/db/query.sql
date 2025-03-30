-- name: CreateUserOrDoNothing :one
INSERT INTO users (name, username, password, email, img, description, provider, verified)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    ON CONFLICT (email) DO NOTHING RETURNING *;

-- name: CreateUserOrThrow :one
INSERT INTO users (name, username, password, email, img, description, provider, verified)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: CreateServer :one
INSERT INTO servers (name, description, img, banner)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: CreateServerAndMapping :one
WITH inserted_server AS (
    INSERT INTO servers (name, description, img, banner)
    VALUES ($1, $2, $3, $4)
        RETURNING id
)
INSERT INTO server_to_user_mapping (user_id, server_id, role)
SELECT $5, id, 'admin' FROM inserted_server
RETURNING server_id;

-- name: GetPasswordFromUserNameEmail :one
SELECT password, id
FROM users
WHERE email = $1 OR username = $1
    LIMIT 1;


-- name: CreateChannel :one
INSERT INTO channels (name, type, server_id, allowed_roles, description)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: CreateChannelIfAuthorized :one
WITH user_role_check AS (
    SELECT role FROM server_to_user_mapping
    WHERE user_id = $1 AND server_id = $3
)
INSERT INTO channels (name, type, server_id, allowed_roles, description)
SELECT $2, $4, $3, $5, $6
FROM user_role_check
WHERE role IN ('admin', 'moderator')
    RETURNING *;


-- name: GetChannelList :many
WITH user_role_cte AS (
    SELECT role FROM server_to_user_mapping
    WHERE user_id = $1 AND server_id = $2
)
SELECT c.*
FROM channels c
         JOIN user_role_cte u
              ON c.server_id = $2
WHERE c.allowed_roles IN (
    CASE
        WHEN u.role = 'admin' THEN ARRAY['admin', 'moderator', 'member']::user_role[]
        WHEN u.role = 'moderator' THEN ARRAY['moderator', 'member']::user_role[]
        WHEN u.role = 'member' THEN ARRAY['member']::user_role[]
        END
);

-- name: CreateServerInvite :one
INSERT INTO invites(id, server_id, expiry, uses, valid)
SELECT $1, $2, $3, $4, $6
FROM server_to_user_mapping
WHERE server_id = $2 AND user_id = $5 AND role IN ('admin', 'moderator')
    RETURNING *;


-- name: GetServerInvite :one
SELECT * FROM invites WHERE id = $1;


-- name: DecrementInviteUses :exec
UPDATE invites
SET uses = uses - 1
WHERE id = $1 AND uses > 0;

-- name: CreateServerToUserMapping :one
INSERT INTO server_to_user_mapping(user_id, server_id, role)
VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET
    name = COALESCE(NULLIF($2, ''), name),
    username = COALESCE(NULLIF($3, ''), username),
    img = COALESCE(NULLIF($4, ''), img),
    description = COALESCE(NULLIF($5, ''), description)
WHERE id = $1;

-- name: UpdateChannel :exec
UPDATE channels
SET
    name = COALESCE(NULLIF($3, ''), name),
    allowed_roles = COALESCE(NULLIF($4, ''), allowed_roles),
    description = COALESCE(NULLIF($5, ''), description)
WHERE channels.id = $1
  AND EXISTS (
    SELECT 1
    FROM server_to_user_mapping
    WHERE server_to_user_mapping.server_id = channels.server_id
      AND server_to_user_mapping.user_id = $2
      AND server_to_user_mapping.role IN ('admin', 'moderator')
);

-- name: GetServersList :many
SELECT * FROM servers s
JOIN server_to_user_mapping m ON s.id = m.server_id
WHERE m.user_id = $1;

-- name: UpdateUserRole :exec
UPDATE server_to_user_mapping AS target
SET role = $3
    FROM server_to_user_mapping AS initiator
WHERE target.user_id = $2
  AND target.server_id = initiator.server_id
  AND initiator.user_id = $1
  AND (
    (initiator.role = 'admin' AND $3 IN ('admin', 'moderator', 'member'))
   OR
    (initiator.role = 'moderator' AND $3 IN ('moderator', 'member'))
    );
