package models

// Server is a space server
type Server struct {
	CreatedAt  string `json:"created_at"`
	Dc         string `json:"dc"`
	Hostname   string `json:"hostname"`
	ID         string `json:"id"`
	IP         string `json:"ip"`
	ModifiedAt string `json:"modified_at"`
	Name       string `json:"name"`
}

// Servers a list of servers
type Servers struct {
	TotalCount int64    `json:"total_count"`
	Data       []Server `json:"data"`
}

// Space description of a space
type Space struct {
	CreatedAt   string    `json:"created_at"`
	Domain      string    `json:"domain"`
	DomainSpace string    `json:"domainSpace"`
	ID          string    `json:"id"`
	ModifiedAt  string    `json:"modified_at"`
	Owner       string    `json:"owner"`
	Status      string    `json:"status"`
	Server      Server    `json:"server,omitempty"`
	Plan        Plan      `json:"plan"`
	Invoice     []Invoice `json:"invoice_list"`
	Type        string    `json:"type"`
}

// IsActive check if a space is active
func (space *Space) IsActive() bool {
	if space.Status != "ACTIVE" {
		return false
	}
	return true
}

// IsPaid check if space is paid for
func (space *Space) IsPaid() bool {
	for _, invoice := range space.Invoice {
		if invoice.AmountDue != invoice.AmountPaid {
			return false
		}
	}
	return true
}

// CanActivate check if space can be activated
func (space *Space) CanActivate() bool {
	return space.IsPaid()
}
