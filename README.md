### OrderSys
该项目是我上Oracle数据库这门课的课程项目，题目是老师说的关于订单管理的系统，没有什么难度，就是CRUD和一些存储过程的使用。
为了让自己找到一个理由认真做下去，我采用了前后端分离的设计模式。前端用vue+axios,后端用go+gin来写(主要是为了学习go,和vue)

### 项目结构
```
-OrderSys
|--src
    |--apis
        |--admin.go       # 管理员的操作
        |--custom.go      # 客户的操作
        |--saler.go       # 销售员的操作
        |--auth.go        # 权限认证
        |--apis.go        # API路由
    |--dataSource      
        |--source.go      # 配置数据源,这里我用的是oracle数据库
    |--models
        |--models.go      # 实体
        |--dao.go         # 数据库接口一些公共操作
        |--orderDao.go    # 关于订单的操作
        |--prosDao.go     # 关于商品的操作
        |--saleDao.go     # 关于销售员的一些操作
        |--userDao.go     # 关于用户的一些操作
    |--sql
        |--OrderSys.sql   # 数据表
    |--util
        |--security.go    # 加密工具
    |--app.go
|--README.md
```


### 第三方依赖

```
go get github.com/elgs/gosqljson
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors

### 国内可能无法直接下载

```

### 运行
* go run app.go

### 截图
![image](https://github.com/yahaa/OrderSys/raw/master/run.png)

### 前端情况下面链接
[OrderSys-ui](https://github.com/yahaa/OrderSys-ui)

