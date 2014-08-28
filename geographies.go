package instagram

import "net/url"
import "github.com/nowk/go-instagram/jsons"

// Geographies Endpoint
// http://instagram.com/developer/endpoints/geographies/
type Geographies struct {
	API *Instagram
}

func NewGeographies(API *Instagram) (geo *Geographies) {
	geo = &Geographies{
		API: API,
	}

	return
}

// MediaRecentCID - Get recent media from a geography subscription that you
// created.
func (g Geographies) MediaRecentCID(geoid interface{}, clientid string, p ...map[string]string) (data *jsons.Medias, err error) {
	endp := NewEndpoint(GeographiesMediaRecent, url.Values{}, geoid)
	endp.Query.Add("client_id", clientid)
	endp.AppendQuery(p...)
	err = g.API.Call(endp, &data)

	return
}
