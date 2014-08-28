package instagram

import "fmt"
import "strings"
import "github.com/nowk/go-instagram/jsons"

// Comments Endpoint
// http://instagram.com/developer/endpoints/comments/
type Comments struct {
	API *Instagram
}

func NewComments(API *Instagram) (com *Comments) {
	com = &Comments{
		API: API,
	}

	return
}

// Comments - Get a full list of comments on a media object.
func (c Comments) Comments(mediaid interface{}) (data *jsons.Comments, err error) {
	endp := NewEndpoint(CommentsComments, c.API.query(), mediaid)
	err = c.API.Call(endp, &data)

	return
}

// Post - Create a comment on a media object.
// NOTE you will be required to request access to the Comments POST endpoint in
// order to get access to this feature
func (c Comments) Post(mediaid interface{}, text string) (data *jsons.Comment, err error) {
	endp := NewEndpoint(CommentsPost, c.API.query(), mediaid)
	endp.Body = strings.NewReader(fmt.Sprintf("text=%s", text))
	err = c.API.Call(endp, &data)

	return
}

// Delete - Remove a comment either on the authenticated user's media object or
// authored by the authenticated user.
func (c Comments) Delete(mediaid, commentid interface{}) (data *jsons.Comment, err error) {
	endp := NewEndpoint(CommentsDelete, c.API.query(), mediaid, commentid)
	err = c.API.Call(endp, &data)

	return
}
