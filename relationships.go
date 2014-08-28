package instagram

import "fmt"
import "strings"
import "github.com/nowk/go-instagram/jsons"

// Relationships Endpoint
// http://instagram.com/developer/endpoints/relationships/
type Relationships struct {
	API *Instagram
}

func NewRelationships(API *Instagram) (rel *Relationships) {
	rel = &Relationships{
		API: API,
	}

	return
}

// Follows - Get the list of users this user follows.
func (r Relationships) Follows(userid interface{}) (data *jsons.Users, err error) {
	endp := NewEndpoint(RelationshipsFollows, r.API.query(), userid)
	err = r.API.Call(endp, &data)

	return
}

// FollowedBy - Get the list of users this user is followed by.
func (r Relationships) FollowedBy(userid interface{}) (data *jsons.Users, err error) {
	endp := NewEndpoint(RelationshipsFollowedBy, r.API.query(), userid)
	err = r.API.Call(endp, &data)

	return
}

// SelfRequestedBy - List the users who have requested this user's permission
// to follow.
func (r Relationships) SelfRequestedBy() (data *jsons.Users, err error) {
	endp := NewEndpoint(RelationshipsSelfRequestedBy, r.API.query())
	err = r.API.Call(endp, &data)

	return
}

// Relationship - Get information about a relationship to another user.
func (r Relationships) Relationship(userid interface{}) (data *jsons.Relationship, err error) {
	endp := NewEndpoint(RelationshipsRelationship, r.API.query(), userid)
	err = r.API.Call(endp, &data)

	return
}

// Post - Modify the relationship between the current user and the target user.
func (r Relationships) Post(userid interface{}, action string) (data *jsons.Relationship, err error) {
	endp := NewEndpoint(RelationshipsPost, r.API.query(), userid)
	endp.Body = strings.NewReader(fmt.Sprintf("action=%s", action))
	err = r.API.Call(endp, &data)

	return
}

// Convenience methods

func (r Relationships) Follow(userid interface{}) (data *jsons.Relationship, err error) {
	return r.Post(userid, "follow")
}

func (r Relationships) Unfollow(userid interface{}) (data *jsons.Relationship, err error) {
	return r.Post(userid, "unfollow")
}

func (r Relationships) Block(userid interface{}) (data *jsons.Relationship, err error) {
	return r.Post(userid, "block")
}

func (r Relationships) Unblock(userid interface{}) (data *jsons.Relationship, err error) {
	return r.Post(userid, "unblock")
}

func (r Relationships) Approve(userid interface{}) (data *jsons.Relationship, err error) {
	return r.Post(userid, "approve")
}

func (r Relationships) Ignore(userid interface{}) (data *jsons.Relationship, err error) {
	return r.Post(userid, "ignore")
}
