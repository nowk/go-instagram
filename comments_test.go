package instagram_test

import "testing"
import "github.com/nowk/assert"

func TestCommentsComments(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/media\/\d+\/comments\?access_token=\w+$`)
	mock.Register("GET", reg, 200, `{"meta": {"code": 200}}`)

	data, _ := api.Comments.Comments(12345)

	assert.Equal(t, data.Meta.Code, 200)
	assert.TypeOf(t, "*jsons.Comments", data)
	assert.TypeOf(t, "[]jsons.CommentData", data.Data)
}

func TestCommentsPost(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/media\/\d+\/comments\?access_token=\w+$`)
	_, mres := mock.Register("POST", reg, 200, `{"meta": {"code": 201}}`)

	data, _ := api.Comments.Post(12345, "Hello World")

	assert.Equal(t, data.Meta.Code, 201)
	assert.TypeOf(t, "*jsons.Comment", data)
	assert.TypeOf(t, "jsons.CommentData", data.Data)

	mres.RequestBody().Equals("text=Hello World")
}

func TestCommentsDelete(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/media\/\d+\/comments\/\d+\?access_token=\w+$`)
	mock.Register("DELETE", reg, 200, `{"meta": {"code": 202}}`)

	data, _ := api.Comments.Delete(12345, 67890)

	assert.Equal(t, data.Meta.Code, 202)
	assert.TypeOf(t, "*jsons.Comment", data)
	assert.TypeOf(t, "jsons.CommentData", data.Data)
}
