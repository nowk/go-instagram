package instagram

const (
	_GET    = "GET"
	_POST   = "POST"
	_DELETE = "DELETE"
)

// Users
const (
	_users = iota
	UsersUser
	UsersSelfFeed
	UsersMediaRecent
	UsersSelfMediaLiked
	UsersSearch
)

// Relationships
const (
	_relationships = iota + 10
	RelationshipsFollows
	RelationshipsFollowedBy
	RelationshipsSelfRequestedBy
	RelationshipsRelationship
	RelationshipsPost
)

// Media
const (
	_media = iota + 20
	MediaMedia
	MediaShortcode
	MediaSearch
	MediaPopular
)

// Comments
const (
	_comments = iota + 30
	CommentsComments
	CommentsPost
	CommentsDelete
)

// Likes
const (
	_likes = iota + 40
	LikesLikes
	LikesPost
	LikesDelete
)

// Tags
const (
	_tags = iota + 50
	TagsName
	TagsMediaRecent
	TagsSearch
)

// Locations
const (
	_locations = iota + 60
	LocationsLocation
	LocationsMediaRecent
	LocationsSearch
)

// Geographies
const (
	_geographies = iota + 70
	GeographiesMediaRecent
)

// Real-time
const (
	_realtime = iota + 80
	RealTimeSubscribe
	RealTimeUnsubscribe
	RealTimeSubscriptions
)

// endpoint provides a structure to create METHOD / URL patterns that
// map to an Instagram endpoint
type endpoint struct {
	Method string
	Path   string
}

// Endpoints is an ID => ethod/url pattern map of all the available Instagram
// endpoints
var Endpoints = map[int]endpoint{
	UsersUser:                    endpoint{_GET, "/users/%v"},
	UsersSelfFeed:                endpoint{_GET, "/users/self/feed"},
	UsersMediaRecent:             endpoint{_GET, "/users/%v/media/recent"},
	UsersSelfMediaLiked:          endpoint{_GET, "/users/self/media/liked"},
	UsersSearch:                  endpoint{_GET, "/users/search"},
	RelationshipsFollows:         endpoint{_GET, "/users/%v/follows"},
	RelationshipsFollowedBy:      endpoint{_GET, "/users/%v/followed-by"},
	RelationshipsSelfRequestedBy: endpoint{_GET, "/users/self/requested-by"},
	RelationshipsRelationship:    endpoint{_GET, "/users/%v/relationship"},
	RelationshipsPost:            endpoint{_POST, "/users/%v/relationship"},
	MediaMedia:                   endpoint{_GET, "/media/%v"},
	MediaShortcode:               endpoint{_GET, "/media/shortcode/%v"},
	MediaSearch:                  endpoint{_GET, "/media/search"},
	MediaPopular:                 endpoint{_GET, "/media/popular"},
	CommentsComments:             endpoint{_GET, "/media/%v/comments"},
	CommentsPost:                 endpoint{_POST, "/media/%v/comments"},
	CommentsDelete:               endpoint{_DELETE, "/media/%v/comments/%v"},
	LikesLikes:                   endpoint{_GET, "/media/%v/likes"},
	LikesPost:                    endpoint{_POST, "/media/%v/likes"},
	LikesDelete:                  endpoint{_DELETE, "/media/%v/likes"},
	TagsName:                     endpoint{_GET, "/tags/%v"},
	TagsMediaRecent:              endpoint{_GET, "/tags/%v/media/recent"},
	TagsSearch:                   endpoint{_GET, "/tags/search"},
	LocationsLocation:            endpoint{_GET, "/locations/%v"},
	LocationsMediaRecent:         endpoint{_GET, "/locations/%v/media/recent"},
	LocationsSearch:              endpoint{_GET, "/locations/search"},
	GeographiesMediaRecent:       endpoint{_GET, "/geographies/%v/media/recent"},
	RealTimeSubscribe:            endpoint{_POST, "/subscriptions"},
	RealTimeUnsubscribe:          endpoint{_DELETE, "/subscriptions"},
	RealTimeSubscriptions:        endpoint{_GET, "/subscriptions"},
}
