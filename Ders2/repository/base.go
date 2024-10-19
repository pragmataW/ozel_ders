package repository

import (
	"example/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
						conf.Host, conf.User, conf.Password, conf.DbName, conf.Port, conf.SslMode, conf.TimeZone)
	
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		return Repo{}, err
	}

	sqlDb, err := db.DB()
	if err != nil{
		return Repo{}, err
	}

	if err = sqlDb.Ping(); err != nil{
		return Repo{}, err
	}

	db.AutoMigrate(&model.User{})

	return Repo{
		Db: db,
	}, nil
}
