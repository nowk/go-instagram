package instagram

import "github.com/nowk/go-instagram/jsons"

// Likes Endpoint
// http://instagram.com/developer/endpoints/likes/
type Likes struct {
	API *Instagram
}

func NewLikes(API *Instagram) (lik *Likes) {
	lik = &Likes{
		API: API,
	}

	return
}

// Likes - Get a list of users who have liked this media.
func (l Likes) Likes(mediaid interface{}) (data *jsons.Likes, err error) {
	endp := NewEndpoint(LikesLikes, l.API.query(), mediaid)
	err = l.API.Call(endp, &data)

	return
}

// Post - Set a like on this media by the currently authenticated user.
func (l Likes) Post(mediaid interface{}) (data *jsons.Like, err error) {
	endp := NewEndpoint(LikesPost, l.API.query(), mediaid)
	err = l.API.Call(endp, &data)

	return
}

// Delete - Remove a like on this media by the currently authenticated user.
func (l Likes) Delete(mediaid interface{}) (data *jsons.Like, err error) {
	endp := NewEndpoint(LikesDelete, l.API.query(), mediaid)
	err = l.API.Call(endp, &data)

	return
}

// Convenience methods

func (l Likes) Like(mediaid interface{}) (data *jsons.Like, err error) {
	return l.Post(mediaid)
}

func (l Likes) Unlike(mediaid interface{}) (data *jsons.Like, err error) {
	return l.Delete(mediaid)
}
