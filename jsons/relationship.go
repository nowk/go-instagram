package jsons

type Relationship struct {
	Meta Meta `json:"meta"`
	Data struct {
		OutgoingStatus string `json:"outgoing_status"`
		IncomingStatus string `json:"incoming_status"`
	} `json:"data"`
}
