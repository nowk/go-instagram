package jsons

type Location struct {
	Meta Meta         `json:"meta"`
	Data LocationData `json:"data"`
}

type Locations struct {
	Meta       Meta           `json:"meta"`
	Pagination Pagination     `json:"pagination"`
	Data       []LocationData `json:"data"`
}

type LocationData struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
