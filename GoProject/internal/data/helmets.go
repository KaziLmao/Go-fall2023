package data

import (
	"GoProject/internal/validator"
	"context"
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return h.DB.QueryRowContext(ctx, query, args...).Scan(&helmet.ID, &helmet.CreatedAt)
}

func (m HelmetModel) GetAll(name string, material string, protection string, filters Filters) ([]*Helmet, Metadata, error) {
	query := fmt.Sprintf(`
		SELECT id, created_at, name, year, material, ventilation, protection, weight, sun_protection
		FROM mhelmets
		WHERE (STRPOS(LOWER(name), LOWER($1)) > 0 OR $1 = '')
		AND (STRPOS(LOWER(material), LOWER($2)) > 0 OR $2 = '')
		AND (STRPOS(LOWER(protection), LOWER($3)) > 0 OR $3 = '')
		ORDER BY %s %s, id ASC
		LIMIT $4 OFFSET $5`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, name, material, filters.limit(), filters.offset())
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	helmets := []*Helmet{}

	for rows.Next() {
		var helmet Helmet
		err := rows.Scan(
			&totalRecords,
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
			return nil, Metadata{}, err
		}

		helmets = append(helmets, &helmet)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return helmets, metadata, nil
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := h.DB.QueryRowContext(ctx, query, id).Scan(
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
		SET name = $1, year = $2, material = $3, ventilation = $4, protection = $5, weight = $6, sun_protection = $7
		WHERE id = $8
		RETURNING id`

	args := []interface{}{
		helmet.Name,
		helmet.Year,
		helmet.Material,
		helmet.Ventilation,
		helmet.Protection,
		helmet.Weight,
		helmet.SunProtection,
		helmet.ID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := h.DB.QueryRowContext(ctx, query, args...).Scan(&helmet.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (h HelmetModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM mhelmets
		WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := h.DB.ExecContext(ctx, query, id)
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

//type MockHelmetModel struct{}
//
//func (h MockHelmetModel) Insert(helmet *Helmet) error {
//	return nil
//}
//
//func (h MockHelmetModel) GetAll(id int64) (*Helmet, error) {
//	return nil, nil
//}
//
//func (h MockHelmetModel) Get(id int64) (*Helmet, error) {
//	return nil, nil
//}
//func (h MockHelmetModel) Update(helmet *Helmet) error {
//	return nil
//}
//func (h MockHelmetModel) Delete(id int64) error {
//	return nil
//}
