package store

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEmptyValue     = errors.New("Empty colums in select")
)

// /user=postgres password=p02tgre2 dbname=restapi_test sslmode=disable
//user=postgres password=admin dbname=users_dev sslmode=disable
