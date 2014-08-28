package instagram_test

import "testing"
import "regexp"
import "github.com/nowk/go-mockhttpc"
import . "github.com/nowk/go-instagram"

const AccessToken = "abcdefg"

func regmc(pat string) *regexp.Regexp {
	return regexp.MustCompile(pat)
}

func tNewInstagram(t *testing.T) (ig *Instagram, mock *mockhttpc.Mock) {
	mock = mockhttpc.NewMock(t)
	ig = &Instagram{
		AccessToken: AccessToken,
	}
	ig.HTTPClient = mock.HTTPClient

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

func tNewRealTime(t *testing.T, clientid, clientSecret string) (rt *RealTime, mock *mockhttpc.Mock) {
	mock = mockhttpc.NewMock(t)
	rt = &RealTime{
		ClientID:     clientid,
		ClientSecret: clientSecret,
	}
	rt.HTTPClient = mock.HTTPClient

	return
}
