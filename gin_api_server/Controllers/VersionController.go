package Controllers

import (
	"fmt"
	"gin_api_server/ApiHelpers"
	"github.com/gin-gonic/gin"
	//"github.com/stretchr/testify/http"
)

type VersionController struct {
}

func (version *VersionController) Router(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	{
		v1.GET("versiontest", version.Version)
	}
}

func (version *VersionController) Version(contex *gin.Context) {
	cookie, _ := contex.Cookie("sessionId")
	fmt.Println(cookie)
	ApiHelpers.RespondJSON(contex, 200, "versiontest")
	//contex.JSON(http.StatusOK, gin.H{
	//	"message": "",
	//	"method":  "DELETE",
	//})
}
