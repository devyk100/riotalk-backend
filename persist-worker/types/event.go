package types

type Event struct {
	Event          string `json:"event"`            // chat, auth, history
	Type           string `json:"type"`             // server, user
	IsTokenExpired bool   `json:"is_token_expired"` // if the server wants to return something like this to render a refresh accept token
	To             int64  `json:"to"`               // server_id, user_id
	Of             int64  `json:"of"`               // In case for requesting the history
	ChannelId      int64  `json:"channel_id"`       // channel_id, if Type is server
	Token          string `json:"token"`            // in auth
	FromID         int64  `json:"from_id"`          //
	FromName       string `json:"from_name"`        //
	FromUsername   string `json:"from_username"`    //
	FromImg        string `json:"from_img"`         //
	Content        string `json:"content"`          //
	TimeAt         int64  `json:"time_at"`          //
	MessageType    string `json:"message_type"`     //
	ReplyOf        int64  `json:"reply_of"`         // references another chat id in same list
}
