package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"tantan-simplify/devops"
)

func SetupRoutes(e *gin.Engine, logger *logrus.Logger) {
	// devops interface: health checking
	e.GET("/devops/status", devops.CheckHealth)

	// devops interface: get current version
	e.GET("/devops/version", devops.Version)

	// List all users
	e.GET("/users", ListAllUsers)

	// Create a user
	e.POST("/users", CreateUser)

	// List a users all relationships
	e.GET("/users/:user_id/relationships", ListRelationships)

	// Create/Update relationship state to another user
	e.PUT("/users/:user_id/relationships/:other_user_id", UpdateRelationship)
}
