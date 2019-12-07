package models

// Profile defines a profile
type Profile struct {
	Address   string `json:"address"`
	Address2  string `json:"address2"`
	CreatedAt string `json:"created_at"`
	ID        string `json:"id"`
	FullName  string `json:"fullName"`
	Owner     string `json:"owner"`
	Phone     string `json:"phone"`
}

// ProfileRequest defines a request for a profile
type ProfileRequest struct {
	Owner    string `json:"owner"`
	FullName string `json:"fullName,omitempty"`
	Address  string `json:"address,omitempty"`
	Address2 string `json:"address2,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

// ProfilesResult a result of profiles
type ProfilesResult struct {
	TotalCount int       `json:"total_count"`
	Count      int       `json:"count"`
	Data       []Profile `json:"data"`
}
