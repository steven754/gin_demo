package mysql

import (
	"fmt"
	"gin_visit/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)


var (
	DB  *gorm.DB
)

func ConnectMysql(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("mysql连接失败，请检查")
		log.Fatal(err)
		return
	}
	fmt.Println("mysql连接成功")
	return
}

