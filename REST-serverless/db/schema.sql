

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(244) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    img TEXT,
    since TIMESTAMP DEFAULT now(),
    desc TEXT
);

CREATE UNIQUE INDEX idx_users_username ON users(username);

CREATE UNIQUE INDEX idx_users_name ON users(name);

CREATE TABLE servers (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    desc TEXT,
    since TIMESTAMP DEFAULT now(),
    img TEXT,
    banner TEXT
);

CREATE TYPE message_type AS ENUM ('image', 'video', 'document', 'text', 'link');

CREATE TABLE user_to_user_chat_mapping (
    id BIGSERIAL PRIMARY KEY,
    content TEXT,
    from_user_id BIGINT,
    to_user_id BIGINT,
    type message_type,
    FOREIGN KEY (from_user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (to_user_id) REFERENCES  user(id) ON DELETE CASCADE
);


CREATE TYPE user_role AS ENUM ('admin', 'moderator', 'member');

CREATE TABLE server_to_user_mapping (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT,
    server_id BIGINT,
    role user_role,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (server_id) REFERENCES servers(id) ON DELETE CASCADE
);

CREATE TYPE channel_type AS ENUM ('text', 'voice', 'stage', 'announcement');

CREATE TABLE channels (
    id BIGSERIAL PRIMARY KEY,
    name TEXT,
    type channel_type,
    allowed_roles user_role              e
);

CREATE TABLE channel_to_server_mapping (
    id BIGSERIAL PRIMARY KEY,
    channel_id BIGINT,
    server_id BIGINT,
    FOREIGN KEY (channel_id) REFERENCES channels(id) ON DELETE CASCADE,
    FOREIGN KEY (server_id) REFERENCES servers(id) ON DELETE CASCADE
);

CREATE TABLE user_to_channel_chat_mapping (
    id BIGSERIAL PRIMARY KEY,
    content TEXT,
    from_user_id BIGINT,
    to_user_id BIGINT,
    type message_type,
    FOREIGN KEY (from_user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (to_user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE user_to_channel_session_mapping (
      id BIGSERIAL PRIMARY KEY,
      user_id BIGSERIAL,
      joined_at TIMESTAMP,
      left_at TIMESTAMP,
      FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);