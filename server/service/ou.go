package service

import (
	"gin-manager/global"
	"gin-manager/model/request"
	"gin-manager/model/result"

	"gin-manager/model"
)

// @title    GetOuList
// @description   get ou list by pagination, 分页获取数据
// @auth                      （2020/04/05  20:22）
// @param     info             request.PageInfo
// @return    err              error
// @return    list             []result.Ou
// @return    total            int
func GetOuList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.DB.Model(&model.User{}).
		Table("ous as a").
		Joins("join ous as b ON a.pid = b.id")

	var userList []result.Ou

	err = db.Count(&total).Error

	err = db.Select("a.*", "b.name as parent_name").
		Limit(limit).Offset(offset).Find(&userList).Error

	return err, userList, total
}
