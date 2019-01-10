package devops

import (
	"net/http"
	"tantan-simplify/version"

	"github.com/gin-gonic/gin"
)

func Version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version":      version.Version,
		"build_commit": version.BuildCommit,
		"build_date":   version.BuildDate,
	})
}
