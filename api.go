package instagram

import "encoding/json"
import "io"
import "net/http"
import . "github.com/nowk/go-httpclienti"

// API is to be *inherited* by other *clients* and contains the core method for
// making HTTP requests and json encoding back to the final data struct
type API struct {
	HTTPClient HTTPClient
	RateLimit  *RateLimit
}

// Call takes an http request and decodes the json data back to the given data
// struct
func (api *API) Call(req *http.Request, data interface{}) error {
	resp, err := api.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	api.RateLimit = NewRateLimit(&resp.Header)

	err = decode(resp.Body, &data)
	if err != err {
		return err
	}

	// return RateLimitError, as the response will still be OK
	if resp.StatusCode == _RateLimitErrorStatusCode {
		return RateLimitError{}
	}

	return nil
}

// decode decodes the give io.ReadCloser (http.Response.Body) to the
// *data* struct
func decode(body io.ReadCloser, data interface{}) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(&data)
}
