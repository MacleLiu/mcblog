package routers

import (
	v1 "mcblog/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouterInit(r *gin.Engine) {
	router_v1 := r.Group("/api/v1")
	{
		router_v1.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "API V1")
		})

		//用户模块路由组
		user := router_v1.Group("/user")
		{
			user.POST("/add", v1.AddUser)
			user.GET("/users", v1.GetUsers)
			user.PUT("/:id", v1.EditUser)
			user.DELETE("/:id", v1.DeleteUser)
		}

		//分类模块路由组
		cate := router_v1.Group("/cate")
		{
			cate.POST("/add", v1.AddCategory)
			cate.GET("/cates", v1.GetCategories)
			cate.PUT("/:id", v1.EditCategory)
			cate.DELETE("/:id", v1.DeleteCategory)
		}
		//文章模块路由组
	}
}
