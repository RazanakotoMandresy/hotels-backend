package handler
// can add all info needed in the request
type fullRequest struct {
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Place            string   `json:"place"`
	Prix             uint     `json:"prix"`
	Status           bool     `json:"status"`
	ReservationLists []string `json:"reservation_lists"`
	Service          []string `json:"service"`
}

type filterRequest struct {
	Name      string   `json:"name"`
	Ouverture string   `json:"ouverture"`
	Place     string   `json:"place"`
	MinBudget uint     `json:"min_budget"`
	MaxBudget uint     `json:"max_budget"`
	Service   []string `json:"service"`
}

type userReq struct {
	Name      string `json:"name"`
	Passwords string `json:"passwords"`
	Mail      string `json:"mail"`
}
type reserveRequests struct {
	Starting_date string `json:"starting_date"`
	Ending_date   string `json:"ending_date"`
	Password      string `json:"passwords"`
}
