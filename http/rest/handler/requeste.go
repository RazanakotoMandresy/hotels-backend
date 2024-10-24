package handler

// can add all info needed in the request
type fullRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	AvailableOn string   `json:"available_on"`
	Place       string   `json:"place"`
	Service     []string `json:"service"`
	Prix        uint     `json:"prix"`
	Status      bool     `json:"status"`
}

type filterRequest struct {
	Name      string   `json:"name"`
	Ouverture string   `json:"ouverture"`
	Place     string   `json:"place"`
	Service   []string `json:"service"`
	MinBudget uint     `json:"min_budget"`
	MaxBudget uint     `json:"max_budget"`
}

type userReq struct {
	Name      string `json:"name"`
	Passwords string `json:"passwords"`
	Mail      string `json:"mail"`
}
type reserveRequests struct {
	Date     string `json:"date"`
	Password string `json:"password"`
	// using the go gin banks api for
}
