package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"strings"
	"sync/atomic"
	"time"
)

type PostgreSQLCluster struct {
	master                  *sql.DB
	slaves                  []*sql.DB
	slaveIndex              uint64
	dbMaxOpenConnections    int
	dbMaxIdleConnections    int
	dbConnectionMaxLifetime int
}

func (cluster *PostgreSQLCluster) Master() *sql.DB {
	return cluster.master
}

func (cluster *PostgreSQLCluster) Slave() *sql.DB {
	index := atomic.AddUint64(&cluster.slaveIndex, 1)
	return cluster.slaves[index%uint64(len(cluster.slaves))]
}

func NewPostgreSQLCluster(postgresqlMaster string, postgresqlSlaves string, dbMaxOpenConnections int, dbMaxIdleConnections int, dbConnectionMaxLifetime int) *PostgreSQLCluster {
	cluster := &PostgreSQLCluster{
		dbMaxOpenConnections:    dbMaxOpenConnections,
		dbMaxIdleConnections:    dbMaxIdleConnections,
		dbConnectionMaxLifetime: dbConnectionMaxLifetime,
	}

	slaveDSNList := strings.Split(postgresqlSlaves, ",")
	slaves := make([]*sql.DB, len(slaveDSNList))
	for i, dsn := range slaveDSNList {
		slaves[i] = cluster.newPostgreSQLDatabase(dsn)
	}

	cluster.master = cluster.newPostgreSQLDatabase(postgresqlMaster)
	cluster.slaves = slaves

	return cluster
}

func (cluster *PostgreSQLCluster) newPostgreSQLDatabase(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(cluster.dbMaxOpenConnections)
	db.SetMaxIdleConns(cluster.dbMaxIdleConnections)
	db.SetConnMaxLifetime(time.Minute * time.Duration(cluster.dbConnectionMaxLifetime))

	return db
}
