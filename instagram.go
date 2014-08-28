package instagram

import "net/http"
import "net/url"

const VERSION = "0.0.0"

const (
	_RateLimitErrorStatusCode = 429
)

// Instagram is a struct to create a *client* handling requests out to Instagram
// endpoints
type Instagram struct {
	API

	AccessToken         string
	EnforceSignedHeader *EnforceSignedHeader

	Users         *Users
	Relationships *Relationships
	Media         *Media
	Comments      *Comments
	Likes         *Likes
	Tags          *Tags
	Locations     *Locations
	Geographies   *Geographies
}

// NewClient returns a new *client* with all of the endpoints mounted
func NewClient(accessToken string) (ig *Instagram) {
	ig = &Instagram{
		AccessToken: accessToken,
	}
	ig.HTTPClient = http.DefaultClient

	// mount endpoints
	ig.Users = NewUsers(ig)
	ig.Relationships = NewRelationships(ig)
	ig.Media = NewMedia(ig)
	ig.Comments = NewComments(ig)
	ig.Likes = NewLikes(ig)
	ig.Tags = NewTags(ig)
	ig.Locations = NewLocations(ig)
	ig.Geographies = NewGeographies(ig)

	return
}

// query returns the basic query values required for the final url
// for requests to endpoints that are `client_id` based, this is not required.
func (i Instagram) query() url.Values {
	return url.Values{
		"access_token": {i.AccessToken},
	}
}

// SetSignedHeader sets the X-Insta-Forwarded-For header
func (i *Instagram) SetSignedHeader(secret string, ips ...string) error {
	s := EnforceSignedHeader{
		IPs: ips,
	}

	err := s.Sign(secret)
	if err != nil {
		return err
	}

	i.EnforceSignedHeader = &s

	return nil
}

// Call preps the http.Request from the Endpoint and provides additional headers
// then calls the request on API
func (i *Instagram) Call(ep *Endpoint, data interface{}) error {
	req, err := ep.Request()
	if err != nil {
		return err
	}

	if i.EnforceSignedHeader != nil {
		req.Header.Add("X-Insta-Forwarded-For", i.EnforceSignedHeader.String())
	}

	return i.API.Call(req, &data)
}
