package apis

import (
	"ginstart/global"
	"ginstart/pkg/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HeathCheck(c *gin.Context) {
	// SQL连接检查
	db, _ := global.Db.DB()
	err := db.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errno.ErrServer)
		return
	}
	c.JSON(http.StatusOK, errno.OK)
}
