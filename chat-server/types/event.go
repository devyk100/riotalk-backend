package types

type Event struct {
	Event string `json:"event"`
	Type  string `json:"type"`
	To    int64  `json:"to"`
	Token string `json:"token"`
}
