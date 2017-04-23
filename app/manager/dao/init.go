package dao

import (
	"github.com/jackc/pgx"
	"gopkg.in/inconshreveable/log15.v2"
	"os"
)

var connPool *pgx.ConnPool

func InitDB() {
	var err error
	connPool, err = pgx.NewConnPool(pgx.ConnPoolConfig{
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
	connPool.Prepare("SaveServer", `
	INSERT INTO "server" (ip, alive, port, status, updateat) VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (ip)
  	DO UPDATE
    	SET alive = $2, port = $3, status = $4, updateat = $5
    	WHERE server.ip = $1
	`)
	connPool.Prepare("FindAllServers", `
	SELECT ip,port,alive,status,updateat FROM server
	`)
	connPool.Prepare("GetNextServer", `
	SELECT ip,port,alive,status,updateat FROM server WHERE alive = TRUE LIMIT 1
	`)

	//Service
	connPool.Prepare("FindAllServicesByIp", `
	SELECT addr,password FROM service WHERE ip=$1
	`)
	connPool.Prepare("CreateService", `
	INSERT INTO service (addr, password,ip) VALUES ($1,$2,$3)
	`)

	//User
	connPool.Prepare("UpdateUser", `
	UPDATE "user" SET password=$1,token=$2,balance=$3,cost=$4,config=$5 WHERE email=$6;
	`)
	connPool.Prepare("CreateUser", `
	INSERT INTO "user" (email,password,token) VALUES ($1,$2,$3)
	`)
	connPool.Prepare("GetUserByEmail", `
	SELECT password, token,balance, cost, config
	FROM "user" WHERE email = $1;
	`)
	connPool.Prepare("GetUserByToken", `
	SELECT email,password,balance, cost, config
	FROM "user" WHERE token = $1;
	`)

	//UserService
	connPool.Prepare("UserService.Create", `
	INSERT INTO user_service(email, addr,uuid,device) VALUES ($1,$2,$3,$4)
	`)

}
