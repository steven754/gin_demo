package models

import (
	"gin_visit/mysql"
)

// VisitInfo Model
type VisitInfo struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Sex string `json:"sex"`
	Age int 	`json:"age	"`
	Old bool `json:"old"`
	Card string `json:"card"`
	Iphone string `json:"iphone"`
}
//
//func (v VisitInfo) IsEmpty() (err error){
//	if v.Name == "" || v.Name == null {
//		err = "name不能为空"
//		return
//	}
//}

/*
	visitInfo这个Model的增删改查操作都放在这里
*/
// CreateVisitInfo 创建VisitInfo
func CreateVisitInfo(visitInfo *VisitInfo) (err error){
	err = mysql.DB.Create(&visitInfo).Error
	return
}

func GetVisitInfo() (visitInfoList []*VisitInfo, err error){
	if err = mysql.DB.Find(&visitInfoList).Error; err != nil{
		return nil, err
	}
	return
}

func GetAVisitInfo(id string)(visitInfo *VisitInfo, err error){
	visitInfo = new(VisitInfo)
	if err = mysql.DB.Debug().Where("id=?", id).First(visitInfo).Error; err!=nil{
		return nil, err
	}
	return
}

func UpdateAVisitInfo(visitInfo *VisitInfo)(err error){
	err = mysql.DB.Save(visitInfo).Error
	return
}

func DeleteAVisitInfo(id string)(err error){
	err = mysql.DB.Where("id=?", id).Delete(&VisitInfo{}).Error
	return
}

