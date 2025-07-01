package db

import (
	"database/sql"
	"fmt"
	"go-api/src/config"
	_ "github.com/lib/pq"
)

// LoadDataBase se conecta e carrega a base de dados
func LoadDataBase() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DB_host, config.DB_port, config.DB_user, config.DB_password, config.DB_name,
	)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	fmt.Println("Connected to " + config.DB_name)

	return db, nil

}
