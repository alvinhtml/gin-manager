package service

import (
	"errors"

	"github.com/alvinhtml/gin-manager/server/global"
	"github.com/alvinhtml/gin-manager/server/model/request"

	"github.com/alvinhtml/gin-manager/server/model"
	"github.com/alvinhtml/gin-manager/server/utils"

	"gorm.io/gorm"
)

// @title    CreateUser
// @description   create user, 用户创建
// @return    err             error
// @return    user       *User
func CreateUser(u model.User) (err error, user model.User) {
	if !errors.Is(global.DB.Where("name = ?", u.Name).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("User name registered"), user
	}
	// 使用 md5 加密
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.DB.Create(&u).Error
	return err, u
}

// @title    GetUsers
// @description   get user list by pagination, 分页获取数据
// @return    err              error
// @return    list             []result.User
// @return    total            int
func GetUsers(pageQuery request.PageQuery) (err error, list []model.UserJoinOu, total int64) {
	db := global.DB.Model(&model.User{}).
		Table("users as a").
		Joins("join ous as b ON a.ou_refer = b.id")

	// search
	if len(pageQuery.Search) > 0 {
		for k, v := range pageQuery.Search {
			db = db.Where("a."+k+" LIKE ?", "%"+v+"%")
		}
	}

	// sort
	if len(pageQuery.Sort) > 0 {
		length := len(pageQuery.Sort)
		sort := ""
		for k, v := range pageQuery.Sort {
			length--
			if length > 0 {
				sort += "a." + k + " " + v + ","
			} else {
				sort += "a." + k + " " + v
			}
		}
		db.Order(sort)
	}

	// filter
	if len(pageQuery.Filter) > 0 {
		for k, v := range pageQuery.Filter {
			db = db.Where("a."+k+" = ?", v)
		}
	}

	// count
	err = db.Count(&total).Error

	// limit
	limit := pageQuery.Size
	offset := pageQuery.Size * (pageQuery.Page - 1)

	err = db.Select("a.*", "b.name as ou_name").
		Limit(limit).Offset(offset).Find(&list).Error

	return err, list, total
}

// @title    		GetUser
// @description	get user by id 分页获取数据
// @return    err             error
// @return    user       *User
func GetUser(id uint) (err error, user model.UserJoinOu) {
	err = global.DB.
		Table("users as a").
		Joins("join ous as b ON a.ou_refer = b.id").
		Where("a.id = ?", id).
		Select("a.*", "b.name as ou_name").
		First(&user).Error
	return err, user
}

// @title    UpdateUser
// @description   update user info, 用户更新自己信息
// @return    err             error
// @return    user       *User
func UpdateUser(u model.User) (err error, user model.User) {
	err = global.DB.Model(&u).Updates(&u).First(&user).Error
	return err, user
}

// @title    UpdateUserPassword
// @description   update user password, 用户更新自己密码
// @return    err             error
// @return    user       *User
func UpdateUserPassword(u model.User) (err error, user model.User) {
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.DB.Model(&u).Updates(&u).First(&user).Error
	return err, user
}

// @title    DeleteUser
// @description   delete user by id, 用户删除自己
// @return    err             error
func DeleteUser(id uint) (err error) {
	// err = global.DB.Where("id = ?", id).Delete(&model.User{}).Error
	err = global.DB.Delete(&model.User{}, id).Error
	return err
}

// @title    Login
// @description   login, 用户登录
// @return    err             error
// @return    user       *User
func Login(u model.User) (err error) {
	u.Password = utils.MD5V([]byte(u.Password))

	db := global.DB.Where("name = ? AND password = ?", u.Name, u.Password).First(&u)

	if !errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return errors.New("User name or password error")
	}

	return nil
}

// @title    GetUserJoinOu
// @description   get user join ou, 获取用户所属组织
// @return    err             error
// @return    user       *User
func GetUserJoinOu(id uint) (err error, user model.UserJoinOu) {
	err = global.DB.Model(&user).
		Joins("join ous as b ON a.ou_refer = b.id").
		Where("a.id = ?", id).
		First(&user).Error

	return err, user
}
