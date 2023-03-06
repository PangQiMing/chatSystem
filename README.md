## 使用到的包
```text
go get github.com/gin-gonic/gin
go get gorm.io/driver/mysql
go get gorm.io/gorm
go get github.com/joho/godotenv
go get github.com/dgrijalva/jwt-go
go get golang.org/x/crypto/bcrypt
go get github.com/mashingan/smapping
go get github.com/gorilla/websocket
```

## 需求分析
### 用户注册

1. 设置名称、email、密码
2. 利用JWT（json web token）生成token,然后发放token

### 用户登陆

1. 获取email和密码
2. 验证email和密码
3. 验证成功，生成最新的token,发放最新的token

### 好友列表

登陆成功后，验证token，获取token里面加密的id，才能进行一下操作

1. 添加好友（添加好友时，要判断好友是否已存在，或者没有此好友）
2. 删除好友
3. 获取好友列表

### 群组

1. 创建群组
2. 更新群组信息
3. 获取群组成员列表

### 朋友圈

1. 发布动态

2. 删除动态

3. 获取动态

   

### 好友一对一通信




### 群组广播通信

### 朋友圈
~~~json
{
	"user_id": 0,
	"avatar": "头像URL",
	"nick_name": "JACK",
	"email": "123456789@qq.com",
	"password": "12345678",
	"age": "20",
	"sex": "男",
	"user_status": 0,
	"circle_id": 122,
	"circle": {
		"id": 122,
		"content": "内容：HelloWorld",
		"picture_url": "图片url_1"
	}
}
~~~




