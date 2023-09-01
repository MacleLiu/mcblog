package routers

import (
	v1 "mcblog/api/v1"
	"mcblog/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouterInit(r *gin.Engine) {
	// 需要鉴权的资源
	auth := r.Group("/api/v1")
	auth.Use(middlewares.JWTAuth())
	{
		auth.GET("/users", v1.GetUsers)        // 获取用户列表
		auth.GET("/user/:id", v1.GetUser)      // 查询用户信息
		auth.PUT("user/:id", v1.EditUser)      // 修改用户信息
		auth.POST("user/add", v1.AddUser)      // 添加用户
		auth.DELETE("user/:id", v1.DeleteUser) // 删除用户

		auth.POST("category/add", v1.AddCategory)      // 新增分类
		auth.PUT("category/:id", v1.EditCategory)      // 修改分类信息
		auth.DELETE("category/:id", v1.DeleteCategory) // 删除分类

		auth.POST("article/add", v1.AddArticle)      //新增文章
		auth.PUT("article/:id", v1.EditArticle)      //修改文章
		auth.DELETE("article/:id", v1.DeleteArticle) //删除文章

		auth.POST("upload", v1.UpLoad) // 上传文件

		auth.POST("tool/add", v1.AddTool)      // 新增工具
		auth.PUT("tool/:id", v1.EditTool)      // 修改工具信息
		auth.DELETE("tool/:id", v1.DeleteTool) // 删除工具

		auth.POST("wish/add", v1.AddWish)      // 新增心愿
		auth.PUT("wish/:id", v1.EditWish)      // 修改心愿信息
		auth.DELETE("wish/:id", v1.DeleteWish) // 删除心愿

		auth.POST("tag/add", v1.AddTag)                   // 新增标签
		auth.PUT("tag/article/:id", v1.UpdateArticleTags) //修改文章标签
	}

	//公共资源
	router := r.Group("/api/v1")
	{
		router.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "API V1")
		})
		router.POST("login", v1.Login) // 用户登录

		router.GET("/categories", v1.GetCategories) //获取分类列表
		router.GET("/catestat", v1.GetCateStat)     //获取分类统计
		router.GET("/category/:id", v1.GetCategory) // 查询分类信息

		router.GET("/articlecount", v1.GetArticleCount)    //获取文章总数
		router.GET("/articles", v1.GetArticles)            //获取文章列表
		router.GET("/winnow", v1.GetWinnowArticles)        //获取精选文章列表
		router.GET("/allartinfo", v1.GetAllArtInfo)        //获取全部文章的基础信息
		router.GET("article/:id", v1.GetArticle)           //获取指定文章
		router.GET("article/cate/:id", v1.GetCateArticles) //获取一个分类下的文章
		router.GET("article/tag/:id", v1.GetTagArticles)   //获取一个标签下的文章

		router.GET("/tools", v1.GetTools)   //获取工具列表
		router.GET("/tool/:id", v1.GetTool) //获取一个工具信息

		router.GET("/wishes", v1.GetWishes) //获取心愿列表
		router.GET("/wish/:id", v1.GetWish) //获取一个心愿信息

		router.GET("/tags", v1.GetTags)            //获取标签列表
		router.GET("/tag/:id", v1.GetTag)          //获取一个标签信息
		router.GET("/tags/:id", v1.GetArticleTags) //获取指定文章的标签列表
	}
}
