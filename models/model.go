package models

import (
	"database/sql"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func NewConfig() *Config {
	c := &Config{}
	c.Host, _ = beego.AppConfig.String("mysql_host")
	c.Port, _ = beego.AppConfig.Int("mysql_port")
	c.User, _ = beego.AppConfig.String("mysql_user")
	c.Password, _ = beego.AppConfig.String("mysql_password")
	c.Database, _ = beego.AppConfig.String("mysql_db")
	return c
}

var DB *sql.DB

func InitDB(config *Config) error {
	dsn := config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + strconv.Itoa(config.Port) + ")/" + config.Database + "?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	DB = db
	logs.Info("Connected to MySQL database!")
	return nil
}

// Close closes the database connection
func Close() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			return
		}
	}
}
