package instagram

import "fmt"
import "net/url"
import "github.com/nowk/go-instagram/jsons"

// Users Endpoint
// http://instagram.com/developer/endpoints/users/
type Users struct {
	API *Instagram
}

func NewUsers(API *Instagram) (users *Users) {
	users = &Users{
		API: API,
	}

	return
}

// User - Get basic information about a user.
func (u Users) User(userid interface{}) (data *jsons.User, err error) {
	endp := NewEndpoint(UsersUser, u.API.query(), userid)
	err = u.API.Call(endp, &data)

	return
}

// SelfFeed - See the authenticated user's feed.
func (u Users) SelfFeed(p ...map[string]string) (data *jsons.Medias, err error) {
	endp := NewEndpoint(UsersSelfFeed, u.API.query())
	endp.AppendQuery(p...)
	err = u.API.Call(endp, &data)

	return
}

// MediaRecent - Get the most recent media published by a user.
func (u Users) MediaRecent(userid interface{}, p ...map[string]string) (data *jsons.Medias, err error) {
	endp := NewEndpoint(UsersMediaRecent, u.API.query(), userid)
	endp.AppendQuery(p...)
	err = u.API.Call(endp, &data)

	return
}

// MediaRecentCID - The functionality is the same as `MediaRecent()`, but use
// `client_id` instead of access_token.
func (u Users) MediaRecentCID(clientid interface{}, p ...map[string]string) (data *jsons.Medias, err error) {
	endp := NewEndpoint(UsersMediaRecent, url.Values{}, clientid)
	endp.Query.Add("client_id", fmt.Sprintf("%v", clientid))
	endp.AppendQuery(p...)
	err = u.API.Call(endp, &data)

	return
}

// SelfMediaLiked - See the authenticated user's list of media they've liked.
func (u Users) SelfMediaLiked(p ...map[string]string) (data *jsons.Medias, err error) {
	endp := NewEndpoint(UsersSelfMediaLiked, u.API.query())
	endp.AppendQuery(p...)
	err = u.API.Call(endp, &data)

	return
}

// Search - Search for a user by name.
func (u Users) Search(q string, p ...map[string]string) (data *jsons.Users, err error) {
	endp := NewEndpoint(UsersSearch, u.API.query())
	endp.Query.Add("q", q)
	endp.AppendQuery(p...)
	err = u.API.Call(endp, &data)

	return
}
