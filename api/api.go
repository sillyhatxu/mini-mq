package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sillyhatxu/convenient-utils/response"
	"github.com/sillyhatxu/mini-mq/dto"
	"github.com/sillyhatxu/mini-mq/service/topic"
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
	var requestBody dto.AddTopicDTO
	err := context.ShouldBindJSON(&requestBody)
	if err != nil {
		context.JSON(http.StatusOK, response.ServerParamsValidateError(nil, err.Error()))
		return
	}
	logrus.Infof("[addTopic] requestBody : %#v", requestBody)
	err = requestBody.Validate()
	if err != nil {
		context.JSON(http.StatusOK, response.ServerParamsValidateError(nil, err.Error()))
		return
	}
	err = topic.CreateTopic(requestBody.TopicName)
	if err != nil {
		logrus.Errorf("[addTopic] create topic error : %v", err)
		context.JSON(http.StatusOK, response.ServerError(nil, err.Error(), nil))
		return
	}
	context.JSON(http.StatusOK, response.ServerSuccess(nil, nil))
}

func updateTopic(context *gin.Context) {
	topicName := context.Param("topicName")
	var requestBody dto.UpdateTopicDTO
	err := context.ShouldBindJSON(&requestBody)
	if err != nil {
		context.JSON(http.StatusOK, response.ServerParamsValidateError(nil, err.Error()))
		return
	}
	logrus.Infof("[updateTopic] topicName [%s] requestBody : %#v", topicName, requestBody)
	err = requestBody.Validate()
	if err != nil {
		context.JSON(http.StatusOK, response.ServerParamsValidateError(nil, err.Error()))
		return
	}
	err = topic.UpdateTopic(topicName, requestBody.Offset)
	if err != nil {
		logrus.Errorf("[updateTopic] update topic error : %v", err)
		context.JSON(http.StatusOK, response.ServerError(nil, err.Error(), nil))
		return
	}
	context.JSON(http.StatusOK, response.ServerSuccess(nil, nil))
}

func deleteTopic(context *gin.Context) {
	topicName := context.Param("topicName")
	logrus.Infof("[deleteTopic] topicName [%s]", topicName)
	err := topic.DeleteTopic(topicName)
	if err != nil {
		logrus.Errorf("[deleteTopic] delete topic error : %v", err)
		context.JSON(http.StatusOK, response.ServerError(nil, err.Error(), nil))
		return
	}
	context.JSON(http.StatusOK, response.ServerSuccess(nil, nil))
}

func getTopic(context *gin.Context) {
	context.JSON(http.StatusOK, response.ServerSuccess(nil, nil))
}

func topicList(context *gin.Context) {

	context.JSON(http.StatusOK, response.ServerSuccess(nil, nil))
}
