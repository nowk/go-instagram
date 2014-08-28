package jsons

type Media struct {
	Meta Meta      `json:"meta"`
	Data MediaData `json:"data"`
}

type Medias struct {
	Meta       Meta        `json:"meta"`
	Pagination Pagination  `json:"Pagination"`
	Data       []MediaData `json:"data"`
}

type MediaData struct {
	Attribution interface{} `json:"attribution"`
	Tags        []string    `json:"tags"`
	Type        string      `json:"type"`

	Location struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`

	Comments struct {
		Count int           `json:"count"`
		Data  []CommentData `json:"data"`
	} `json:"comments"`

	Filter      string `json:"filter"`
	CreatedTime string `json:"created_time"`
	Link        string `json:"link"`

	Likes struct {
		Count int        `json:"count"`
		Data  []UserData `json:"data"`
	} `json:"likes"`

	Images struct {
		LowResolution      Image `json:"low_resolution"`
		Thumbnail          Image `json:"thumbnail"`
		StandardResolution Image `json:"standard_resolution"`
	} `json:"images"`

	Videos struct {
		LowResolution      Image `json:"low_resolution"`
		StandardResolution Image `json:"standard_resolution"`
	} `json:"videos"`

	UsersInPhoto []struct {
		Position struct {
			Y float64 `json:"y"`
			X float64 `json:"x"`
		} `json:"position"`
		User UserData `json:"user"`
	} `json:"users_in_photo"`

	Caption struct {
		CreatedTime string   `json:"created_time"`
		Text        string   `json:"text"`
		From        UserData `json:"from"`
		ID          string   `json:"id"`
	} `json:"caption"`

	UserHasLiked bool   `json:"user_has_liked"`
	ID           string `json:"id"`

	User UserData `json:"user"`
}
