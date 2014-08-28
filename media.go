package instagram

import "fmt"
import "github.com/nowk/go-instagram/jsons"

// Media Endpoint
// http://instagram.com/developer/endpoints/media/
type Media struct {
	API *Instagram
}

func NewMedia(API *Instagram) (media *Media) {
	media = &Media{
		API: API,
	}

	return
}

// Media - Get information about a media object.
func (m Media) Media(mediaid interface{}) (data *jsons.Media, err error) {
	endp := NewEndpoint(MediaMedia, m.API.query(), mediaid)
	err = m.API.Call(endp, &data)

	return
}

// Shortcode - This endpoint returns the same response as `Media()`.
func (m Media) Shortcode(shortcode interface{}) (data *jsons.Media, err error) {
	endp := NewEndpoint(MediaShortcode, m.API.query(), shortcode)
	err = m.API.Call(endp, &data)

	return
}

// Search - Search for media in a given area.
func (m Media) Search(lat, lng float64, p ...map[string]string) (data *jsons.Medias, err error) {
	endp := NewEndpoint(MediaSearch, m.API.query())
	endp.Query.Add("lat", fmt.Sprintf("%.6f", lat))
	endp.Query.Add("lng", fmt.Sprintf("%.6f", lng))
	endp.AppendQuery(p...)
	err = m.API.Call(endp, &data)

	return
}

// Popular - Get a list of what media is most popular at the moment.
func (m Media) Popular() (data *jsons.Medias, err error) {
	endp := NewEndpoint(MediaPopular, m.API.query())
	err = m.API.Call(endp, &data)

	return
}
