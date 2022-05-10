package service

import (
	"go-admin/go/dao"
	"go-admin/go/entity"
)

// 创建User
func CreateUser(user *entity.User)(err error){
	if err = dao.SqlSession.Create(user).Error; err!=nil{
		return err
	}
	return
}

// 获取User列表
func GetUserList() (userList []*entity.User, err error){
	if err = dao.SqlSession.Find(&userList).Error; err!=nil{
		return nil, err
	}
	return
}

//删除用户
func DeleteUserById(id string) (err error){
	err = dao.SqlSession.Where("id = ?", id).Delete(&entity.User{}).Error
	return
}

// 查询用户
func GetUserById(id string)(user *entity.User, err error){
	user = &entity.User{}
	if err = dao.SqlSession.Where("id = ?", id).First(user).Error;err !=nil{
		return nil, err
	}
	return
}

// 更新用户信息
func UpdateUser(user *entity.User)(err error){
	if err = dao.SqlSession.Save(user).Error;err!=nil{
		return err
	}
	return
}

