package main

import (
	"fmt"
	"go-admin/go/dao"
	"go-admin/go/entity"
	"go-admin/go/routes"
)

func main(){
	err := dao.InitMySql()
	if err != nil{
		fmt.Println("初始化数据库失败")
		panic(err)
	}
	defer dao.Close()
	// 绑定模型
	dao.SqlSession.AutoMigrate(&entity.User{})
	r := routes.SetRouter()
	r.Run(":8081")

}
