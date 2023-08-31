package routes

import (
	"Business_Management/api"
	"Business_Management/middleware"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "wechat/admin.html")
	p.AddFromFiles("user", "wechat/user.html")
	p.AddFromFiles("login", "wechat/login.html")

	return p
}

func InitRouter() {

	r := gin.Default()
	r.Static("/wechat", "./wechat")

	r.HTMLRender = createMyRender()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login", nil)
	})

	r.GET("userPage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user", nil)
	})

	r.GET("adminPage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin", nil)
	})

	//user 注册
	r.POST("register", api.Register)

	//user登陆
	r.POST("login", api.Login)

	//user 查看个人信息
	r.GET("infoUser", middleware.SessionMiddleWare(), api.GetPersonnalInfo)

	//user查看未完成业务
	r.GET("modifyUser", middleware.SessionMiddleWare(), api.GetUnfinished)

	//user查看历史业务
	r.GET("historyUser", middleware.SessionMiddleWare(), api.ViewHistory)

	//user 提交客户信息
	r.POST("submitClient", middleware.SessionMiddleWare(), api.SubmitClientInfo)

	//user 提交业务进展情况
	r.POST("submitProgress", middleware.SessionMiddleWare(), api.UpdateProgress)

	//user 查看冲突消息
	r.GET("conflictUser", middleware.SessionMiddleWare(), api.ViewConfict)

	//admin 查看历史业务
	r.GET("historyAdmin", middleware.SessionMiddleWare(), api.ViewHistoryAdmin)

	//admin 查看业务员表
	r.GET("infoAdmin", middleware.SessionMiddleWare(), api.InfoSaleMan)

	//admin 查看注册申请
	r.GET("addAdmin", middleware.SessionMiddleWare(), api.ViewSaleman)

	//admin 审核注册信息
	r.POST("submitApply", middleware.SessionMiddleWare(), api.SubmitApplyHandler)

	//admin 查看冲突消息
	r.GET("conflictAdmin", middleware.SessionMiddleWare(), api.ViewConflictAdmin)

	r.Run(":8080")
}
