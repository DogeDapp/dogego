# DogeGo

使用DogeGo开发分布式WebAPI: 用最简单的架构, 实现支持分布式的够用的框架, 服务大量的用户.

本框架是我从我编写的很多上生产的项目中拆分出来的(包括个人外包和学校项目), 所以性能是有保证的, 而且得益于Go的超高性能, 而且我的设计, 本身就是支持分布式的, 所以, 可以放心使用.

### 目的
本项目采用了一些Golang里面比较流行的我试用过的组件，让大家快速搭建支持分布式的Restful Web API.

### 特色
本项目整合了开发API的基础类库:
1. [Gin](https://github.com/gin-gonic/gin): 轻量级Web框架，自称路由速度是golang最快的 
2. [GORM](http://gorm.io/docs/index.html): ORM工具。本项目需要配合Mysql使用 
3. [Gin-Session](https://github.com/gin-contrib/sessions): Gin框架提供的Session操作工具
4. [Go-Redis](https://github.com/go-redis/redis): Golang Redis客户端
5. [godotenv](https://github.com/joho/godotenv): 开发环境下的环境变量工具，方便使用环境变量
6. [Gin-Cors](https://github.com/gin-contrib/cors): Gin框架提供的跨域中间件
7. [DogeGo Task](https://github.com/xiangrui2019/dogego_task) 自行实现了分布式定时任务
8. 本项目使用Redis分布式Session存储登录信息, 默认支持分布式, 无需更换

本项目已经预先实现了一些常用的代码方便参考和复用:

1. 创建了用户模型
2. 实现了```/api/v1/ping```心跳检查接口
3. 实现了```/api/v1/user/register```用户注册接口
4. 实现了```/api/v1/user/login```用户登录接口
5. 实现了```/api/v1/user/me```用户信息(需要登录后获取session)
5. 实现了```/api/v1/user/change_password```用户密码修改接口(需要登录后获取session)
5. 实现了```/api/v1/user/update_profile```用户Profile更新接口(需要登录后获取session)
5. 实现了```/api/v1/user/logout```用户登出接口(需要登录后获取session)
6. 实现了```time```定时任务, 配合[DogeGo Task](https://github.com/xiangrui2019/dogego_task)使用
