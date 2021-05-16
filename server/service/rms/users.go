package rms

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/rms"
	rp "gin-vue-admin/model/request/rms"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateUsers
//@description: 创建Users记录
//@param: Users model.Users
//@return: err error

func CreateUsers(Users mp.Users) (err error) {
	err = global.GVA_DB.Create(&Users).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUsers
//@description: 删除Users记录
//@param: Users model.Users
//@return: err error

func DeleteUsers(Users mp.Users) (err error) {
	err = global.GVA_DB.Delete(Users).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUsersByIds
//@description: 批量删除Users记录
//@param: ids request.IdsReq
//@return: err error

func DeleteUsersByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.Users{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateUsers
//@description: 更新Users记录
//@param: Users *model.Users
//@return: err error

func UpdateUsers(Users *mp.Users) (err error) {
	err = global.GVA_DB.Save(Users).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUsers
//@description: 根据id获取Users记录
//@param: id uint
//@return: err error, Users model.Users

func GetUsers(id uint) (err error, Users mp.Users) {
	err = global.GVA_DB.Where("id = ?", id).First(&Users).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUsersInfoList
//@description: 分页获取Users记录
//@param: info request.UsersSearch
//@return: err error, list interface{}, total int64

func GetUsersInfoList(info rp.UsersSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mp.Users{})
    var Users []mp.Users
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Users).Error
	err = db.Preload("Macs").Find(&Users).Error
	return err, Users, total
}

func GetUIList(userid string, password string, mac string) (err error, list interface{}, total int64) {
	var info rp.UsersSearch
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mp.Users{}).Where("name = ? and password = ?", userid,password)
    var Users []mp.Users
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Users).Error
	err = db.Preload("Macs", "mac_address = ?", mac).Find(&Users).Error
	return err, Users, total
}
