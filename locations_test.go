package instagram_test

import "testing"
import "github.com/nowk/assert"

func TestLocationsLocation(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/locations\/\d+\?access_token=\w+$`)
	mock.Register("GET", reg, 200, `{"meta": {"code": 200}}`)

	data, _ := api.Locations.Location(12345)

	assert.Equal(t, data.Meta.Code, 200)
	assert.TypeOf(t, "*jsons.Location", data)
	assert.TypeOf(t, "jsons.LocationData", data.Data)
}

func TestLocationsMediaRecent(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/locations\/\d+\/media\/recent\?access_token=\w+&min_id=\d+$`)
	mock.Register("GET", reg, 200, `{"meta": {"code": 201}}`)

	data, _ := api.Locations.MediaRecent(12345, map[string]string{
		"min_id": "67890",
	})
	assert.Equal(t, data.Meta.Code, 201)
	assert.TypeOf(t, "*jsons.Locations", data)
	assert.TypeOf(t, "[]jsons.LocationData", data.Data)
}

func TestLocationsSearch(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/locations\/search\?access_token=\w+&distance=\d+&lat=(\d+\.\d+)&lng=(\d+\.\d+)$`)
	mock.Register("GET", reg, 200, `{"meta": {"code": 202}}`)

	data, _ := api.Locations.Search(40.7127, 74.0059, map[string]string{
		"distance": "5000",
	})
	assert.Equal(t, data.Meta.Code, 202)
	assert.TypeOf(t, "*jsons.Locations", data)
	assert.TypeOf(t, "[]jsons.LocationData", data.Data)
}
