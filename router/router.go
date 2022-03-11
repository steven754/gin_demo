package router

import (
	"gin_visit/controller"
	"gin_visit/setting"
	"github.com/gin-gonic/gin"
)


type VisitInfo struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Sex string `json:"sex"`
	Age int 	`json:"age	"`
	Old bool `json:"old"`
	Card string `json:"card"`
	Iphone string `json:"iphone"`
}

func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	v1Group := r.Group("/v1")
	{ //获取访客信息
		v1Group.GET("/visitInfo", controller.GetVisitInfoList)

		//获取某一访客详细信息
		v1Group.GET("/visitInfo/:id", controller.GetAVisitInfo)

		//添加访客信息
		v1Group.POST("/visitInfo", controller.CreateVisitInfo)

		//修改访客信息
		v1Group.PUT("/visitInfo/:id", controller.UpdateAVisitInfo)

		//删除访客信息
		v1Group.DELETE("/visitInfo/:id", controller.DeleteAVisitInfo)
	}
	return r
}
