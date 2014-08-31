package instagram_test

import "testing"
import "regexp"
import "github.com/nowk/go-mockhttpc"
import . "github.com/nowk/go-instagram"
import . "github.com/nowk/go-instagram/testing"

const AccessToken = "abcdefg"

func regmc(pat string) *regexp.Regexp {
	return regexp.MustCompile(pat)
}

func tNewInstagram(t *testing.T) (ig *Instagram, mock *mockhttpc.Mock) {
	return NewInstagramT(t, AccessToken)
}

func tNewRealTime(t *testing.T, clientid, clientSecret string) (rt *RealTime, mock *mockhttpc.Mock) {
	return NewRealTimeT(t, clientid, clientSecret)
}
