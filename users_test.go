package instagram_test

import "testing"
import "github.com/nowk/assert"

func TestUsersSelfFeed(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/users\/self\/feed\?access_token=\w+&count=\d+&max_id=\d+$`)
	mock.Register("GET", reg, 200, `{"meta": {"code": 200}}`)

	data, _ := api.Users.SelfFeed(map[string]string{
		"count":  "50",
		"max_id": "12345678",
	})

	assert.Equal(t, data.Meta.Code, 200)
	assert.TypeOf(t, "*jsons.Medias", data)
	assert.TypeOf(t, "[]jsons.MediaData", data.Data)
}

func TestUsersUser(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/users\/\d+\?access_token=`)
	mock.Register("GET", reg, 200, `{"meta": {"code": 201}}`)

	data, _ := api.Users.User(1234)

	assert.Equal(t, data.Meta.Code, 201)
	assert.TypeOf(t, "*jsons.User", data)
	assert.TypeOf(t, "jsons.UserData", data.Data)
}

func TestUsersMediaRecent(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/users\/\d+\/media\/recent\?access_token=\w+&count=\d+&min_id=\d+$`)
	mock.Register("GET", reg, 200, `{"meta": {"code": 202}}`)

	data, _ := api.Users.MediaRecent(1234, map[string]string{
		"count":  "50",
		"min_id": "12345678",
	})

	assert.Equal(t, data.Meta.Code, 202)
	assert.TypeOf(t, "*jsons.Medias", data)
	assert.TypeOf(t, "[]jsons.MediaData", data.Data)
}

func TestUsersMediaRecentCID(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/users\/\d+\/media\/recent\?client_id=clientid&count=\d+&min_id=\d+$`)
	mock.Register("GET", reg, 200, `{"meta": {"code": 203}}`)

	data, _ := api.Users.MediaRecentCID(1234, "clientid", map[string]string{
		"count":  "50",
		"min_id": "12345678",
	})

	assert.Equal(t, data.Meta.Code, 203)
	assert.TypeOf(t, "*jsons.Medias", data)
	assert.TypeOf(t, "[]jsons.MediaData", data.Data)
}

func TestUsersSelfMediaLiked(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/users\/self\/media\/liked\?access_token=\w+&count=\d+$`)
	mock.Register("GET", reg, 200, `{"meta": {"code": 204}}`)

	data, _ := api.Users.SelfMediaLiked(map[string]string{
		"count": "50",
	})

	assert.Equal(t, data.Meta.Code, 204)
	assert.TypeOf(t, "*jsons.Medias", data)
	assert.TypeOf(t, "[]jsons.MediaData", data.Data)
}

func TestUsersSearch(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/users\/search\?access_token=\w+&count=\d+&q=foo\+bar$`)
	mock.Register("GET", reg, 200, `{"meta": {"code": 205}}`)

	data, _ := api.Users.Search("foo bar", map[string]string{
		"count": "50",
	})

	assert.Equal(t, data.Meta.Code, 205)
	assert.TypeOf(t, "*jsons.Users", data)
	assert.TypeOf(t, "[]jsons.UserData", data.Data)
}
