package models

type Plan struct {
	Price            int    `json:"price,omitempty"`
	UsedDisk         int    `json:"usedDisk,omitempty"`
	BandwidthAllowed int    `json:"bandwidthAllowed"`
	CreatedAt        string `json:"created_at"`
	ModifiedAt       string `json:"modified_at"`
	DiskAllowed      int    `json:"diskAllowed"`
	EmailsAllowed    int    `json:"emailsAllowed"`
	ID               string `json:"id"`
	Name             string `json:"name"`
	PlanName         string `json:"planName"`
}

type PlanList struct {
	Count      int    `json:"count"`
	Data       []Plan `json:"data"`
	TotalCount int    `json:"total_count"`
}
