package assert

import (
	"com.sj/admin/pkg/model"
	"github.com/gin-gonic/gin"
)

// AssertIsTrue 如果条件为false 时则返回
func AssertIsTrue(c *gin.Context, condition bool, data interface{}, err error) {
	if !condition {
		message := "请求失败"
		if err != nil {
			message = err.Error()
		}
		c.JSON(200, model.Response.FailWithMsg(message))
		return
	}
}
