package api

import (
	"Business_Management/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

// 查看注册申请
func ViewSaleman(c *gin.Context) {
	Applicants, err := model.GetApplicantInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, Applicants)
}

// 通过申请
func SubmitApplyHandler(c *gin.Context) {
	var applycant struct {
		Session    string                `json:"session"`
		Table_data []model.ApplicantInfo `json:"table_data"`
	}
	if err := c.ShouldBindBodyWith(&applycant, binding.JSON); err != nil {
		fmt.Println("error1:", err)
		//c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	if err := model.UpdatePermitted(applycant.Table_data); err != nil {
		fmt.Println("error2:", err)
		return
	}
	fmt.Println("hahahha2")

	c.JSON(http.StatusOK, gin.H{"result": 1, "msg": "提交成功"})
}

// 查看历史业务
func ViewHistoryAdmin(c *gin.Context) {
	var Business []model.HistoryBusiness
	var err error
	if Business, err = model.ViewHistoryAdmin(); err != nil {
		fmt.Println("error:", err)
	}
	c.JSON(http.StatusOK, Business)
}

// 查看业务员列表
func InfoSaleMan(c *gin.Context) {
	var SMan []model.SaleMan
	var err error
	if SMan, err = model.InfoSaleMan(); err != nil {
		fmt.Println("error:", err)
	}
	c.JSON(http.StatusOK, SMan)
}

// 查看冲突记录
func ViewConflictAdmin(c *gin.Context) {
	var ConflictInfo []model.ConflictAdmin
	var err error
	if ConflictInfo, err = model.ViewConflictAdmin(); err != nil {
		fmt.Println("error:", err)
	}
	c.JSON(http.StatusOK, ConflictInfo)
}
