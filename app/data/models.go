package data

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type Database struct {
	*gorm.DB
}

func NewDB(dataSource string) (*Database, error) {
	database, err := gorm.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}
	if err := database.DB().Ping(); err != nil {
		return nil, err
	}

	// go-sql-driver/mysql - Best practice
	// https://github.com/go-sql-driver/mysql/issues/461#issuecomment-227008369
	database.DB().SetConnMaxLifetime(time.Minute * 5)
	database.DB().SetMaxIdleConns(0)
	database.DB().SetMaxOpenConns(5)

	return &Database{
		DB: database,
	}, nil

}

func (database *Database) Migrate() {
	database.AutoMigrate(&User{})
}

type User struct {
	gorm.Model
	UserName string `gorm:"unique_index"`
	Password []byte
}

type Book struct {
	gorm.Model
	Title  string
	Author string
}
