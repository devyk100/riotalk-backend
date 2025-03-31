DROP TYPE IF EXISTS message_type CASCADE;
CREATE TYPE message_type AS ENUM ('image', 'video', 'document', 'text', 'link');

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