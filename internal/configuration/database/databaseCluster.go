package database

import "database/sql"

type DatabaseCluster interface {
	Master() *sql.DB
	Slave() *sql.DB
}
