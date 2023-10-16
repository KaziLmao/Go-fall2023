package data

import (
	"GoProject/internal/validator"
	"encoding/json"
	"fmt"
	"time"
)

type Helmet struct {
	ID            int64     `json:"id"`             // Unique integer ID for the helmet
	CreatedAt     time.Time `json:"-"`              // Timestamp for when the helmet is added to our database
	Name          string    `json:"name"`           // Helmet name
	Year          int32     `json:"year"`           // Helmet release year
	Material      string    `json:"material"`       // Material used in the construction of the helmet.
	Ventilation   bool      `json:"ventilation"`    // Ventilation system in the helmet.
	Protection    string    `json:"protection"`     // Safety certification of the helmet (e.g., "DOT", "ECE", "Snell").
	Weight        float64   `json:"weight"`         // Weight of the helmet in kilograms.
	SunProtection bool      `json:"sun_protection"` // Whether the helmet has an integrated sun protection visor.
	Lining        string    `json:"lining"`         // The material of the lining.
	Fastening     string    `json:"fastening"`      // The helmet's fastening system (e.g., "Quick-release buckle").
}

func ValidateHelmet(v *validator.Validator, helmet *Helmet) {
	v.Check(helmet.Name != "", "title", "must be provided")
	v.Check(len(helmet.Name) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(helmet.Year != 0, "year", "must be provided")
	v.Check(helmet.Year >= 1888, "year", "must be greater than 1888")
	v.Check(helmet.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	//	Needs to be added some checks
}

func (h Helmet) MarshalJSON() ([]byte, error) {
	var year string
	if h.Year != 0 {
		year = fmt.Sprintf("%d year", h.Year)
	}

	aux := struct {
		ID            int64   `json:"id"`
		Name          string  `json:"name"`
		Year          string  `json:"year"`
		Material      string  `json:"material"`
		Ventilation   bool    `json:"ventilation"`
		Protection    string  `json:"protection"`
		Weight        float64 `json:"weight"`
		SunProtection bool    `json:"sun_protection"`
		Lining        string  `json:"lining"`
		Fastening     string  `json:"fastening"`
	}{
		ID:            h.ID,
		Name:          h.Name,
		Year:          year,
		Material:      h.Material,
		Ventilation:   h.Ventilation,
		Protection:    h.Protection,
		Weight:        h.Weight,
		SunProtection: h.SunProtection,
		Lining:        h.Lining,
		Fastening:     h.Fastening,
	}
	return json.Marshal(aux)
}
