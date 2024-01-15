package processor

type Message struct {
	Sender     string `json:"sender"`
	SenderName string `json:"senderName"`
	Room       string `json:"room"`
	Type       string `json:"type"`
	Content    string `json:"content"`
	Server     string `json:"server"`
}
