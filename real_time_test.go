package instagram_test

import "net/url"
import "testing"
import "github.com/nowk/assert"

func TestRealTimeSubscribe(t *testing.T) {
	api, mock := tNewRealTime(t, "my-clientid", "a-secret")

	reg := regmc(`v1\/subscriptions$`)
	_, mres := mock.Register("POST", reg, 200, `{"meta": {"code": 200}}`)

	data, _ := api.Subscribe("tag", "token", "http://example.com/rt/user", map[string]string{
		"object_id": "notifier",
	})

	assert.Equal(t, data.Meta.Code, 200)
	assert.TypeOf(t, "*jsons.Subscriptions", data)
	assert.TypeOf(t, "[]jsons.SubscriptionData", data.Data)

	v := url.Values{}
	v.Add("client_id", "my-clientid")
	v.Add("client_secret", "a-secret")
	v.Add("object", "tag")
	v.Add("aspect", "media")
	v.Add("verify_token", "token")
	v.Add("callback_url", "http://example.com/rt/user")
	v.Add("object_id", "notifier")
	mres.Body().Equals(v.Encode())
}

func TestRealTimeUnsubscribe(t *testing.T) {
	api, mock := tNewRealTime(t, "my-clientid", "a-secret")

	reg := regmc(`v1\/subscriptions\?client_id=[a-z-]+&client_secret=[a-z-]+&id=\d+$`)
	mock.Register("DELETE", reg, 200, `{"meta": {"code": 201}}`)

	data, _ := api.Unsubscribe("id", "1234")

	assert.Equal(t, data.Meta.Code, 201)
}

func TestRealTimeSubscriptions(t *testing.T) {
	api, mock := tNewRealTime(t, "my-clientid", "a-secret")

	reg := regmc(`v1\/subscriptions\?client_id=[a-z-]+&client_secret=[a-z-]+$`)
	mock.Register("GET", reg, 200, `{"meta": {"code": 202}}`)

	data, _ := api.Subscriptions()

	assert.Equal(t, data.Meta.Code, 202)
	assert.TypeOf(t, "*jsons.Subscriptions", data)
	assert.TypeOf(t, "[]jsons.SubscriptionData", data.Data)
}
