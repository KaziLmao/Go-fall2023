package data

import (
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
	Design        string    `json:"design"`         // The aesthetic design of the helmet.
	Weight        float64   `json:"weight"`         // Weight of the helmet in kilograms.
	SunProtection bool      `json:"sun_protection"` // Whether the helmet has an integrated sun protection visor.
	Lining        string    `json:"lining"`         // The material of the lining.
	Fastening     string    `json:"fastening"`      // The helmet's fastening system (e.g., "Quick-release buckle").
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
		Design        string  `json:"design"`
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
		Design:        h.Design,
		Weight:        h.Weight,
		SunProtection: h.SunProtection,
		Lining:        h.Lining,
		Fastening:     h.Fastening,
	}
	return json.Marshal(aux)
}
