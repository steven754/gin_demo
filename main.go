package main

import (
	"fmt"
	"gin_visit/models"
	"gin_visit/mysql"
	"gin_visit/router"
	"gin_visit/setting"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage：./bubble conf/config.ini")
		return
	}
	// 加载配置文件
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := mysql.ConnectMysql(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("connect mysql failed, err:%v\n", err)
		return
	}
	//defer dao.Close() // 程序退出关闭数据库连接
	// 模型绑定
	mysql.DB.AutoMigrate(&models.VisitInfo{})
	// 注册路由
	r := router.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}

