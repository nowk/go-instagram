package jsons

type Like struct {
	Meta Meta     `json:"meta"`
	Data LikeData `json:"data"`
}

type Likes struct {
	Meta       Meta       `json:"meta"`
	Pagination Pagination `json:"pagination"`
	Data       []LikeData `json:"data"`
}

type LikeData struct {
	UserName  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Type      string `json:"type"`
	ID        string `json:"id"`
}
