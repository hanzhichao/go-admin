# go-admin
Gin + Gorm 框架接口示例

## 使用方式
下载源码
修改resource/application.yaml中数据库配置
```yaml
url: 127.0.0.1  # 数据库地址
userName: root  # 数据库用户名
password: passw0rd  # 数据库密码
dbName: go_admin  # 数据库名称，需要手动创建
port: 3306   # 数据库端口号
```
> 注意：dbName对应数据库需要手动创建，如`CREATE DATABASE go_admin`

## 接口列表

### 1.创建用户
- 接口路径：/api/users
- 请求方法：POST
- 内容类型：application/json

示例
```shell
curl --location --request POST 'http://127.0.0.1:8081/api/users' \
--header 'Content-Type: application/json' \
--data-raw '{
"name": "张三",
"nickName": "三儿",
"avatar": "https://",
"password": "123",
"email": "zhangshan@123.com",
"mobile": "12112345678",
"delstatus": 0,
"createTime": 1652168218
}'
```
### 2.查询用户列表
- 接口路径：/api/users
- 请求方法：GET
示例
```shell
curl --location --request GET 'http://127.0.0.1:8081/api/users'
```

### 3.查询用户信息
- 接口路径：/api/users/<id>
- 请求方法：GET

示例
```shell
curl --location --request GET 'http://127.0.0.1:8081/api/users/1'
```

### 4.更新用户信息
- 接口路径：/api/users/<id>
- 请求方法：PUT
- 内容类型：application/json

示例
```shell
curl --location --request PUT 'http://127.0.0.1:8081/api/users/11' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "张三儿",
    "nickName": "三儿",
    "avatar": "https://",
    "password": "123",
    "email": "zhangshan@123.com",
    "mobile": "12112345678",
    "delstatus": 0,
    "createTime": 1652168218
}'
```

### 5.删除用户
- 接口路径：/api/users/<id>
- 请求方法：DELETE

示例
```shell
curl --location --request DELETE 'http://127.0.0.1:8081/api/users/1'
```


## 参考
- 教程链接：https://www.cnblogs.com/zhujiqian/p/15074010.html
- Gin: https://github.com/gin-gonic/gin
- Gin Web Framework: https://chenyitian.gitbooks.io/gin-web-framework/content/
