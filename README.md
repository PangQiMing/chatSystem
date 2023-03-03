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



### json
```text
User
{
	"id": 1,
	"userName": "jack",
	"password": "123456",
	"email": "953559192@qq.com"
}

Friends
{
	"id": 1,
    "friend_email": "111222333@qq.com",
	"user_id": 953559192,
    "user":{
        "id": 953559192,
	    "userName": "jack",
	    "email": "123456789@qq.com"
    }
}

Group
{
	"id": 1,
	"name": "群名",
	"notice": "群公告",
	"group_leaderID": "群主ID",
	"user_id": 953559192,
    "user":{
        "id": 953559192,
	    "userName": "jack",
	    "email": "123456789@qq.com"
    }
}

动态 
{
    "id": 1,
    "description": "今天也是焦头烂额的一天",
    "user_id"：953559192,
    "user": {
    	"id": 953559192,
		"user_name": "jack"
	}
}


OneToOneMessage
{
	"id": 1,
	"from_user_id": 953559192,
	"to_user_id": 111222333,
	"content": "hello golang"
}

GroupMessage
{
    "id": 1,
    "group_id":112233,
    "user_id": 953559192,
    "content": "群聊广播消息，欢迎进群"
}
```

