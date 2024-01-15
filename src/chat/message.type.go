package chat

type MessageType string

const (
	JoinRoomType    MessageType = "join_room"
	LeaveRoomType   MessageType = "leave_room"
	ChatMessageType MessageType = "chat_message"
)
