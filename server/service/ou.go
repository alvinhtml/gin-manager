package service

import (
	"github.com/alvinhtml/gin-manager/server/global"
	"github.com/alvinhtml/gin-manager/server/model"
	"github.com/alvinhtml/gin-manager/server/model/request"
)

// @title    CreateOu
// @description   create ou, 创建组织单位
// @return    err              error
func CreateOu(o model.Ou) (err error, ou model.Ou) {
	err = global.DB.Create(&o).First(&ou).Error
	return err, ou
}

// @title    GetOus
// @description   get ou list by pagination, 分页获取数据
// @return    err              error
// @return    list             []result.Ou
// @return    total            int
func GetOus(info request.PageInfo) (err error, ous []model.Ou, total int64) {
	limit := info.Size
	offset := info.Size * (info.Page - 1)

	err = global.DB.Model(&model.Ou{}).Offset(offset).Limit(limit).Find(&ous).Count(&total).Error

	return err, ous, total
}

// @title    GetOu
// @description   get ou by id, 根据id获取组织单位
// @return    err              error
// @return    ou               result.Ou
func GetOu(id uint) (err error, ou model.Ou) {
	err = global.DB.First(&ou, id).Error
	return err, ou
}

// @title    UpdateOu
// @description   update ou, 更新组织单位
// @return    err              error
func UpdateOu(o model.Ou) (err error, ou model.Ou) {
	err = global.DB.Save(&o).First(&ou).Error
	return err, ou
}

// @title    DeleteOu
// @description   delete ou, 删除组织单位
// @return    err              error
func DeleteOu(id uint) (err error) {
	err = global.DB.Delete(&model.Ou{}, id).Error
	return err
}
