package instagram

import "github.com/nowk/go-instagram/jsons"

// Tags Endpoint
// http://instagram.com/developer/endpoints/tags/
type Tags struct {
	API *Instagram
}

func NewTags(API *Instagram) (tags *Tags) {
	tags = &Tags{
		API: API,
	}

	return
}

// Name - Get information about a tag object.
func (t Tags) Name(tagname string) (data *jsons.Tag, err error) {
	endp := NewEndpoint(TagsName, t.API.query(), tagname)
	err = t.API.Call(endp, &data)

	return
}

// MediaRecent - Get a list of recently tagged media.
func (t Tags) MediaRecent(tagname string, p ...map[string]string) (data *jsons.Medias, err error) {
	endp := NewEndpoint(TagsMediaRecent, t.API.query(), tagname)
	endp.AppendQuery(p...)
	err = t.API.Call(endp, &data)

	return
}

// Search - Search for tags by name.
func (t Tags) Search(q string) (data *jsons.Tags, err error) {
	endp := NewEndpoint(TagsSearch, t.API.query())
	endp.Query.Add("q", q)
	err = t.API.Call(endp, &data)

	return
}
