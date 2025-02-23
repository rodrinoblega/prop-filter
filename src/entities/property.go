package entities

type Property struct {
	SquareFootage int             `json:"squareFootage"`
	Lighting      string          `json:"lighting"`
	Price         float64         `json:"price"`
	Rooms         int             `json:"rooms"`
	Bathrooms     int             `json:"bathrooms"`
	Location      [2]float64      `json:"location"`
	Description   string          `json:"description"`
	Amenities     map[string]bool `json:"amenities"`
}
