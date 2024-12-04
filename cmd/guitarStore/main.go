package main

import (
	_ "github.com/lib/pq"
	middleware "github.com/oapi-codegen/nethttp-middleware"
	"guitarStore/api"
	"guitarStore/internal/configuration"
	"log"
	"net"
	"net/http"
	"strconv"
)

func main() {
	config := configuration.New()

	postgresqlMaster := config.Get("POSTGRESQL_MASTER")
	postgresqlSlaves := config.Get("POSTGRESQL_SLAVES")
	dbMaxOpenConnections, err := strconv.Atoi(config.Get("DB_MAX_OPEN_CONNECTIONS"))
	if err != nil {
		panic(err)
	}

	dbMaxIdleConnections, err := strconv.Atoi(config.Get("DB_MAX_IDLE_CONNECTIONS"))
	if err != nil {
		panic(err)
	}

	dbConnectionMaxLifetime, err := strconv.Atoi(config.Get("DB_CONNECTION_MAX_LIFETIME"))
	if err != nil {
		panic(err)
	}

	redisAddrs := config.Get("REDIS_ADDRS")

	swagger, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}
	swagger.Servers = nil

	server := NewServer(postgresqlMaster, postgresqlSlaves, dbMaxOpenConnections, dbMaxIdleConnections, dbConnectionMaxLifetime, redisAddrs)

	r := http.NewServeMux()

	api.HandlerFromMux(server, r)

	h := middleware.OapiRequestValidator(swagger)(r)

	s := &http.Server{
		Handler: h,
		Addr:    net.JoinHostPort("0.0.0.0", config.Get("PORT")),
	}

	log.Fatal(s.ListenAndServe())
}
