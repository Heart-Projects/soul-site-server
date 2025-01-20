package controller

import (
	"com.sj/admin/pkg/entity/vo"
	"com.sj/admin/pkg/service"
	"com.sj/admin/pkg/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type IArticleColumnController interface {
	create(c *gin.Context)
	save(c *gin.Context)
	remove(c *gin.Context)
	get(c *gin.Context)
	list(c *gin.Context)
}

type ArticleColumnController struct {
	columnService service.IArticleColumnService
}

func NewArticleColumnController(columnService service.IArticleColumnService) IArticleColumnController {
	return &ArticleColumnController{
		columnService: columnService,
	}
}

func (a *ArticleColumnController) create(c *gin.Context) {
	var params vo.ArticleColumnAddVo
	if err := c.ShouldBind(&params); err != nil {
		utils.DoResponseErrorMessage(c, fmt.Sprintf("参数验证失败,%s", err.Error()))
	}
	_, principal := utils.GetUserInfo(c)
	params.UserId = principal.UserId
	success, err, newItem := a.columnService.New(&params)
	utils.DoResponseWithCondition(c, success, newItem, err)
}

func (a *ArticleColumnController) save(c *gin.Context) {
	var params vo.ArticleColumnUpdateVo
	if err := c.ShouldBind(&params); err != nil {
		utils.DoResponseErrorMessage(c, fmt.Sprintf("参数验证失败,%s", err.Error()))
	}
	success, err := a.columnService.Modify(&params)
	utils.DoResponseWithCondition(c, success, nil, err)
}

func (a *ArticleColumnController) remove(c *gin.Context) {
	removeIdStr := c.Param("id")

	id, err := strconv.ParseUint(removeIdStr, 10, 64)
	if err != nil || id <= 0 {
		utils.DoResponseErrorMessage(c, fmt.Sprintf("参数验证失败,无效的ID %s", removeIdStr))
	}
	success, err := a.columnService.Remove(id)
	utils.DoResponseWithCondition(c, success, nil, err)
}

func (a *ArticleColumnController) get(c *gin.Context) {
	removeIdStr := c.Param("id")
	identify := c.Param("identify")
	id := uint64(0)
	if removeIdStr != "" {
		id, err := strconv.ParseUint(removeIdStr, 10, 64)
		if err != nil || id <= 0 {
			utils.DoResponseErrorMessage(c, fmt.Sprintf("参数验证失败,无效的ID %s", removeIdStr))
		}
	}
	item := a.columnService.GetOne(id, identify)
	utils.DoResponseWithCondition(c, item != nil, item, errors.New("文章栏目不存在"))
}

func (a *ArticleColumnController) list(c *gin.Context) {
	categoryIdStr := c.Query("categoryId")

	categoryId, err := strconv.ParseUint(categoryIdStr, 10, 64)
	if err != nil || categoryId <= 0 {
		utils.DoResponseErrorMessage(c, fmt.Sprintf("参数验证失败,无效的categoryId %s", categoryIdStr))
	}
	items := a.columnService.List(categoryId)
	utils.DoResponseSuccessWithData(c, items)
}
