package api

import (
	"Business_Management/model"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// 注册
func Register(c *gin.Context) {
	var user model.User_info
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "err"})
		return
	}

	existingUser := model.CheckNumber(user.PhoneNumber)
	if existingUser != nil {
		c.JSON(http.StatusOK, gin.H{
			"result": -1,
			"msg":    "该手机号已被注册，若有问题请联系网站管理员",
		})
		return
	}

	err := model.InsertUser(user)
	if err != nil {
		fmt.Println("Error inserting user:", err)

	}
	c.JSON(http.StatusOK, gin.H{
		"result": 1,
		"msg":    "信息已成功提交，等待管理员审核",
	})
}

// 登陆
func Login(c *gin.Context) {

	var request struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println("errorjson:", err)
		return
	}

	password, err, flagAdmin := model.GetPasswrd(request.Account)
	if err != nil {
		fmt.Println(err)
		return
	}

	if password == "" {
		c.JSON(http.StatusOK, gin.H{
			"result": 0,
			"msg":    "用户不存在，请注册",
		})
		return
	}

	//是否审核通过判定
	if model.IsPermitted(request.Account) == false {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": -1,
			"msg":    "审核还未通过",
		})
		return

	}

	decodedPassword, err := model.DecodeBase64(request.Password)
	if err != nil {
		fmt.Println("docode error:", err)
		return
	}

	if strings.TrimSpace(password) != strings.TrimSpace(decodedPassword) {
		c.JSON(http.StatusOK, gin.H{
			"result": -1,
			"msg":    "密码错误，请重新输入",
		})
		return
	}

	// 生成session
	sessionToken := model.GenerateSession()
	session := base64.StdEncoding.EncodeToString([]byte(sessionToken))
	// Store session information with create time
	model.Sessions[session] = model.SessionInfo{
		PhoneNumber:    request.Account,
		ExpirationTime: time.Now().Add(time.Minute * 5),
	}
	//fmt.Println(sessionToken)

	c.JSON(http.StatusOK, gin.H{
		"result":  1,
		"msg":     "登录成功",
		"isadmin": flagAdmin,
		"session": session,
	})

}
