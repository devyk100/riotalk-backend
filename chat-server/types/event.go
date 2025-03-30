package types

type Event struct {
	Event string `json:"event"` // chat, auth, active
	Type  string `json:"type"`  // channel, user
	To    int64  `json:"to"`    // channel_id, user_id
	Token string `json:"token"` // in auth
	From  int64  `json:"from"`  // from user_id, at the user end

}
