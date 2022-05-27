package main

import (
	"fmt"
	"gin_api_server/Config"
	"gin_api_server/Models"
	"gin_api_server/Routers"
	"gin_api_server/base_info"
	"gin_api_server/handles"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

var err error

func main() {
	fmt.Println("hello gin api server.")
	fmt.Println("Ver:" + base_info.GetVersion())

	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	//router := gin.Default()
	//
	//router.GET("/someGet", getting)
	//router.POST("/somePost", posting)
	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	//router.Run()
	// router.Run(":3000") for a hard coded por

	//router := gin.Default()
	//
	//// This handler will match /user/john but will not match /user/ or /user
	//router.GET("/user/:name", func(c *gin.Context) {
	//	name := c.Param("name")
	//	c.String(http.StatusOK, "Hello %s", name)
	//})
	//
	//// However, this one will match /user/john/ and also /user/john/send
	//// If no other routers match /user/john, it will redirect to /user/john/
	//router.GET("/user/:name/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	message := name + " is " + action
	//	c.String(http.StatusOK, message)
	//})
	//
	//// For each matched request Context will hold the route definition
	//router.POST("/user/:name/*action", func(c *gin.Context) {
	//	b := c.FullPath() == "/user/:name/*action" // true
	//	c.String(http.StatusOK, "%t", b)
	//})
	//
	//// This handler will add a new router for /user/groups.
	//// Exact routes are resolved before param routes, regardless of the order they were defined.
	//// Routes starting with /user/groups are never interpreted as /user/:name/... routes
	//router.GET("/user/groups", func(c *gin.Context) {
	//	c.String(http.StatusOK, "The available groups are [...]")
	//})
	//
	//router.Run(":8080")
	//router := gin.Default()
	//
	//// Query string parameters are parsed using the existing underlying request object.
	//// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	//router.GET("/welcome", func(c *gin.Context) {
	//	firstname := c.DefaultQuery("firstname", "Guest")
	//	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	//
	//	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	//})
	//router.Run(":8080")
	//router := gin.Default()
	//
	//router.POST("/form_post", func(c *gin.Context) {
	//	message := c.PostForm("message")
	//	nick := c.DefaultPostForm("nick", "anonymous")
	//
	//	c.JSON(200, gin.H{
	//		"status":  "posted",
	//		"message": message,
	//		"nick":    nick,
	//	})
	//})
	//router.Run(":8080")
	//router := gin.Default()
	//
	//router.POST("/post", func(c *gin.Context) {
	//
	//	id := c.Query("id")
	//	page := c.DefaultQuery("page", "0")
	//	name := c.PostForm("name")
	//	message := c.PostForm("message")
	//
	//	fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	//})
	//router.Run(":8080")
	//router := gin.Default()
	//
	//router.POST("/post", func(c *gin.Context) {
	//
	//	ids := c.QueryMap("ids")
	//	names := c.PostFormMap("names")
	//
	//	fmt.Printf("ids: %v; names: %v", ids, names)
	//})
	//router.Run(":8080")
	//router := gin.Default()
	//// Set a lower memory limit for multipart forms (default is 32 MiB)
	//router.MaxMultipartMemory = 8 << 20 // 8 MiB
	//router.POST("/upload", func(c *gin.Context) {
	//	// Single file
	//	file, _ := c.FormFile("file")
	//	log.Println(file.Filename)
	//
	//	// Upload the file to specific dst.
	//	c.SaveUploadedFile(file, dst)
	//
	//	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	//})
	//router.Run(":8080")
	//router := gin.Default()
	//// Set a lower memory limit for multipart forms (default is 32 MiB)
	//router.MaxMultipartMemory = 8 << 20 // 8 MiB
	//router.POST("/upload", func(c *gin.Context) {
	//	// Multipart form
	//	form, _ := c.MultipartForm()
	//	files := form.File["upload[]"]
	//
	//	for _, file := range files {
	//		log.Println(file.Filename)
	//
	//		// Upload the file to specific dst.
	//		c.SaveUploadedFile(file, dst)
	//	}
	//	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	//})
	//router.Run(":8080")
	//router := gin.Default()

	// Simple group: v1
	//v1 := router.Group("/v1")
	//{
	//	v1.POST("/login", loginEndpoint)
	//	v1.POST("/submit", submitEndpoint)
	//	v1.POST("/read", readEndpoint)
	//}
	//
	//// Simple group: v2
	//v2 := router.Group("/v2")
	//{
	//	v2.POST("/login", loginEndpoint)
	//	v2.POST("/submit", submitEndpoint)
	//	v2.POST("/read", readEndpoint)
	//}
	//
	//router.Run(":8080")
	// Creates a router without any middleware by default
	//r := gin.New()
	//
	//// Global middleware
	//// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	//// By default gin.DefaultWriter = os.Stdout
	//r.Use(gin.Logger())
	//
	//// Recovery middleware recovers from any panics and writes a 500 if there was one.
	//r.Use(gin.Recovery())
	//
	//// Per route middleware, you can add as many as you desire.
	//r.GET("/benchmark", MyBenchLogger(), benchEndpoint)
	//
	//// Authorization group
	//// authorized := r.Group("/", AuthRequired())
	//// exactly the same as:
	//authorized := r.Group("/")
	//// per group middleware! in this case we use the custom created
	//// AuthRequired() middleware just in the "authorized" group.
	//authorized.Use(AuthRequired())
	//{
	//	authorized.POST("/login", loginEndpoint)
	//	authorized.POST("/submit", submitEndpoint)
	//	authorized.POST("/read", readEndpoint)
	//
	//	// nested group
	//	testing := authorized.Group("testing")
	//	// visit 0.0.0.0:8080/testing/analytics
	//	testing.GET("/analytics", analyticsEndpoint)
	//}
	//
	//// Listen and serve on 0.0.0.0:8080
	//r.Run(":8080")
	// Creates a router without any middleware by default
	//r := gin.New()
	//
	//// Global middleware
	//// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	//// By default gin.DefaultWriter = os.Stdout
	//r.Use(gin.Logger())
	//
	//// Recovery middleware recovers from any panics and writes a 500 if there was one.
	//r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
	//	if err, ok := recovered.(string); ok {
	//		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	//	}
	//	c.AbortWithStatus(http.StatusInternalServerError)
	//}))
	//
	//r.GET("/panic", func(c *gin.Context) {
	//	// panic with a string -- the custom middleware could save this to a database or report it to the user
	//	panic("foo")
	//})
	//
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "ohai")
	//})
	//
	//// Listen and serve on 0.0.0.0:8080
	//r.Run(":8080")
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//
	//// Use the following code if you need to write the logs to file and console at the same time.
	//// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	//
	//router := gin.Default()
	//router.GET("/ping", func(c *gin.Context) {
	//	c.String(200, "pong")
	//})
	//
	//router.Run(":8080")
	//router := gin.New()
	//
	//// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	//// By default gin.DefaultWriter = os.Stdout
	//router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	//
	//	// your custom format
	//	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	//		param.ClientIP,
	//		param.TimeStamp.Format(time.RFC1123),
	//		param.Method,
	//		param.Path,
	//		param.Request.Proto,
	//		param.StatusCode,
	//		param.Latency,
	//		param.Request.UserAgent(),
	//		param.ErrorMessage,
	//	)
	//}))
	//router.Use(gin.Recovery())
	//
	//router.GET("/ping", func(c *gin.Context) {
	//	c.String(200, "pong")
	//})
	//
	//router.Run(":8080")
	// Disable log's color
	//gin.DisableConsoleColor()
	//
	//// Creates a gin router with default middleware:
	//// logger and recovery (crash-free) middleware
	//router := gin.Default()
	//
	//router.GET("/ping", func(c *gin.Context) {
	//	c.String(200, "pong")
	//})
	//
	//router.Run(":8080")
	// Force log's color
	//gin.ForceConsoleColor()
	//
	//// Creates a gin router with default middleware:
	//// logger and recovery (crash-free) middleware
	//router := gin.Default()
	//
	//router.GET("/ping", func(c *gin.Context) {
	//	c.String(200, "pong")
	//})
	//
	//router.Run(":8080")
	//router := gin.Default()
	//
	//// Example for binding JSON ({"user": "manu", "password": "123"})
	//router.POST("/loginJSON", func(c *gin.Context) {
	//	var json Login
	//	if err := c.ShouldBindJSON(&json); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		return
	//	}
	//
	//	if json.User != "manu" || json.Password != "123" {
	//		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	//		return
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	//})
	//
	//// Example for binding XML (
	////	<?xml version="1.0" encoding="UTF-8"?>
	////	<root>
	////		<user>manu</user>
	////		<password>123</password>
	////	</root>)
	//router.POST("/loginXML", func(c *gin.Context) {
	//	var xml Login
	//	if err := c.ShouldBindXML(&xml); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		return
	//	}
	//
	//	if xml.User != "manu" || xml.Password != "123" {
	//		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	//		return
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	//})
	//
	//// Example for binding a HTML form (user=manu&password=123)
	//router.POST("/loginForm", func(c *gin.Context) {
	//	var form Login
	//	// This will infer what binder to use depending on the content-type header.
	//	if err := c.ShouldBind(&form); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		return
	//	}
	//
	//	if form.User != "manu" || form.Password != "123" {
	//		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	//		return
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	//})
	//
	//// Listen and serve on 0.0.0.0:8080
	//router.Run(":8080")
	// Creates a router without any middleware by default
	//r := gin.New()
	//
	//// Global middleware
	//// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	//// By default gin.DefaultWriter = os.Stdout
	//r.Use(gin.Logger())
	//
	//// Recovery middleware recovers from any panics and writes a 500 if there was one.
	//r.Use(gin.Recovery())
	//
	//// Per route middleware, you can add as many as you desire.
	//r.GET("/benchmark", MyBenchLogger(), benchEndpoint)
	//
	//// Authorization group
	//// authorized := r.Group("/", AuthRequired())
	//// exactly the same as:
	//authorized := r.Group("/")
	//// per group middleware! in this case we use the custom created
	//// AuthRequired() middleware just in the "authorized" group.
	//authorized.Use(AuthRequired())
	//{
	//	authorized.POST("/login", loginEndpoint)
	//	authorized.POST("/submit", submitEndpoint)
	//	authorized.POST("/read", readEndpoint)
	//
	//	// nested group
	//	testing := authorized.Group("testing")
	//	// visit 0.0.0.0:8080/testing/analytics
	//	testing.GET("/analytics", analyticsEndpoint)
	//}
	//
	//// Listen and serve on 0.0.0.0:8080
	//r.Run(":8080")

	// 创建记录日志的文件
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	base_info.WriteLog("hello gin.")

	Cfg, err := ini.Load("conf/app.conf") //此路径为ini文件的路径
	if err != nil {
		fmt.Println("读取ini文件失败")
	}
	section := Cfg.Section("server")          //读取名字为server的区域部分
	port := section.Key("httpport").MustInt() //选择区域内的键并指定类型
	portStr := fmt.Sprintf("%d", port)
	fmt.Println("port:" + portStr) //最后输出8085
	Config.RedisInit()

	Config.DB, err = gorm.Open("mysql", "root:1q2w3e4r@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Book{})

	//router := Routers.SetupRouter()
	//// running method 1
	////r.Run(":8080")
	//// running method 2
	////http.ListenAndServe(":8080", r)
	//// running method 3
	//server := &http.Server{
	//	Addr:           ":8080",
	//	Handler:        router,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//server.ListenAndServe()

	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		err := server01.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	g.Go(func() error {
		err := server02.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

var (
	g errgroup.Group
)

func router01() http.Handler {
	//e := gin.New()
	e := Routers.SetupRouter()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01",
			},
		)
	})

	return e
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 02",
			},
		)
	})

	//r := gin.Default()
	greeter := e.Group("v1")
	//业务调用
	//匹配路由进行返回回调进入到CallGreeter方法里面处理
	greeter.GET("/callsay", handles.CallGreeter)

	//if err := e.Run(); err != nil {
	//	panic(err)
	//}

	return e
}
