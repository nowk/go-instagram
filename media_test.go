package instagram_test

import "testing"
import "github.com/nowk/assert"

func TestMediaMedia(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/media\/\d+\?access_token=\w+$`)
	mock.Register("GET", reg, 200, `{"meta": {"code": 200}}`)

	data, _ := api.Media.Media(1234)

	assert.Equal(t, 200, data.Meta.Code)
	assert.TypeOf(t, "*jsons.Media", data)
	assert.TypeOf(t, "jsons.MediaData", data.Data)
}

func TestMediaShortcode(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/media\/shortcode\/\d+\?access_token=\w+$`)
	mock.Register("GET", reg, 200, `{"meta": {"code": 201}}`)

	data, _ := api.Media.Shortcode(1234)

	assert.Equal(t, 201, data.Meta.Code)
	assert.TypeOf(t, "*jsons.Media", data)
	assert.TypeOf(t, "jsons.MediaData", data.Data)
}

func TestMediaSearch(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/media\/search\?access_token=\w+&distance=\d+&lat=(\d+\.\d+)&lng=(\d+\.\d+)$`)
	mock.Register("GET", reg, 200, `{"meta": {"code": 202}}`)

	data, _ := api.Media.Search(40.7127, 74.0059, map[string]string{
		"distance": "10000",
	})

	assert.Equal(t, 202, data.Meta.Code)
	assert.TypeOf(t, "*jsons.Medias", data)
	assert.TypeOf(t, "[]jsons.MediaData", data.Data)
}

func TestMediaPopular(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/media\/popular\?access_token=\w+$`)
	mock.Register("GET", reg, 200, `{"meta": {"code": 203}}`)

	data, _ := api.Media.Popular()

	assert.Equal(t, 203, data.Meta.Code)
	assert.TypeOf(t, "*jsons.Medias", data)
	assert.TypeOf(t, "[]jsons.MediaData", data.Data)
}
