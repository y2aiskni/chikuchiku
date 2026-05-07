package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/y2aiskni/chikuchiku/internal/usecase"
	"github.com/y2aiskni/chikuchiku/internal/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var cmd Cmd
	ctx := kong.Parse(&cmd)
	switch command := ctx.Command(); command {
	case "post discord":
		if err := runPostDiscord(cmd.Post.Config); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	default:
		panic("command not implemented: " + command)
	}
}

func runPostDiscord(configFilePath string) error {
	conf, err := util.ReadYamlFile[config](configFilePath)
	if err != nil {
		return fmt.Errorf("failed to util.ReadYamlFile(): %w", err)
	}

	db, err := openDB(conf.Mysql.Address, conf.Mysql.Port, conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Database)
	if err != nil {
		return fmt.Errorf("failed to openDB(): %w", err)
	}

	chikuchiku := usecase.NewChikuchiku(db)
	if err := chikuchiku.PostTodayToDiscord(conf.Discord.PostURL); err != nil {
		return fmt.Errorf("failed to PostTodayToDiscord(): %w", err)
	}

	return nil
}

func openDB(address string, port uint, username string, password string, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
