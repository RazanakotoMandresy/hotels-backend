package handlers

// import "github.com/RazanakotoMandresy/hotels-backend/internal/model"

// can add all info needed in the request
type fullRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Service     []string `json:"service"`
	Prix        uint     `json:"prix"`
	Status      int      `json:"status"`
	Ouverture   string   `json:"ouverture"`
}
