package instagram_test

import "testing"
import "regexp"
import . "github.com/nowk/go-instagram"
import . "github.com/nowk/go-instagram/testing"
import "gopkg.in/nowk/go-netmock.v0"

const AccessToken = "abcdefg"

func regmc(pat string) *regexp.Regexp {
	return regexp.MustCompile(pat)
}

func tNewInstagram(t *testing.T) (ig *Instagram, mock *netmock.Mock) {
	return NewInstagramT(t, AccessToken)
}

func tNewRealTime(t *testing.T, clientid, clientSecret string) (rt *RealTime, mock *netmock.Mock) {
	return NewRealTimeT(t, clientid, clientSecret)
}
