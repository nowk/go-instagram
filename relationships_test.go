package instagram_test

import "testing"
import "github.com/nowk/assert"

func TestRelationshipsFollows(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/users\/\d+\/follows\?access_token=\w+$`)
	mock.Register("GET", reg, 200, `{"meta": {"code": 200}}`)

	data, _ := api.Relationships.Follows(123)

	assert.Equal(t, data.Meta.Code, 200)
	assert.TypeOf(t, "*jsons.Users", data)
	assert.TypeOf(t, "[]jsons.UserData", data.Data)
}

func TestRelationshipsFollowedBy(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/users\/\d+\/followed-by\?access_token=\w+$`)
	mock.Register("GET", reg, 200, `{"meta": {"code": 201}}`)

	data, _ := api.Relationships.FollowedBy(123)

	assert.Equal(t, data.Meta.Code, 201)
	assert.TypeOf(t, "*jsons.Users", data)
	assert.TypeOf(t, "[]jsons.UserData", data.Data)
}

func TestRelationshipsSelfRequestedBy(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/users\/self\/requested-by\?access_token=\w+$`)
	mock.Register("GET", reg, 200, `{"meta": {"code": 202}}`)

	data, _ := api.Relationships.SelfRequestedBy()

	assert.Equal(t, data.Meta.Code, 202)
	assert.TypeOf(t, "*jsons.Users", data)
	assert.TypeOf(t, "[]jsons.UserData", data.Data)
}

func TestRelationshipsRelationship(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/users\/\d+\/relationship\?access_token=\w+$`)
	mock.Register("GET", reg, 200, `{"meta": {"code": 203}}`)

	data, _ := api.Relationships.Relationship(12345)

	assert.Equal(t, data.Meta.Code, 203)
	assert.TypeOf(t, "*jsons.Relationship", data)
	assert.TypeOf(t, "jsons.RelationshipData", data.Data)
}

func TestRelationshipsPost(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/users\/\d+\/relationship\?access_token=\w+$`)
	_, mres := mock.Register("POST", reg, 200, `{"meta": {"code": 204}}`)

	data, _ := api.Relationships.Post(12345, "follow")

	assert.Equal(t, data.Meta.Code, 204)
	assert.TypeOf(t, "*jsons.Relationship", data)
	assert.TypeOf(t, "jsons.RelationshipData", data.Data)

	mres.RequestBody().Equals("action=follow")
}
