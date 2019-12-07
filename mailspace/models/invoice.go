package models

// UpdateInvoice update invoice model
type UpdateInvoice struct {
	AmountDue  int    `json:"amountDue,omitempty"`
	AmountPaid int    `json:"amountPaid,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	DueBy      string `json:"dueBy,omitempty"`
	Domain     string `json:"domain,omitempty"`
	Owner      string `json:"owner,omitempty"`
	Plan       string `json:"plan,omitempty"`
	Status     string `json:"status,omitempty"`
	Space      string `json:"space,omitempty"`
}

// Invoice model
type Invoice struct {
	ID         string `json:"id"`
	AmountDue  int    `json:"amountDue"`
	AmountPaid int    `json:"amountPaid"`
	CreatedAt  string `json:"created_at"`
	DueBy      string `json:"dueBy"`
	Domain     string `json:"domain"`
	Owner      string `json:"owner"`
	Plan       Plan   `json:"plan"`
	Status     string `json:"status"`
	Space      string `json:"space"`
	ModifiedAt string `json:"modified_at"`
}

// BillPeriod get bill period
func (i *Invoice) BillPeriod() string {
	return "yearly"
}

// Title get title
func (i *Invoice) Title() string {
	return i.Plan.PlanName + " Invoice - " + i.ID
}

// ShortID get short id
func (i *Invoice) ShortID() string {
	return string(i.ID[0:5]) + " ... " + string(i.ID[len(i.ID)-5:])
}

// Description get description
func (i *Invoice) Description() string {
	return i.Domain + " - " + i.Plan.PlanName + " " + i.Plan.Name
}
