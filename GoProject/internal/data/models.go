package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Helmets interface {
		Insert(helmet *Helmet) error
		Get(id int64) (*Helmet, error)
		Update(movie *Helmet) error
		Delete(id int64) error
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Helmets: HelmetModel{DB: db},
	}
}

func NewMockModels() Models {
	return Models{
		Helmets: MockHelmetModel{},
	}
}
