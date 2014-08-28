package instagram

import "fmt"
import "io"
import "net/http"
import "net/url"

const (
	API_PROTOCOL = "https"
	API_HOST     = "api.instagram.com"
	API_VERSION  = "v1"
)

// Endpoint is a struct to build a request ready for an Instagram API
// endpoint
type Endpoint struct {
	Method string
	Path   string
	Query  url.Values
	Params []interface{}
	Body   io.Reader
	key    int
}

func NewEndpoint(key int, query url.Values, params ...interface{}) (endpoint *Endpoint) {
	ep := Endpoints[key]

	endpoint = &Endpoint{
		Method: ep.Method,
		Path:   ep.Path,
		Query:  query,
		Params: params,
		key:    key,
	}

	return
}

// GetKey returns the endpoint const ID used to build the endpoint
func (e Endpoint) GetKey() int {
	return e.key
}

// URL builds a url.Url from the properties of the given endpoint
func (e Endpoint) URL() *url.URL {
	path := e.Path
	if len(e.Params) > 0 {
		path = fmt.Sprintf(e.Path, e.Params...)
	}

	u := &url.URL{
		Scheme:   API_PROTOCOL,
		Host:     API_HOST,
		Path:     fmt.Sprintf("/%s%s", API_VERSION, path),
		RawQuery: e.Query.Encode(),
	}

	return u
}

// Request builds a request from the Endpoint and returns an http.Request
func (e Endpoint) Request() (*http.Request, error) {
	return http.NewRequest(e.Method, e.URL().String(), e.Body)
}

// AppendQuery appends e.Query from a ...map[string]string
func (e *Endpoint) AppendQuery(p ...map[string]string) {
	for _, par := range p {
		for k, v := range par {
			e.Query.Add(k, v)
		}
	}
}
