package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const DRIVER = "mysql"
var SqlSession *gorm.DB

type conf struct {
	Url string `yaml:"url"`
	UserName string `yaml:"userName"`
	Passwrod string `yaml:"password"`
	DbName string `yaml:"dbName"`
	Port string `yaml:"port"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("./resources/application.yaml")
	if err !=nil{
		fmt.Printf("读取yaml文件 %s 出错\n", yamlFile)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil{
		fmt.Printf("加载yaml %s 出错\n", yamlFile)
	}
	return c
}

func InitMySql()(err error){
	var c conf
	conf := c.getConf()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName, conf.Passwrod, conf.Url, conf.Port, conf.DbName)
	SqlSession, err = gorm.Open(DRIVER, dsn)
	if err != nil{
		fmt.Printf("连接数据库失败 host=%s port=%s user=%s pwd=%s db=%s\n", conf.Url, conf.Port, conf.UserName, conf.Passwrod, conf.DbName)
		panic(err)
	}
	return SqlSession.DB().Ping()
}

func Close(){
	SqlSession.Close()
}