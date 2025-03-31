DROP TYPE IF EXISTS auth_type CASCADE;
CREATE TYPE auth_type AS ENUM ('google', 'email');

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(244) UNIQUE NOT NULL,
    password TEXT,
    email VARCHAR(255) UNIQUE NOT NULL,
    img TEXT,
    since TIMESTAMP DEFAULT now(),
    description TEXT,
    provider auth_type NOT NULL,
    verified BOOLEAN NOT NULL
);

CREATE UNIQUE INDEX idx_users_username ON users(username);

CREATE UNIQUE INDEX idx_users_name ON users(name);

CREATE TABLE servers (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    since TIMESTAMP DEFAULT now(),
    img TEXT,
    banner TEXT
);

DROP TYPE IF EXISTS message_type CASCADE;
CREATE TYPE message_type AS ENUM ('image', 'video', 'document', 'text', 'link');

CREATE TABLE user_to_user_chat_mapping (
    id BIGSERIAL PRIMARY KEY,
    content TEXT,
    from_user_id BIGINT NOT NULL,
    to_user_id BIGINT  NOT NULL,
    type message_type NOT NULL,
    time_at BIGINT NOT NULL,
    reply_of BIGINT,
    FOREIGN KEY (reply_of) REFERENCES user_to_user_chat_mapping(id)
        ON DELETE SET NULL DEFERRABLE INITIALLY DEFERRED,
    FOREIGN KEY (from_user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (to_user_id) REFERENCES  users(id) ON DELETE CASCADE
);


DROP TYPE IF EXISTS user_role CASCADE;
CREATE TYPE user_role AS ENUM ('admin', 'moderator', 'member');

CREATE TABLE server_to_user_mapping (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    server_id BIGINT NOT NULL,
    role user_role NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (server_id) REFERENCES servers(id) ON DELETE CASCADE,
    UNIQUE (user_id, server_id)
);
CREATE INDEX idx_user_server_id ON server_to_user_mapping(server_id, user_id); -- for search of roles and all, known the server
CREATE INDEX idx_server_id_of_mapping ON server_to_user_mapping(server_id); -- For simple seach of servers for a user

DROP TYPE IF EXISTS channel_type CASCADE;
CREATE TYPE channel_type AS ENUM ('text', 'voice', 'stage', 'announcement');

CREATE TABLE channels (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    type channel_type NOT NULL,
    server_id BIGINT NOT NULL,
    allowed_roles user_role NOT NULL, -- it is conditional based on sets permissions, eg., if member is allowed, that means everyone, if mods allowed, then below them no one is allowed
    description TEXT,
    FOREIGN KEY (server_id) REFERENCES servers(id) ON DELETE CASCADE
);

CREATE INDEX idx_channels_server_id ON channels(server_id);

CREATE TABLE user_to_channel_chat_mapping (
    id BIGSERIAL PRIMARY KEY,
    content TEXT,
    reply_of BIGINT,
    from_user_id BIGINT NOT NULL,
    channel_id BIGINT NOT NULL,
    type message_type NOT NULL,
    time_at BIGINT NOT NULL,
    FOREIGN KEY (from_user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (reply_of) REFERENCES user_to_channel_chat_mapping(id)
        ON DELETE SET NULL DEFERRABLE INITIALLY DEFERRED,
    FOREIGN KEY (channel_id) REFERENCES channels(id) ON DELETE SET NULL
);
-- CREATE INDEX idx_reply_of ON user_to_channel_chat_mapping(reply_of);
-- CREATE INDEX idx_channel_id ON user_to_channel_chat_mapping(channel_id, from_user_id, time_at);

CREATE TABLE user_to_channel_session_mapping (
      id BIGSERIAL PRIMARY KEY,
      user_id BIGINT NOT NULL,
      joined_at TIMESTAMP,
      left_at TIMESTAMP,
      FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE invites (
    id VARCHAR(10) PRIMARY KEY,
    server_id BIGINT NOT NULL REFERENCES servers(id) ON DELETE CASCADE,
    expiry BIGINT, -- assume -1 means no expiry
    uses INTEGER, -- assume -1 means unlimited uses
    valid BOOLEAN NOT NULL,
    FOREIGN KEY (server_id) REFERENCES servers(id) ON DELETE CASCADE
);

ALTER TABLE invites
    ADD CONSTRAINT check_expiry_or_uses
        CHECK (expiry IS NOT NULL OR uses IS NOT NULL);
