package sqlstore

import (
	"database/sql"
	"log"
	"restApi/internal/app/store"

	_ "github.com/lib/pq"
)

//Структура хранилища

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

//инициализацияя хранилища

func New(db *sql.DB) *Store {
	log.Println("Store initialization")
	return &Store{
		db: db,
	}
}

//Оболочка пользователя над хранилищем

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}

//store.User().Create()
