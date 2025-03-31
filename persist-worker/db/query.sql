-- name: BatchInsertUserToChannelChat :exec
INSERT INTO user_to_channel_chat_mapping (content, reply_of, from_user_id, channel_id, type, time_at)
SELECT
    u_content,
    u_reply_of,
    u_from_user_id,
    u_channel_id,
    u_type::message_type,
        u_time_at
FROM unnest(@content::TEXT[]) WITH ORDINALITY AS t1(u_content, ord),
     unnest(ARRAY(SELECT CASE WHEN x = 0 THEN NULL ELSE x END FROM unnest(@reply_of::BIGINT[]) AS x)) WITH ORDINALITY AS t2(u_reply_of, ord),
     unnest(@from_user_id::BIGINT[]) WITH ORDINALITY AS t3(u_from_user_id, ord),
     unnest(@channel_id::BIGINT[]) WITH ORDINALITY AS t4(u_channel_id, ord),
     unnest(@type::TEXT[]) WITH ORDINALITY AS t5(u_type, ord),
     unnest(@time_at::BIGINT[]) WITH ORDINALITY AS t6(u_time_at, ord)
WHERE t1.ord = t2.ord
  AND t2.ord = t3.ord
  AND t3.ord = t4.ord
  AND t4.ord = t5.ord
  AND t5.ord = t6.ord;

-- name: BatchInsertUserToUserChat :exec
INSERT INTO user_to_user_chat_mapping (content, from_user_id, to_user_id, reply_of, type, time_at)
SELECT
    u_content,
    u_from_user_id,
    u_to_user_id,
    u_reply_of,
    u_type::message_type,
        u_time_at
FROM unnest(@content::TEXT[]) WITH ORDINALITY AS t1(u_content, ord),
     unnest(@from_user_id::BIGINT[]) WITH ORDINALITY AS t2(u_from_user_id, ord),
     unnest(@to_user_id::BIGINT[]) WITH ORDINALITY AS t3(u_to_user_id, ord),
     unnest(ARRAY(SELECT CASE WHEN x = 0 THEN NULL ELSE x END FROM unnest(@reply_of::BIGINT[]) AS x)) WITH ORDINALITY AS t4(u_reply_of, ord),
     unnest(@type::TEXT[]) WITH ORDINALITY AS t5(u_type, ord),
     unnest(@time_at::BIGINT[]) WITH ORDINALITY AS t6(u_time_at, ord)
WHERE t1.ord = t2.ord
  AND t2.ord = t3.ord
  AND t3.ord = t4.ord
  AND t4.ord = t5.ord
  AND t5.ord = t6.ord;
