package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"zwell.github/mic-server/sms/internal/config"
)

var DB *sqlx.DB

type Connector struct {
}

func GetDB () *sqlx.DB {

	if DB == nil {
		DB = connection()
	}

	return DB
}

func connection () *sqlx.DB {

	conf := config.GetConf()

	DB, err := sqlx.Open("mysql", "" + conf.Mysql.User + ":" + conf.Mysql.Password + "@tcp(" + conf.Mysql.Host + ":" + conf.Mysql.Port + ")/" + conf.Mysql.Db)
	if err != nil {
		log.Fatalln(err)
	}

	return DB
}