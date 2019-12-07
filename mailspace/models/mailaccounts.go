package models

// MailAccount a mail account
type MailAccount struct {
	CreatedAt string `json:"created_at"`
	Space     string `json:"space"`
	Email     string `json:"email"`
	ID        string `json:"id"`
	Owner     string `json:"owner"`
	Quota     int    `json:"quota"`
	Status    string `json:"status"`
}

// MailAccountResult result of mail accounts
type MailAccountResult struct {
	Count      int           `json:"count"`
	Data       []MailAccount `json:"data"`
	TotalCount int           `json:"total_count"`
}
