package jsons

type RealTimeUpdates struct {
	SubscriptionID int    `json:"subscription_id"`
	Object         string `json:"object"`
	ObjectID       string `json:"object_id"`
	ChangedAspect  string `json:"changed_aspect"`
	Time           int    `json:"time"`
}
