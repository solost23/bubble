package mysql

import (
	"bubble/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitDB() (err error) {
	//dsn := "root:123@(localhost:3306)/bubble?charset=utf8mb4"
	mysqlConfig := config.NewConnectins()
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s",
		mysqlConfig.UserName,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.DB,
		mysqlConfig.Charset,
	)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}
