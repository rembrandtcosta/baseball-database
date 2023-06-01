package database

import (
	"fmt"
	"io/ioutil"
	"database/sql"
)

func Connect() (*sql.DB, error) {
	bin, err := ioutil.ReadFile("/run/secrets/db-password")
	if err != nil {
		return nil, err
	}
	return sql.Open("postgres", fmt.Sprintf("postgres://postgres:%s@db:5432/example?sslmode=disable", string(bin)))
}
