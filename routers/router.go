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
		auth.POST("user/add", v1.AddUser)      //添加注册
		auth.DELETE("user/:id", v1.DeleteUser) // 删除用户

		auth.POST("category/add", v1.AddCategory)      // 新增分类
		auth.PUT("category/:id", v1.EditCategory)      // 修改分类信息
		auth.DELETE("category/:id", v1.DeleteCategory) // 删除分类

		auth.POST("article/add", v1.AddArticle)      //新增文章
		auth.PUT("article/:id", v1.EditArticle)      //修改文章
		auth.DELETE("article/:id", v1.DeleteArticle) //删除文章

		auth.POST("upload", v1.UpLoad) // 上传文件
	}

	//公共资源
	router := r.Group("/api/v1")
	{
		router.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "API V1")
		})
		router.POST("login", v1.Login) // 用户登录

		router.GET("/categories", v1.GetCategories) //获取分类列表
		auth.GET("/category/:id", v1.GetCategory)   // 查询分类信息

		router.GET("/articles", v1.GetArticles)            //获取文章列表
		router.GET("article/:id", v1.GetArticle)           //获取指定文章
		router.GET("article/cate/:id", v1.GetCateArticles) //获取一个分类下的文章
	}
}
