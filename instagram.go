package instagram

import "encoding/json"
import "io"
import "net/http"
import "net/url"
import . "github.com/nowk/go-httpclienti"

const VERSION = "0.0.0"

const (
	_RateLimitErrorStatusCode = 429
)

// Instagram is a struct to create a *client* handling requests out to Instagram
// endpoints
type Instagram struct {
	AccessToken         string
	HTTPClient          HTTPClient
	RateLimit           *RateLimit
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
		HTTPClient:  http.DefaultClient,
	}

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

// Call takes a Endpoint and *data* struct, makes the request and maps the
// returned data back to give *data* struct
func (i *Instagram) Call(ep *Endpoint, data interface{}) error {
	resp, err := i.request(ep)
	if err != nil {
		return err
	}

	i.RateLimit = NewRateLimit(&resp.Header)

	err = i.decode(resp.Body, &data)
	if err != err {
		return err
	}

	// return RateLimitError, as the response will still be OK
	if resp.StatusCode == _RateLimitErrorStatusCode {
		return RateLimitError{}
	}

	return err
}

// request requests an http.Request to the given Endpoint
func (i Instagram) request(ep *Endpoint) (*http.Response, error) {
	req, err := ep.Request()
	if err != nil {
		return nil, err
	}

	if i.EnforceSignedHeader != nil {
		req.Header.Add("X-Insta-Forwarded-For", i.EnforceSignedHeader.String())
	}

	return i.HTTPClient.Do(req)
}

// decode decodes the give io.ReadCloser (http.Response.Body) to the
// *data* struct
func (i Instagram) decode(body io.ReadCloser, data interface{}) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(&data)
}
