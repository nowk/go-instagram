package jsons

type Tag struct {
	Meta Meta    `json:"meta"`
	Data TagData `json:"data"`
}

type Tags struct {
	Meta       Meta       `json:"meta"`
	Pagination Pagination `json:"pagination"`
	Data       []TagData  `json:"data"`
}

type TagData struct {
	MediaCount int    `json:"media_count"`
	Name       string `json:"name"`
}
