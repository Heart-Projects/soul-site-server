package controller

import (
	"com.sj/admin/docs"
	"com.sj/admin/pkg/model"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// gin-swagger middleware

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags        example
// @Accept      json
// @Produce     json
// @Success     200 {string} Helloworld
// @Router      /ping [get]
func HelloWord(c *gin.Context) {
	c.JSON(200, model.Response.S())
}

func InitRouters(engine *gin.Engine) {
	//r := gin.Default()
	//
	//// 添加跨域处理
	//r.Use(cors.New(cors.Config{
	//	AllowCredentials: true,
	//	AllowHeaders:     []string{"*"},
	//	AllowMethods:     []string{"POST,GET,OPTIONS,DELETE,PUT"},
	//	AllowAllOrigins:  true,
	//	MaxAge:           12 * time.Hour,
	//}))
	//r.Use(middleware.Auth(r))
	//
	//docs.SwaggerInfo.BasePath = "/"
	//
	//r.GET("/ping", HelloWord)
	//
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//
	//port := os.Getenv("server.port")
	//if "" == port {
	//	port = "9999"
	//}
	//address := "localhost:" + port
	//logrus.Info("Server stated and listen on port ", port)
	//r.Run(address)
	docs.SwaggerInfo.BasePath = "/"
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 注册用户相关路由
	u := getOrCreate().UserController()
	r := engine.Group("/user")
	{
		r.POST("/login", u.login)
		r.GET("/info", u.getUserByToken)
		r.GET("/logout", u.logout)
		// 用户文章数据汇总
		r.GET("/article/summary", u.userArticleSummaryInfo)
	}

	// 这册tag相关路由
	t := getOrCreate().TagController()
	r = engine.Group("/tag")
	{
		r.POST("/add", t.addTag)
	}

	// 注册文章相关路由
	a := getOrCreate().ArticleController()
	r = engine.Group("/article")
	{
		r.POST("/save", a.saveArticle)
		r.GET("/list", a.listUserArticles)
		r.GET("/detail/:userId/:articleId", a.articleDetail)
		r.GET("/edit/detail/:articleId", a.articleItemDetail)
	}

	// 注册文章分类相关路由
	ac := getOrCreate().ArticleCategoryController()
	r = engine.Group("/article-category")
	{
		r.GET("/list", ac.UserArticleCategory)
	}

	// 注册站点category
	sc := getOrCreate().SiteCategoryController()
	r = engine.Group("/site-category")
	{
		r.GET("/list", sc.listAllSiteCategory)
	}

	// 专栏分类
	columnCategory := getOrCreate().ArticleColumnCategoryController()
	r = engine.Group("/column-category")
	{
		r.GET("/tree", columnCategory.treeData)
		r.POST("/add", columnCategory.add)
		r.POST("/rename", columnCategory.rename)
		r.POST("/remove/:id", columnCategory.remove)
	}

	// 专栏
	column := getOrCreate().ArticleColumnController()
	r = engine.Group("/column")
	{
		r.POST("/create", column.create)
		r.POST("/save", column.save)
		r.POST("/remove/:id", column.remove)
		r.GET("/:identify", column.get)
		r.GET("/detail/:id", column.get)
		r.GET("/list", column.list)
	}

	// 专栏文章
	columnArticles := getOrCreate().ColumnArticlesController()
	r = engine.Group("/column-articles")
	{
		r.GET("/tree/:columnIdentify", columnArticles.TreeData)
		r.POST("/add", columnArticles.Add)
		r.POST("/rename", columnArticles.Rename)
		r.POST("/remove/:id", columnArticles.Remove)
		r.POST("/modify-order", columnArticles.ModifyOrder)
		r.POST("/save-content", columnArticles.SaveContent)
		r.GET("/content", columnArticles.ContentItem)
		r.POST("modify-parent", columnArticles.ModifyParent)
	}

	articleContent := getOrCreate().ArticleContentController()
	r = engine.Group("/article-content")
	{
		r.POST("/save", articleContent.Save)
		r.GET("/get", articleContent.Content)
	}

	// 文件上传
	file := getOrCreate().FileController()
	r = engine.Group("/file")
	{
		r.POST("/upload", file.UploadFile)
	}
}
