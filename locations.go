package instagram

import "fmt"
import "github.com/nowk/go-instagram/jsons"

// Locations Endpoint
// http://instagram.com/developer/endpoints/locations/
type Locations struct {
	API *Instagram
}

func NewLocations(API *Instagram) (loc *Locations) {
	loc = &Locations{
		API: API,
	}

	return
}

// Location - Get information about a location.
func (l Locations) Location(locid interface{}) (data *jsons.Location, err error) {
	endp := NewEndpoint(LocationsLocation, l.API.query(), locid)
	err = l.API.Call(endp, &data)

	return
}

// MediaRecent - Get a list of recent media objects from a given location.
func (l Locations) MediaRecent(locid interface{}, p ...map[string]string) (data *jsons.Locations, err error) {
	endp := NewEndpoint(LocationsMediaRecent, l.API.query(), locid)
	endp.AppendQuery(p...)
	err = l.API.Call(endp, &data)

	return
}

// Search - Search for a location by geographic coordinate.
func (l Locations) Search(lat, lng float64, p ...map[string]string) (data *jsons.Locations, err error) {
	endp := NewEndpoint(LocationsSearch, l.API.query())
	endp.Query.Add("lat", fmt.Sprintf("%.6f", lat))
	endp.Query.Add("lng", fmt.Sprintf("%.6f", lng))
	endp.AppendQuery(p...)
	err = l.API.Call(endp, &data)

	return
}
