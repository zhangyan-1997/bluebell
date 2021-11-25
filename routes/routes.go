package routes

import (
	"bluebell/controllers"
	_ "bluebell/docs" // 千万不要忘了导入把你上一步生成的docs
	"bluebell/logger"
	"bluebell/middlewares"
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	//覆盖默认日志
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//注册业务路由
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "请输入有效地址")
	})
	v1 := r.Group("/api/v1")
	//用户登录
	v1.POST("/login", controllers.LoginHandle)
	//用户注册
	v1.POST("/signup", controllers.SignUpHandle)

	//应用JWT认证的中间件
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		//社区分类接口
		v1.GET("/community", controllers.CommunityHandle)
		//社区详情接口
		v1.GET("/community/:id", controllers.CommunityDetailHandler)

		//创建帖子接口
		v1.POST("/post", controllers.CreatePostHandler)
		//帖子详情
		v1.GET("/post/:id", controllers.GetPostDetailHandler)
		//获取帖子列表
		v1.GET("/posts", controllers.GetPostListHandle)
		//投票
		v1.POST("/vote", controllers.PostVoteController)
	}

	r.GET("/ping", func(c *gin.Context) {
		// 如果是登录的用户,判断请求头中是否有 有效的JWT  ？
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})

	//处理错误页面
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
