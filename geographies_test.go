package instagram_test

import "testing"
import "github.com/nowk/assert"

func TestGeographiesMediaRecentCID(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/geographies\/\d+\/media\/recent\?client_id=clientid&count=\d+&min_id=\d+$`)
	mock.Expect("GET", reg).Respond(200, `{"meta": {"code": 200}}`)

	data, _ := api.Geographies.MediaRecentCID(12345, "clientid", map[string]string{
		"count":  "50",
		"min_id": "12345678",
	})

	assert.Equal(t, data.Meta.Code, 200)
	assert.TypeOf(t, "*jsons.Medias", data)
	assert.TypeOf(t, "[]jsons.MediaData", data.Data)
}
