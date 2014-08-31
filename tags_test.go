package instagram_test

import "testing"
import "github.com/nowk/assert"

func TestTagsName(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/tags\/\w+\?access_token=\w+$`)
	mock.Expect("GET", reg).Respond(200, `{"meta": {"code": 200}}`)

	data, _ := api.Tags.Name("Foo")

	assert.Equal(t, data.Meta.Code, 200)
	assert.TypeOf(t, "*jsons.Tag", data)
	assert.TypeOf(t, "jsons.TagData", data.Data)
}

func TestTagsMediaRecent(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/tags\/\w+\/media\/recent\?access_token=\w+&count=\d+$`)
	mock.Expect("GET", reg).Respond(200, `{"meta": {"code": 201}}`)

	data, _ := api.Tags.MediaRecent("Foo", map[string]string{
		"count": "50",
	})

	assert.Equal(t, data.Meta.Code, 201)
	assert.TypeOf(t, "*jsons.Medias", data)
	assert.TypeOf(t, "[]jsons.MediaData", data.Data)
}

func TestTagsSearch(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/tags\/search\?access_token=\w+&q=\w+$`)
	mock.Expect("GET", reg).Respond(200, `{"meta": {"code": 202}}`)

	data, _ := api.Tags.Search("Foo")

	assert.Equal(t, data.Meta.Code, 202)
	assert.TypeOf(t, "*jsons.Tags", data)
	assert.TypeOf(t, "[]jsons.TagData", data.Data)
}
