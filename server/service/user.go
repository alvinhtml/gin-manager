package service

import (
	"errors"
	"gin-manager/global"
	"gin-manager/model/request"
	"gin-manager/model/result"

	"gin-manager/model"

	"github.com/alvinhtml/gin-manager/server/model"
	"github.com/alvinhtml/gin-manager/server/utils"
)

// @title    Register
// @description   register, 用户注册
// @param     u               model.User
// @return    err             error
// @return    user       *User

func Register(u model.User) (err error, user model.User) {
	var user model.User
	if !errors.Is(global.GVA_DB.Where("name = ?", u.Name).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("User name registered"), user
	}
	// 使用 md5 加密
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.DB.Create(&u).Error
	return err, u
}

// @title    Login
// @description   login, 用户登录
// @param     u               *model.User
// @return    err             error
// @return    user       *User

func Login(u *model.User) (err error, user *model.User) {
	var user model.User
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.DB.Where("name = ? AND password = ?", u.Name, u.Password).Preload("Authority").First(&user).Error
	return err, &user
}

// @title    GetUsers
// @description   get user list by pagination, 分页获取数据
// @param     info             request.PageInfo
// @return    err              error
// @return    list             []result.User
// @return    total            int
func GetUsers(info request.PageInfo) (err error, list []result.User, total int64) {
	limit := info.Size
	offset := info.Size * (info.Page - 1)

	db := global.DB.Model(&model.User{}).
		Table("users as a").
		Joins("join ous as b ON a.ou_refer = b.id")

	err = db.Count(&total).Error

	err = db.Select("a.*", "b.name as ou_name").
		Limit(limit).Offset(offset).Find(&list).Error

	return err, list, total
}

// @title    		GetUser
// @description	get user by id 分页获取数据
// @param     id              int
// @return    err             error
// @return    user       *User
func GetUser(id int) (err error, user model.User) {
	err = global.DB.Where("id = ?", id).First(&user).Error
	return err, user
}

// @title    UpdateUser
// @description   update user info, 用户更新自己信息
// @param     u               model.User
// @return    err             error
// @return    user       *User
func UpdateUser(u model.User) (err error, user model.User) {
	err = global.DB.Model(&u).Update(u).Error
	return err, u
}

// @title    UpdateUserPassword
// @description   update user password, 用户更新自己密码
// @param     u               model.User
// @return    err             error
// @return    user       *User
func UpdateUserPassword(u model.User) (err error, user model.User) {
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.DB.Model(&u).Update(u).Error
	return err, u
}

// @title    DeleteUser
// @description   delete user by id, 用户删除自己
// @param     id              int
// @return    err             error
func DeleteUser(id int) (err error) {
	err = global.DB.Where("id = ?", id).Delete(&model.User{}).Error
	return err
}
