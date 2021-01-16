package service

import (
	"gin-manager/global"
	"gin-manager/model/request"
	"gin-manager/model/result"

	"gin-manager/model"
)

// @title    GetInfoList
// @description   get user list by pagination, 分页获取数据
// @auth                      （2020/04/05  20:22）
// @param     info             request.PageInfo
// @return    err              error
// @return    list             []result.User
// @return    total            int
func GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.DB.Model(&model.User{}).
		Table("users as a").
		Joins("join ous as b ON a.ou_refer = b.id")

	var userList []result.User

	err = db.Count(&total).Error

	err = db.Select("a.*", "b.name as ou_name").
		Limit(limit).Offset(offset).Find(&userList).Error

	return err, userList, total
}
