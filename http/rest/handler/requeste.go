package handler

// can add all info needed in the request
type fullRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Service     []string `json:"service"`
	Prix        uint     `json:"prix"`
	Status      bool     `json:"status"`
	Ouverture   string   `json:"ouverture"`
	Place       string   `json:"place"`
}
// notes ny ouverture date TODO de raha asiana front-end de de mampiasa input type date
type filterRequest struct {
	Name      string `json:"name"`
	Ouverture string `json:"ouverture"`
	Place     string `json:"place"`
	Service   string `json:"service"`
	Prix      uint   `json:"prix"`
}
type userReq struct {
	Name      string `json:"name"`
	Passwords string `json:"passwords"`
	Mail      string `json:"mail"`
}