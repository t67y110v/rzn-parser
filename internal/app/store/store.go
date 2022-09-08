package store

type UserStore interface {
	User() UserRepository
}

type DepartmentStore interface {
	Department() DepartmentRepository
}
