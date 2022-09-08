package department

import (
	"database/sql"
	"restApi/internal/app/store"

	_ "github.com/lib/pq"
)

type Store struct {
	db                   *sql.DB
	DepartmentRepository *DepartmentRepositor
}

func (s *Store) Department() store.DepartmentRepository {
	if s.DepartmentRepository != nil {
		return s.DepartmentRepository
	}

	s.DepartmentRepository = &DepartmentRepositor{
		store: s,
	}
	return s.DepartmentRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}
