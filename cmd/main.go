package main

import (
	"audit/common"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func setupRouter() *gin.Engine {
	r := gin.New()
	r.Use(common.Logger(), gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "xiao-li is pretty!")
	})

	r.POST("/audit", auditProcess)

	return r
}

func auditProcess(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	log.Info(string(buf[0:n]))
	resp := map[string]string{"hello": "world"}
	c.JSON(http.StatusOK, resp)
}

func main() {
	common.InitLogger()
	r := setupRouter()
	log.Info("audit starting...")
	r.Run(":8888")
}
