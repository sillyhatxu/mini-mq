package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sillyhatxu/convenient-utils/response"
	"github.com/sillyhatxu/mini-mq/dto"
	"github.com/sillyhatxu/mini-mq/service/topic"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"strconv"
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
	usersGroup := router.Group("/users")
	{
		usersGroup.POST("", addUser)
		usersGroup.PUT("/:id", updateUser)
		usersGroup.DELETE("/:id", deleteUser)
		usersGroup.GET("", userList)
		usersGroup.GET("/:id", getUserById)
	}
	topicGroup := router.Group("/topic")
	{
		topicGroup.POST("", addTopic)
		topicGroup.PUT("/:topicName", updateTopic)
		topicGroup.DELETE("/:topicName", deleteTopic)
		topicGroup.GET("", topicList)
		topicGroup.GET("/:topicName", getTopic)
	}
	tgGroup := router.Group("/topic-group")
	{
		tgGroup.GET("", topicGroupList)
	}
	return router
}

func addUser(context *gin.Context) {

}

func updateUser(context *gin.Context) {

}

func deleteUser(context *gin.Context) {

}

func userList(context *gin.Context) {

}

func getUserById(context *gin.Context) {

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
	topicName := context.Param("topicName")
	if topicName == "" {
		context.JSON(http.StatusOK, response.ServerError(nil, "topicName is required", nil))
		return
	}
	q := context.Request.URL.Query()
	offsetReq := q["offset"]
	limitReq := q["limit"]
	var offset int64 = 0
	var limit int = 20
	if offsetReq != nil && len(offsetReq) > 0 {
		i, err := strconv.ParseInt(offsetReq[0], 10, 64)
		if err != nil {
			context.JSON(http.StatusOK, response.ServerError(nil, "offset must be number", nil))
			return
		}
		offset = i
	}
	if limitReq != nil && len(limitReq) > 0 {
		i, err := strconv.Atoi(limitReq[0])
		if err != nil {
			context.JSON(http.StatusOK, response.ServerError(nil, "limit must be number", nil))
			return
		}
		limit = i
	}
	logrus.Infof("[getTopic] topicName [%s]; offset [%d]; limit [%d]", topicName, offset, limit)
	array, count, err := topic.FindTopicData(topicName, offset, limit)
	if err != nil {
		logrus.Errorf("[topicList] topic list error : %v", err)
		context.JSON(http.StatusOK, response.ServerError(nil, err.Error(), nil))
		return
	}
	logrus.Infof("[getTopic] topicName [%s]; offset [%d]; limit [%d] response : %#v ; totalCount : %d", topicName, offset, limit, array, count)
	context.JSON(http.StatusOK, response.ServerSuccess(array, count))
}

func topicList(context *gin.Context) {
	array, err := topic.FindTopicList()
	if err != nil {
		logrus.Errorf("[topicList] topic list error : %v", err)
		context.JSON(http.StatusOK, response.ServerError(nil, err.Error(), nil))
		return
	}
	logrus.Infof("[topicList] response : %#v", array)
	context.JSON(http.StatusOK, response.ServerSuccess(array, nil))
}

func topicGroupList(context *gin.Context) {
	q := context.Request.URL.Query()
	offsetReq := q["offset"]
	limitReq := q["limit"]
	var offset int64 = 0
	var limit int = 20
	if offsetReq != nil && len(offsetReq) > 0 {
		i, err := strconv.ParseInt(offsetReq[0], 10, 64)
		if err != nil {
			context.JSON(http.StatusOK, response.ServerError(nil, "offset must be number", nil))
			return
		}
		offset = i
	}
	if limitReq != nil && len(limitReq) > 0 {
		i, err := strconv.Atoi(limitReq[0])
		if err != nil {
			context.JSON(http.StatusOK, response.ServerError(nil, "limit must be number", nil))
			return
		}
		limit = i
	}
	logrus.Infof("[topicGroupList] offset [%d]; limit [%d]", offset, limit)
	array, err := topic.FindTopicGroupList(offset, limit)
	if err != nil {
		logrus.Errorf("[topicList] topic list error : %v", err)
		context.JSON(http.StatusOK, response.ServerError(nil, err.Error(), nil))
		return
	}
	logrus.Infof("[topicGroupList] response : %#v", array)
	context.JSON(http.StatusOK, response.ServerSuccess(array, nil))
}
