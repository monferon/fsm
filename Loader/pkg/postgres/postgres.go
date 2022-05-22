package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func New(url string) (*Postgres, error) {
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		//TODO error
		return nil, err
	}
	return &Postgres{
		DB: db,
	}, nil
}
