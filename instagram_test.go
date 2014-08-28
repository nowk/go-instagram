package instagram_test

import "testing"
import "github.com/nowk/assert"

func TestRateLimitError(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/users\/self\/feed\?access_token=\w+$`)
	mock.Register("GET", reg, 429, `{"meta": {"code": 429}}`)

	data, err := api.Users.SelfFeed()

	assert.Equal(t, data.Meta.Code, 429)
	assert.Equal(t, err.Error(), "The maximum number of requests per hour has been exceeded.")
	assert.TypeOf(t, "instagram.RateLimitError", err)
}

func TestRateLimit(t *testing.T) {
	api, mock := tNewInstagram(t)

	reg := regmc(`v1\/users\/self\/feed\?access_token=\w+$`)
	resp := mock.Register("GET", reg, 200, `{"meta": {"code": 200}}`)
	resp.Header.Add("X-Ratelimit-Remaining", "3000")
	resp.Header.Add("X-Ratelimit-Limit", "5000")

	api.Users.SelfFeed()

	assert.Equal(t, api.RateLimit.Remaining, int64(3000))
	assert.Equal(t, api.RateLimit.Limit, int64(5000))
}
