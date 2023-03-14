package store

type UserStore interface {
	User() UserRepository
}
