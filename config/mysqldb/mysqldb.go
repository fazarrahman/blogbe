package mysqldb

import (
	"log"

	"github.com/fazarrahman/blogbe/lib"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New() (*sqlx.DB, error) {
	username := lib.GetEnv("DB_USERNAME")
	password := lib.GetEnv("DB_PASSWORD")
	host := lib.GetEnv("DB_HOST")
	port := lib.GetEnv("DB_PORT")
	dbname := lib.GetEnv("DB_NAME")

	db, err := sqlx.Connect("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+dbname)
	if err != nil {
		log.Fatalln(err)
	}

	return db, nil
}
