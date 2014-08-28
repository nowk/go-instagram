package jsons

type Subscription struct {
	Meta Meta             `json:"meta"`
	Data SubscriptionData `json:"data"`
}

type Subscriptions struct {
	Meta Meta               `json:"meta"`
	Data []SubscriptionData `json:"data"`
}

type SubscriptionData struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Object      string `json:"object"`
	Aspect      string `json:"aspect"`
	CallbackURL string `json:"callback_url"`
}
