package users

type postRequestModel struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

type putRequestModel struct {
	ChatRoomID string `json:"chat_room_id"`
	UserID     string `json:"user_id"`
}
