package api

import (
	"encoding/base64"
	"game/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var conn *sql.DB

type userData struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func Login(c *gin.Context) {
	/*
		mock data
		{
			"Username":"test",
			"Password":"test"
		}
	*/
	var realpass string
	var logdata userData

	// 将Postdata绑定到logdata上
	if err := c.BindJSON(&logdata); err != nil {
		return
	}

	//连接数据库并查询指定用户名的密码，将密码保存在realpass变量中
	// !!! 现存缺陷：没法将数据库连接保持在整个程序中，每次运行login都要重新做一次数据库连接
	conn := config.Dbconnect()
	if err := conn.QueryRow("select password from users where username =?", logdata.Username).Scan(&realpass); err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"msg": "No such user!"}) //没有该用户或语句错误
		return
	}

	// 校验密码是否正确
	// 若密码正确，返回status值为1并设置cookie
	// 否则，返回status值为0
	if realpass == logdata.Password {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Login success!",
		})
		c.SetCookie(
			"auth", // 设置cookie的key
			base64.StdEncoding.EncodeToString([]byte(logdata.Username)), // 设置cookie的值为用户名的base64值
			6000,        // 过期时间
			"/",         // 所在目录
			"127.0.0.1", //域名
			false,       // 是否只能通过https访问
			true)        // 是否允许别人通过js获取自己的cookie
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Wrong Pass",
		})
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Request.Cookie("auth")
		if cookie == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "请登录"})
		}
	}
}
