package apiserver

import (
	"database/sql"
	"log"
	"net/http"
	"restApi/internal/app/logging"
	sqlstore "restApi/internal/app/store/sqlstore/users"
)

// запуск сервера
func Start(config *Config) error {

	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	store := sqlstore.New(db)
	logger := logging.GetLogger()
	server := newServer(store, config, logger)
	return http.ListenAndServe(config.BindAddr, server)
}

// инициализация бд
func newDB(databaseURL string) (*sql.DB, error) {
	log.Printf("Database initialization: %s\n", databaseURL)
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
