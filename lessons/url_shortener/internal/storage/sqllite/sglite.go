package sqllite

import "database/sql"

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqllite.New"
	
}
