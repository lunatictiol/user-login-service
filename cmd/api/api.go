package api

import "database/sql"

type ApiServer struct {
	address string
	db      *sql.DB
}

func (a *ApiServer) New(address string,
	db *sql.DB) {
	a.address = address
	a.db = db

}
