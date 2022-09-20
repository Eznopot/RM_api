package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

var once sync.Once

var (
	instance *sql.DB
)

type DbConfig struct {
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
		Ip       string `yaml:"ip"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func GetDb() *sql.DB {

	once.Do(func() {
		f, err := os.Open("config.yml")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		defer f.Close()
		var cfg DbConfig
		decoder := yaml.NewDecoder(f)
		err = decoder.Decode(&cfg)
		if err != nil {
			os.Exit(-1)
		}
		instance, err = sql.Open("mysql", cfg.Database.Username+":"+cfg.Database.Password+"@tcp("+cfg.Database.Ip+":"+cfg.Database.Port+")/"+cfg.Database.Name)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		var version string
		instance.QueryRow("SELECT VERSION()").Scan(&version)
	})
	return instance
}
