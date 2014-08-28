package jsons

type User struct {
	Meta Meta     `json:"meta"`
	Data UserData `json:"data"`
}

type Users struct {
	Meta       Meta       `json:"meta"`
	Pagination Pagination `json:"pagination"`
	Data       []UserData `json:"data"`
}

type UserData struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	FullName       string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
	Bio            string `json:"bio"`
	Website        string `json:"website"`

	Counts struct {
		Media      int `json:"media"`
		Follows    int `json:"follows"`
		FollowedBy int `json:"followed_by"`
	} `json:"counts"`
}
