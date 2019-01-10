package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
	"tantan-simplify/errors"
	"tantan-simplify/logger"
	"tantan-simplify/pkg/constant"
	"tantan-simplify/service"
	"tantan-simplify/types"
)

func ListAllUsers(c *gin.Context) {
	httpLogger := logger.GetHTTPLogger(c)

	users, err := service.ListAllUsers()
	if err != nil {
		httpLogger.WithField("error", err.Error()).Warn("List all users failed")
		responseWithItems(c, err)
		return
	}
	responseWithItems(c, users)
}

func CreateUser(c *gin.Context) {
	httpLogger := logger.GetHTTPLogger(c)

	var user types.User
	if err := c.BindJSON(&user); err != nil {
		httpLogger.WithField("error", err.Error()).Warn("Create user failed")
		responseWithItems(c, err)
		return
	}
	userResult, err := service.CreateUser(user.Name)
	if err != nil {
		httpLogger.WithField("error", err.Error()).Warn("Create user failed")
		responseWithItems(c, err)
		return
	}
	responseWithItems(c, userResult)
}

func ListRelationships(c *gin.Context) {
	httpLogger := logger.GetHTTPLogger(c)

	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		httpLogger.WithFields(logrus.Fields{
			"user_id": userId,
			"error":   err.Error(),
		}).Warn("	List relationships failed")
		responseWithError(c, err)
		return
	}

	relationshipHandler := service.GetRelationshipHandler(userId)

	relationships, err := relationshipHandler.ListRelationships()
	if err != nil {
		httpLogger.WithFields(logrus.Fields{
			"user_id": userId,
			"error":   err.Error(),
		}).Warn("	List relationships failed")
		responseWithError(c, err)
		return
	}
	responseWithItems(c, relationships)
}

func UpdateRelationship(c *gin.Context) {
	httpLogger := logger.GetHTTPLogger(c)

	fromUserId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		httpLogger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warn("	Update relationship failed")
		responseWithError(c, err)
		return
	}

	toUserId, err := strconv.Atoi(c.Param("other_user_id"))
	if err != nil {
		httpLogger.WithFields(logrus.Fields{
			"from_user_id": fromUserId,
			"error":        err.Error(),
		}).Warn("	Update relationship failed")
		responseWithError(c, err)
		return
	}

	var relationship types.Relationship
	if err := c.BindJSON(&relationship); err != nil {
		httpLogger.WithFields(logrus.Fields{
			"from_user_id": fromUserId,
			"to_user_id":   toUserId,
			"error":        err.Error(),
		}).Warn("	Update relationship failed")
		responseWithError(c, err)
		return
	}
	// check state
	state := constant.ParseState(relationship.State)
	if state != constant.Liked && state != constant.Disliked {
		err = errors.ErrUnknownState
		httpLogger.WithFields(logrus.Fields{
			"from_user_id": fromUserId,
			"to_user_id":   toUserId,
			"state":        relationship.State,
			"error":        err.Error(),
		}).Warn("	Update relationship failed")
		responseWithError(c, err)
		return
	}

	relationshipHandler := service.GetRelationshipHandler(fromUserId)

	relationshipResult, err := relationshipHandler.UpdateRelationship(toUserId, state)
	if err != nil {
		httpLogger.WithFields(logrus.Fields{
			"from_user_id": fromUserId,
			"to_user_id":   toUserId,
			"state":        relationship.State,
			"error":        err.Error(),
		}).Warn("	Update relationship failed")
		responseWithError(c, err)
		return
	}

	responseWithItems(c, relationshipResult)
}
