package Controllers

import (
	"fmt"
	"gin_api_server/ApiHelpers"
	"gin_api_server/Config"
	"gin_api_server/Models"
	"github.com/gin-gonic/gin"
	"time"
)

type BookController struct {
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleware")
	}
}

func (book *BookController) Router(engine *gin.Engine) {
	//engine.GET("/book", book.Book)
	v1 := engine.Group("/v1")
	v1.Use(Middleware())
	{
		v1.GET("booktest", book.Book)
		v1.GET("login", Login)
		v1.GET("book", ListBook)
		v1.POST("book", AddNewBook)
		v1.GET("book/:id", GetOneBook)
		v1.PUT("book/:id", PutOneBook)
		v1.DELETE("book/:id", DeleteBook)
	}
}

func (book *BookController) Book(contex *gin.Context) {
	cookie, _ := contex.Cookie("sessionId")
	fmt.Println(cookie)
	ApiHelpers.RespondJSON(contex, 200, "booktest")
}

// -----------------------------------------------
func Login(c *gin.Context) {
	cookie, _ := c.Cookie("sessionId")
	fmt.Println(cookie)

	Config.SetStr("123", "value_123", 10*time.Second)
	str := Config.GetStr("123")
	fmt.Println(str)

	c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)

	ApiHelpers.RespondJSON(c, 200, "login")
}

func ListBook(c *gin.Context) {

	var book []Models.Book
	err := Models.GetAllBook(&book)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}

func AddNewBook(c *gin.Context) {
	var book Models.Book
	c.BindJSON(&book)
	err := Models.AddNewBook(&book)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}

func GetOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Models.Book
	err := Models.GetOneBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}

func PutOneBook(c *gin.Context) {
	var book Models.Book
	id := c.Params.ByName("id")
	err := Models.GetOneBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	}
	c.BindJSON(&book)
	err = Models.PutOneBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}

func DeleteBook(c *gin.Context) {
	var book Models.Book
	id := c.Params.ByName("id")
	err := Models.DeleteBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}
