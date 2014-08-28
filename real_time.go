package instagram

import "net/http"
import "net/url"
import "strings"
import "github.com/nowk/go-instagram/jsons"

// RealTime creates a Client for executing RealTime only related requests
// http://instagram.com/developer/realtime/
type RealTime struct {
	API

	ClientID     string
	ClientSecret string
}

func NewRealTime(clientid, clientSecret string) (rt *RealTime) {
	rt = &RealTime{
		ClientID:     clientid,
		ClientSecret: clientSecret,
	}
	rt.HTTPClient = http.DefaultClient

	return
}

// query returns the basic query values required for the final request
func (r RealTime) query() url.Values {
	return url.Values{
		"client_id":     {r.ClientID},
		"client_secret": {r.ClientSecret},
	}
}

// Call preps the http.Request from the endpoint and calls the request on API
func (r *RealTime) Call(ep *Endpoint, data interface{}) error {
	req, err := ep.Request()
	if err != nil {
		return err
	}

	return r.API.Call(req, &data)
}

// Subscribe requests a subscription
func (r RealTime) Subscribe(object, token, callbackURL string, p ...map[string]string) (data *jsons.Subscriptions, err error) {
	query := r.query()
	query.Add("object", object)
	query.Add("aspect", "media")
	query.Add("verify_token", token)
	query.Add("callback_url", callbackURL)

	for _, params := range p {
		for k, v := range params {
			query.Add(k, v)
		}
	}

	endp := NewEndpoint(RealTimeSubscribe, nil)
	endp.Body = strings.NewReader(query.Encode())
	err = r.Call(endp, &data)

	return
}

// Unsubscribe deletes an existing subscription
func (r RealTime) Unsubscribe(k, v string) (data *jsons.Subscriptions, err error) {
	query := r.query()
	query.Add(k, v)
	endp := NewEndpoint(RealTimeUnsubscribe, query)
	err = r.Call(endp, &data)

	return
}

// Subscriptions returns a list of current subscriptions
func (r RealTime) Subscriptions() (data *jsons.Subscriptions, err error) {
	endp := NewEndpoint(RealTimeSubscriptions, r.query())
	err = r.Call(endp, &data)

	return
}
