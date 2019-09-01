package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sillyhatxu/convenient-utils/response"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
)

func InitialAPI(listener net.Listener) {
	logrus.Info("---------- initial internal api start ----------")
	router := SetupRouter()
	//return router.Run(config.Conf.Http.Listen)
	//return router.Run(":8080")
	err := http.Serve(listener, router)
	if err != nil {
		panic(err)
	}
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"message": "OK"}) })
	topicGroup := router.Group("/topic")
	{
		topicGroup.POST("", addTopic)
		topicGroup.PUT("/:topicName", updateTopic)
		topicGroup.DELETE("/:topicName", deleteTopic)
		topicGroup.GET("", topicList)
		topicGroup.GET("/:topicName", getTopic)
	}
	return router
}

func addTopic(context *gin.Context) {
	context.JSON(http.StatusOK, response.ServerSuccess(nil, nil))
}

func updateTopic(context *gin.Context) {
	context.JSON(http.StatusOK, response.ServerSuccess(nil, nil))
}

func getTopic(context *gin.Context) {
	context.JSON(http.StatusOK, response.ServerSuccess(nil, nil))
}

func topicList(context *gin.Context) {
	context.JSON(http.StatusOK, response.ServerSuccess(nil, nil))
}

func deleteTopic(context *gin.Context) {
	context.JSON(http.StatusOK, response.ServerSuccess(nil, nil))
}
