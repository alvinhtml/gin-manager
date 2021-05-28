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
// @auth                     （2020/04/05  20:22）
// @param     u               model.User
// @return    err             error
// @return    userInter       *User

func Register(u model.User) (err error, userInter model.User) {
	var user model.User
	if !errors.Is(global.GVA_DB.Where("name = ?", u.Name).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("User name registered"), userInter
	}
	// 使用 md5 加密
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.DB.Create(&u).Error
	return err, u
}

// @title    Login
// @description   login, 用户登录
// @auth                     （2020/04/05  20:22）
// @param     u               *model.User
// @return    err             error
// @return    userInter       *User

func Login(u *model.User) (err error, userInter *model.User) {
	var user model.User
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.DB.Where("name = ? AND password = ?", u.Name, u.Password).Preload("Authority").First(&user).Error
	return err, &user
}

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
