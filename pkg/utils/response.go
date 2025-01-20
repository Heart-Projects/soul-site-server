package utils

import (
	"com.sj/admin/pkg/model"
	"github.com/gin-gonic/gin"
)

func DoResponseWithCondition(c *gin.Context, condition bool, data interface{}, err error) {
	if condition {
		c.JSON(200, model.Response.SuccessWithData(data))
	} else {
		message := "请求失败"
		if err != nil {
			message = err.Error()
		}
		c.JSON(200, model.Response.FailWithMsg(message))
	}
}

func DoResponseSuccess(c *gin.Context, message string) {
	c.JSON(200, model.Response.SuccessWithMsg(message))
}

func DoResponseSuccessWithData(c *gin.Context, data interface{}) {
	c.JSON(200, model.Response.SuccessWithData(data))
}

func DoResponseSuccessMessage(c *gin.Context, data interface{}, message string) {
	c.JSON(200, model.Response.SuccessWithMsgData(message, data))
}

func DoResponseWithConditionMessage(c *gin.Context, condition bool, message string, data interface{}) {
	if condition {
		c.JSON(200, model.Response.SuccessWithMsgData(message, data))
	} else {
		c.JSON(200, model.Response.FailWithMsg(message))
	}
}

func DoResponseWithError(c *gin.Context, err error, data interface{}) {
	if err == nil {
		c.JSON(200, model.Response.SuccessWithData(data))
	} else {
		message := err.Error()
		c.JSON(200, model.Response.FailWithMsg(message))
	}
}

func DoResponseError(c *gin.Context, err error) {
	message := err.Error()
	c.JSON(200, model.Response.FailWithMsg(message))
}

func DoResponseErrorMessage(c *gin.Context, message string) {
	c.JSON(200, model.Response.FailWithMsg(message))
}

func DoResponseWithErrorMessage(c *gin.Context, err error, errorMessage string, data interface{}) {
	if err == nil {
		c.JSON(200, model.Response.SuccessWithData(data))
	} else {
		c.JSON(200, model.Response.FailWithMsg(errorMessage))
	}
}
