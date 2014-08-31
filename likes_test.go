package instagram_test

import "testing"
import "github.com/nowk/assert"

func TestLikesLikes(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/media\/\d+\/likes\?access_token=\w+$`)
	mock.Expect("GET", reg).Respond(200, `{"meta": {"code": 200}}`)

	data, _ := api.Likes.Likes(12345)

	assert.Equal(t, data.Meta.Code, 200)
	assert.TypeOf(t, "*jsons.Likes", data)
	assert.TypeOf(t, "[]jsons.LikeData", data.Data)
}

func TestLikesPost(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/media\/\d+\/likes\?access_token=\w+$`)
	mock.Expect("POST", reg).Respond(200, `{"meta": {"code": 201}}`)

	data, _ := api.Likes.Post(12345)

	assert.Equal(t, data.Meta.Code, 201)
	assert.TypeOf(t, "*jsons.Like", data)
	assert.TypeOf(t, "jsons.LikeData", data.Data)
}

func TestLikesDelete(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/media\/\d+\/likes\?access_token=\w+$`)
	mock.Expect("DELETE", reg).Respond(200, `{"meta": {"code": 202}}`)

	data, _ := api.Likes.Delete(12345)

	assert.Equal(t, data.Meta.Code, 202)
	assert.TypeOf(t, "*jsons.Like", data)
	assert.TypeOf(t, "jsons.LikeData", data.Data)
}
