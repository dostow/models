package models

// DomainRequest a request to create a domain
type DomainRequest struct {
	Domain string `json:"domain" url:"domain"`
}

// DomainResult a domain result
type DomainResult struct {
	Total   int      `json:"total_count"`
	Domains []Domain `json:"data"`
}

// Domain a domain
type Domain struct {
	ID     string `json:"id"`
	Owner  User   `json:"owner"`
	Domain string `json:"domain"`
}
