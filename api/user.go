package api

import (
	"Business_Management/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

// 提交客户信息
func SubmitClientInfo(c *gin.Context) {
	var clientInfo *model.Client_Info
	if err := c.ShouldBindBodyWith(&clientInfo, binding.JSON); err != nil {
		fmt.Println("error:", err)
		return
	}

	var request struct {
		Session string `json:"session"`
	}
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		fmt.Println("errorJson:", err)
		c.Abort()
		return
	}
	session := request.Session

	pNumber := model.Sessions[session].PhoneNumber

	err := model.SubmitClientInfo(clientInfo, pNumber)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": 1,
		"msg":    "提交成功",
	})
	return

}

// 查看未完成业务
func GetUnfinished(c *gin.Context) {

	session := c.DefaultQuery("session", "")
	phoneNumber := model.Sessions[session].PhoneNumber

	var Unfinished []model.ViewBusiness
	var err error
	Unfinished, err = model.GetUnfinished(phoneNumber)
	if err != nil {
		fmt.Println("error:", err)
	}
	c.JSON(http.StatusOK, Unfinished)

}

// 获取个人信息
func GetPersonnalInfo(c *gin.Context) {

	session := c.DefaultQuery("session", "")
	phoneNumber := model.Sessions[session].PhoneNumber
	var saleMan *model.SaleMan_Info
	saleMan = model.GetPersonnalInfo(phoneNumber)

	c.JSON(http.StatusOK, gin.H{
		"name":    saleMan.Name,
		"company": saleMan.Company,
		"phone":   saleMan.PhoneNumber,
	})
}

// 更新进度
func UpdateProgress(c *gin.Context) {

	var request struct {
		Session string `json:"session"`
	}
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		fmt.Println("errorJson:", err)
		c.Abort()
		return
	}
	session := request.Session

	pNumber := model.Sessions[session].PhoneNumber

	var PInfo struct {
		ProgressInfo []model.UpdateBusiness `json:"table_data"`
	}
	if err := c.ShouldBindBodyWith(&PInfo, binding.JSON); err != nil {
		fmt.Println("errorJson:", err)
		c.Abort()
		return
	}

	if err := model.UpdateProgress(PInfo.ProgressInfo, pNumber); err != nil {
		fmt.Println("errorJson:", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": 1,
		"msg":    "提交成功",
	})
}

// 查看历史业务
func ViewHistory(c *gin.Context) {
	session := c.DefaultQuery("session", "")
	phoneNumber := model.Sessions[session].PhoneNumber

	var Finished []model.ViewBusiness
	var err error
	Finished, err = model.ViewHistory(phoneNumber)
	if err != nil {
		fmt.Println("error:", err)
	}
	c.JSON(http.StatusOK, Finished)

}

// 查看冲突记录
func ViewConfict(c *gin.Context) {
	session := c.DefaultQuery("session", "")
	phoneNumber := model.Sessions[session].PhoneNumber

	ConflictInfo, err := model.ViewConflict(phoneNumber)
	if err != nil {
		fmt.Println("error:", err)
	}
	c.JSON(http.StatusOK, ConflictInfo)

}
