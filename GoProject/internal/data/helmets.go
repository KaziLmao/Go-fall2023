package data

import (
	"GoProject/internal/validator"
	"database/sql"
	"encoding/json"
	"errors"
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
	}{
		ID:            h.ID,
		Name:          h.Name,
		Year:          year,
		Material:      h.Material,
		Ventilation:   h.Ventilation,
		Protection:    h.Protection,
		Weight:        h.Weight,
		SunProtection: h.SunProtection,
	}
	return json.Marshal(aux)
}

type HelmetModel struct {
	DB *sql.DB
}

func (h HelmetModel) Insert(helmet *Helmet) error {

	query := `
		INSERT INTO mhelmets (name, year, material, ventilation, protection, weight, sun_protection)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at`

	args := []interface{}{
		helmet.Name,
		helmet.Year,
		helmet.Material,
		helmet.Ventilation,
		helmet.Protection,
		helmet.Weight,
		helmet.SunProtection,
	}

	return h.DB.QueryRow(query, args...).Scan(&helmet.ID, &helmet.CreatedAt)
}

func (h HelmetModel) Get(id int64) (*Helmet, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
		SELECT id, created_at, name, year, material, ventilation, protection, weight, sun_protection
		FROM mhelmets
		WHERE id = $1`

	var helmet Helmet

	err := h.DB.QueryRow(query, id).Scan(
		&helmet.ID,
		&helmet.CreatedAt,
		&helmet.Name,
		&helmet.Year,
		&helmet.Material,
		&helmet.Ventilation,
		&helmet.Protection,
		&helmet.Weight,
		&helmet.SunProtection,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &helmet, nil
}

func (h HelmetModel) Update(helmet *Helmet) error {
	query := `
		UPDATE mhelmets
		SET name = $1, year = $2, material = $3, weight = $4, ventilation = $5
		WHERE id = $6`

	args := []interface{}{
		helmet.Name,
		helmet.Year,
		helmet.Material,
		helmet.Weight,
		helmet.Ventilation,
		helmet.ID,
	}

	return h.DB.QueryRow(query, args...).Scan(&helmet.ID)
}

func (h HelmetModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM mhelmets
		WHERE id = $1`

	result, err := h.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

type MockHelmetModel struct{}

func (h MockHelmetModel) Insert(helmet *Helmet) error {
	return nil
}
func (h MockHelmetModel) Get(id int64) (*Helmet, error) {
	return nil, nil
}
func (h MockHelmetModel) Update(helmet *Helmet) error {
	return nil
}
func (h MockHelmetModel) Delete(id int64) error {
	return nil
}
