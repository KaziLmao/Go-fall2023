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
		GetAll(name string, material string, protection string, filters Filters) ([]*Helmet, Metadata, error)
	}
	Users UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Helmets: HelmetModel{DB: db},
		Users:   UserModel{DB: db},
	}
}
