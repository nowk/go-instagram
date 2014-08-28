package jsons

type Meta struct {
	ErrorType    string `json:"error_type"`
	Code         int    `json:"code"`
	ErrorMessage string `json:"error_message"`
}
