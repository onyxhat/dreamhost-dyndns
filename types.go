package main

// dhRequest api as described here: http://bit.ly/2MHs8UU
type dhRequest struct {
	URL    string
	CMD    string
	Format string
	APIKey string
	Args   string
}

type dhDNSRecord struct {
	Record    string `json:"record"`
	Value     string `json:"value"`
	Type      string `json:"type"`
	Editable  string `json:"editable"`
	AccountID string `json:"account_id"`
	Comment   string `json:"comment"`
	Zone      string `json:"zone"`
}

type dhResponse struct {
	Result string        `json:"result"`
	Data   []dhDNSRecord `json:"data"`
}
