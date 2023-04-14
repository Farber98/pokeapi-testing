package models

type ErrorResponse struct {
	Error string `json:"error"`
}

type Pokemon struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	Power     string         `json:"type"`
	Abilities map[string]int `json:"abilities"`
}

var Abilities = map[string]int{
	"hp":      0,
	"attack":  0,
	"defense": 0,
	"speed":   0,
}

var AllowedAbilities = map[string]string{
	"hp":      "hp",
	"attack":  "attack",
	"defense": "defense",
	"speed":   "speed",
}
