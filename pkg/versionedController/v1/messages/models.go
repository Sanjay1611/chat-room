package messages

type postRequestModel struct {
	Content string `json:"content"`
	UserID  string `json:"user_id"`
}
