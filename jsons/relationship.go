package jsons

type Relationship struct {
	Meta Meta             `json:"meta"`
	Data RelationshipData `json:"data"`
}

type RelationshipData struct {
	OutgoingStatus string `json:"outgoing_status"`
	IncomingStatus string `json:"incoming_status"`
}
