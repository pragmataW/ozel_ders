package repository

import (
	"database/sql"
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

type Repo struct {
	Db *gorm.DB
}

type DbConf struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SslMode  string
	TimeZone string
}

func New(conf DbConf) (Repo, error) {
	var err error
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			conf.Host, conf.User, conf.Password, conf.DbName, conf.Port, conf.SslMode, conf.TimeZone)

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return
		}

		var sqlDb *sql.DB
		sqlDb, err = db.DB()
		if err != nil {
			return
		}

		err = sqlDb.Ping()
		if err != nil{
			return
		}
	})

	if err != nil {
		return Repo{}, err
	}

	return Repo{Db: db}, nil
}
