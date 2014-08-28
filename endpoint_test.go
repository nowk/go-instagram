package instagram_test

import "net/url"
import "testing"
import "github.com/nowk/assert"
import . "github.com/nowk/go-instagram"

var query = url.Values{
	"access_token": {"abcdefg"},
}

func TestEndpointUrl(t *testing.T) {
	table := map[int]string{
		UsersSelfFeed:       "/v1/users/self/feed",
		UsersSelfMediaLiked: "/v1/users/self/media/liked",
	}

	for k, v := range table {
		ep := NewEndpoint(k, query)
		assert.Equal(t, ep.URL(), &url.URL{
			Scheme:   "https",
			Host:     "api.instagram.com",
			Path:     v,
			RawQuery: "access_token=abcdefg",
		})
	}
}

func TestEndpointUrlWithParams(t *testing.T) {
	table := map[int]string{
		UsersUser:        "/v1/users/12345",
		UsersMediaRecent: "/v1/users/12345/media/recent",
	}

	for k, v := range table {
		ep := NewEndpoint(k, query, 12345)
		assert.Equal(t, ep.URL(), &url.URL{
			Scheme:   "https",
			Host:     "api.instagram.com",
			Path:     v,
			RawQuery: "access_token=abcdefg",
		})
	}
}

func TestEndpointWithQuery(t *testing.T) {
	table := map[int]string{
		UsersSearch: "/v1/users/search",
	}

	query.Add("q", "foo bar")

	for k, v := range table {
		ep := NewEndpoint(k, query)
		assert.Equal(t, ep.URL(), &url.URL{
			Scheme:   "https",
			Host:     "api.instagram.com",
			Path:     v,
			RawQuery: "access_token=abcdefg&q=foo+bar",
		})
	}
}
