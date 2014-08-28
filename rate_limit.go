package instagram

import "net/http"
import "strconv"

// RateLimit contains rate limit data returned from the response header
type RateLimit struct {
	Remaining int64 `header:"X-Ratelimit-Remaining"`
	Limit     int64 `header:"X-Ratelimit-Limit"`
}

func NewRateLimit(head *http.Header) (rate *RateLimit) {
	// TODO should we just be passing nil on err?

	r, err := parseInt(head.Get("X-RateLimit-Remaining"))
	if err != nil {
		return rate
	}

	l, err := parseInt(head.Get("X-RateLimit-Limit"))
	if err != nil {
		return rate
	}

	rate = &RateLimit{
		Remaining: r,
		Limit:     l,
	}

	return
}

// RateLimitError returns an error object specifically for rate limits
type RateLimitError struct{}

func (e RateLimitError) Error() string {
	return "The maximum number of requests per hour has been exceeded."
}

// parseInt converts a string to int64
func parseInt(str string) (int64, error) {
	return strconv.ParseInt(str, 0, 64)
}
