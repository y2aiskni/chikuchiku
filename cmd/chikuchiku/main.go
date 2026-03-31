package main

import (
	"fmt"

	"github.com/y2aiskni/chikuchiku/internal/usecase"
	"github.com/y2aiskni/chikuchiku/internal/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	conf, err := util.ReadYamlFile[config]("./config.yaml")
	if err != nil {
		panic(err)
	}

	db, err := openDB(conf.Mysql.Address, conf.Mysql.Port, conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Database)
	if err != nil {
		panic(err)
	}

	chikuchiku := usecase.NewChikuchiku(db)
	if err := chikuchiku.PostTodayToDiscord(conf.Discord.PostURL); err != nil {
		panic(err)
	}
}

func openDB(address string, port uint, username string, password string, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
