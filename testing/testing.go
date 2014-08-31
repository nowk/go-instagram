package testing

import "testing"
import "github.com/nowk/go-mockhttpc"
import . "github.com/nowk/go-instagram"

func NewInstagramT(t *testing.T, accessToken string) (ig *Instagram, mock *mockhttpc.Mock) {
	mock = mockhttpc.NewMock(t)
	ig = &Instagram{
		AccessToken: accessToken,
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

func NewRealTimeT(t *testing.T, clientid, clientSecret string) (rt *RealTime, mock *mockhttpc.Mock) {
	mock = mockhttpc.NewMock(t)
	rt = &RealTime{
		ClientID:     clientid,
		ClientSecret: clientSecret,
	}
	rt.HTTPClient = mock.HTTPClient

	return
}
