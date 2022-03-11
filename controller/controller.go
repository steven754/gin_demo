package controller

import (
	"gin_visit/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateVisitInfo(c *gin.Context) {
	// 1. 从请求中把数据拿出来
	var visitInfo models.VisitInfo
	c.BindJSON(&visitInfo)
	if visitInfo.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1001,
			"msg":  "name参数无效",
		})
	} else {
	// 2. 存入数据库
		err:=models.CreateVisitInfo(&visitInfo)
		if err != nil{
			c.JSON(http.StatusOK, gin.H{
				"code":1000,
				"msg":"fail",
				"error": err.Error()})
		}else{
			c.JSON(http.StatusOK, gin.H{
				"code": 2000,
				"msg":  "success",
				"data": visitInfo,
			})
		}
	}}

func GetVisitInfoList(c *gin.Context) {
	// 查询VisitInfo这个表里的所有数据
	visitInfoList, err := models.GetVisitInfo()
	if err!= nil {
		c.JSON(http.StatusOK, gin.H{
			"code":1000,
			"msg":"fail",
			"error": err.Error()})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg": "success",
			"data": visitInfoList,
		})
	}
}

func GetAVisitInfo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id"})
		return
	}
	visitInfo, err := models.GetAVisitInfo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1000,
			"msg": "fail",
			"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": visitInfo})
	}
}

func UpdateAVisitInfo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id"})
		return
	}
	visitInfo, err := models.GetAVisitInfo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&visitInfo)
	if err = models.UpdateAVisitInfo(visitInfo); err!= nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, visitInfo)
	}
}

func DeleteAVisitInfo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := models.DeleteAVisitInfo(id);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1000,
			"msg": "fail",
			"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg": "success",
			"data": id + "  deleted"})
	}
}
