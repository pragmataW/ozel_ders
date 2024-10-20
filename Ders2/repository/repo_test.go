package repository

import (
	"example/dto"
	"example/model"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dsn = "host=localhost user=postgres password=123wsedrf dbname=authentication port=5430 sslmode=disable"
)

func setup(model any) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")

	migrator := db.Migrator()
	migrator.CreateTable(model)

	return db
}

func TestCreateUser(t *testing.T) {
	db := setup(&model.User{})

	repo := Repo{
		Db: db,
	}

	user := dto.User{
		UserName: "pragmata",
		Password: "123wsedrf",
	}

	err := repo.AddUser(user)
	assert.NoError(t, err)

	var dbUser model.User
	result := repo.Db.First(&dbUser, "user_name = ?", user.UserName)

	assert.NoError(t, result.Error)
	assert.Equal(t, user.UserName, dbUser.Username)
	assert.Equal(t, user.Password, dbUser.Password)
}

func TestGetCredentials(t *testing.T) {
	db := setup(&model.User{})
	repo := Repo{
		Db: db,
	}

	user := model.User{
		Username: "pragmata",
		Password: "123wsedrf",
	}

	result := repo.Db.Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
	}
	assert.NoError(t, result.Error)

	dbUser, err := repo.GetCredentials(user.Username)
	if err != nil{
		log.Println(err)
	}
	assert.NoError(t, err)
	assert.Equal(t, user.Username, dbUser.Username)
	assert.Equal(t, user.Password, dbUser.Password)
}