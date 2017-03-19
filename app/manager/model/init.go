package model

import (
	"github.com/jackc/pgx"
	"gopkg.in/inconshreveable/log15.v2"
	"os"
)

var DB *pgx.ConnPool

func InitDB() {
	var err error
	DB, err = pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     os.Getenv("PG_LS_HOST"),
			User:     os.Getenv("PG_LS_USER"),
			Password: os.Getenv("PG_LS_PASSWORD"),
			Database: "lightsocks",
			LogLevel: pgx.LogLevelInfo,
			Logger:   log15.New("module", "pgx"),
		},
		MaxConnections: 10,
	})
	if err != nil {
		panic(err)
	}

	//Server
	DB.Prepare("Server.Save", `
	INSERT INTO server(ip, port) VALUES ($1,$2) RETURNING updateat
	`)
	DB.Prepare("Server.UpdateStatus", `
	UPDATE server SET status=$1 WHERE ip=$2
	`)
	DB.Prepare("Server.UpdateAlive", `
	UPDATE server SET alive=$1 WHERE ip=$2
	`)
	DB.Prepare("Server.GetAllServers", `
	SELECT ip,port,alive,status,updateat FROM server
	`)

	//Service
	DB.Prepare("Service.AllServicesByIp", `
	SELECT port,password FROM service WHERE ip=$1
	`)

}
