package Routers

import (
	"gin_api_server/Controllers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Cors() gin.HandlerFunc {

	/*
		return func(c *gin.Context) {
			c.Header("Access-Control-Allow-Origin", "*")  // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Max-Age", "86400")

			if c.Request.Method == http.MethodOptions {
				c.AbortWithStatus(200)
			} else {
				c.Next()
			}
		}
	*/
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}

func RegisterRouter(router *gin.Engine) {
	new(Controllers.BookController).Router(router)
	new(Controllers.VersionController).Router(router)
}

func SetupRouter() *gin.Engine {
	//engine := gin.Default()
	engine := gin.New()
	engine.Use(Cors()) //开启中间件 允许使用跨域请求
	RegisterRouter(engine)
	//v1 := r.Group("/v1")
	//{
	//	v1.GET("login", Controllers.Login)
	//	v1.GET("book", Controllers.ListBook)
	//	v1.POST("book", Controllers.AddNewBook)
	//	v1.GET("book/:id", Controllers.GetOneBook)
	//	v1.PUT("book/:id", Controllers.PutOneBook)
	//	v1.DELETE("book/:id", Controllers.DeleteBook)
	//}

	return engine
}
