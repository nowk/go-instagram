package jsons

type Comment struct {
	Meta Meta        `json:"meta"`
	Data CommentData `json:"data"`
}

type Comments struct {
	Meta       Meta          `json:"meta"`
	Pagination Pagination    `json:"pagination"`
	Data       []CommentData `json:"data"`
}

type CommentData struct {
	CreatedTime string   `json:"created_time"`
	Text        string   `json:"text"`
	From        UserData `json:"from"`
	ID          string   `json:"id"`
}
