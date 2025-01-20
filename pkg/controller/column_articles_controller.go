package controller

import (
	"com.sj/admin/pkg/entity"
	"com.sj/admin/pkg/entity/vo"
	"com.sj/admin/pkg/service"
	"com.sj/admin/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

type IColumnArticlesController interface {
	Add(c *gin.Context)
	Rename(c *gin.Context)
	Remove(c *gin.Context)
	SaveContent(c *gin.Context)
	// ContentItem 获取文章内容
	ContentItem(c *gin.Context)
	TreeData(c *gin.Context)

	ModifyOrder(c *gin.Context)
	ModifyParent(c *gin.Context)
}

type columnArticlesController struct {
	srv       service.IColumnArticlesService
	columnSrv service.IArticleColumnService
}

func NewColumnArticlesController(srv service.IColumnArticlesService, columnSrv service.IArticleColumnService) IColumnArticlesController {
	return &columnArticlesController{
		srv:       srv,
		columnSrv: columnSrv,
	}
}

func (a *columnArticlesController) Add(c *gin.Context) {
	var params vo.ColumnArticlesAddVo
	if err := c.ShouldBind(&params); err != nil {
		logrus.Error(err)
		utils.DoResponseError(c, err)
		return
	}
	if params.ColumnId == 0 && params.ColumnIdentify == "" {
		utils.DoResponseErrorMessage(c, "参数验证失败")
		return
	}
	// 未传递栏目标识时，根据ID 查询
	if params.ColumnId == 0 {
		articleColumn := a.columnSrv.GetOne(0, params.ColumnIdentify)
		if articleColumn == nil {
			utils.DoResponseErrorMessage(c, fmt.Sprintf("栏目 %s 不存在", params.ColumnIdentify))
			return
		}
		params.ColumnId = articleColumn.ID
	}
	success, rerr, newItem := a.srv.New(&params)
	utils.DoResponseWithCondition(c, success, newItem, rerr)
}

func (a *columnArticlesController) Rename(c *gin.Context) {
	var params vo.ColumnArticleRenameVo
	if err := c.ShouldBind(&params); err != nil {
		logrus.Error(err)
		utils.DoResponseErrorMessage(c, "参数验证失败")
		return
	}
	success, rerr := a.srv.Rename(params.Title, params.Id)
	utils.DoResponseWithCondition(c, success, nil, rerr)
}

func (a *columnArticlesController) Remove(c *gin.Context) {
	var idStr = c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id <= 0 {
		utils.DoResponseErrorMessage(c, "参数验证失败")
		return
	}
	success, rerr := a.srv.Remove(id)
	utils.DoResponseWithCondition(c, success, nil, rerr)
}

func (a *columnArticlesController) SaveContent(c *gin.Context) {
	var params vo.ColumnArticleSaveContent
	if err := c.ShouldBind(&params); err != nil {
		logrus.Error(err)
		utils.DoResponseErrorMessage(c, "参数验证失败")
		return
	}
	success, rerr := a.srv.SaveContent(params.Id, params.Content)
	utils.DoResponseWithCondition(c, success, nil, rerr)
}

func (a *columnArticlesController) ContentItem(c *gin.Context) {
	var params vo.ColumnArticleContentVo
	if err := c.ShouldBind(&params); err != nil {
		logrus.Error(err)
		utils.DoResponseErrorMessage(c, "参数验证失败")
		return
	}
	if params.Identify == "" {
		data, done := a.getColumnTocData(c, params.ColumnIdentify)
		if done {
			return
		}
		// 构建栏目的首页，它是一个虚拟的节点
		virtualHome := vo.ColumnArticleVo{Content: data,
			ColumnArticles: &entity.ColumnArticles{
				ParentId:   0,
				Title:      "首页",
				Identify:   "",
				Type:       "index",
				SortNumber: 0,
				BaseModel: entity.BaseModel{
					ID: 0,
				},
			},
		}
		utils.DoResponseSuccessWithData(c, virtualHome)
	} else {
		success, err, data := a.srv.ContentItem(&params)
		utils.DoResponseWithCondition(c, success, data, err)
	}
}

func (a *columnArticlesController) TreeData(c *gin.Context) {
	var columnIdentify = c.Param("columnIdentify")
	treeData, done := a.getColumnTocData(c, columnIdentify)
	if done {
		return
	}
	utils.DoResponseSuccessWithData(c, treeData)
}

// getColumnTocData 获取专栏目录树形数据
func (a *columnArticlesController) getColumnTocData(c *gin.Context, columnIdentify string) ([]*vo.ColumnArticlesListVo, bool) {
	articleColumn := a.columnSrv.GetOne(0, columnIdentify)
	if articleColumn == nil {
		utils.DoResponseErrorMessage(c, fmt.Sprintf("栏目 %s 不存在", columnIdentify))
		return nil, true
	}
	_, articles := a.srv.GetColumnData(articleColumn.ID)
	if len(articles) == 0 {
		utils.DoResponseSuccessWithData(c, make([]utils.TreeItemFeature, 0))
		return nil, true
	}
	vos := make([]*vo.ColumnArticlesListVo, len(articles))
	for i, v := range articles {
		temp := v
		vos[i] = &vo.ColumnArticlesListVo{
			ColumnArticles: &temp,
			Children:       make([]utils.TreeItemFeature, 0),
		}
	}
	treeData := utils.TransformListToTreeData[*vo.ColumnArticlesListVo](vos)
	// sort tree by order
	treeData = utils.SoreTreeData[*vo.ColumnArticlesListVo](treeData, func(i, j utils.TreeItemFeature) bool {
		before := i.(*vo.ColumnArticlesListVo).ColumnArticles
		after := j.(*vo.ColumnArticlesListVo).ColumnArticles
		if before.SortNumber == after.SortNumber {
			return before.CreatedAt.Unix() > after.CreatedAt.Unix()
		}
		return before.SortNumber > after.SortNumber
	})
	return treeData, false
}

func (a *columnArticlesController) ModifyOrder(c *gin.Context) {
	var params vo.ColumnArticleModifyOrderVo
	if err := c.ShouldBind(&params); err != nil {
		logrus.Error(err)
		utils.DoResponseErrorMessage(c, "参数验证失败")
		return
	}
	success, rerr := a.srv.ExchangeOrder(params.SourceItemId, params.TargetItemId)
	utils.DoResponseWithCondition(c, success, nil, rerr)
}

func (a *columnArticlesController) ModifyParent(c *gin.Context) {
	var params vo.ColumnArticleModifyParentVo
	if err := c.ShouldBind(&params); err != nil {
		logrus.Error(err)
		utils.DoResponseErrorMessage(c, "参数验证失败")
		return
	}
	success, rerr := a.srv.ModifyParent(params.Id, params.ParentId)
	utils.DoResponseWithCondition(c, success, nil, rerr)
}
