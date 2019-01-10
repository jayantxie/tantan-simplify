package main

import (
	"fmt"
	"tantan-simplify/config"
	"tantan-simplify/logger"
	"tantan-simplify/model"
	"tantan-simplify/pkg/middleware"
	"tantan-simplify/rest"

	"github.com/gin-gonic/gin"
)

var (
	conf *config.Config
)

func startServer() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(middleware.RequestIDMiddleware, middleware.AccessLogMiddleware(logger.Logger.AccessLogger), gin.RecoveryWithWriter(logger.Logger.ErrorLogger.Out))

	// setup route
	rest.SetupRoutes(engine, logger.Logger.ErrorLogger)

	shutdown := make(chan struct{})
	registerSignal(shutdown, func() {
		logger.ReopenLogs(conf.LogDir)
	})
	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	logger.Logger.ErrorLogger.Infof("Server listening at %s", addr)
	go engine.Run(addr)
	<-shutdown
	logger.Logger.ErrorLogger.Infof("Shutting down...")
}

// every panic causes starting up failed
func setup() {
	var err error

	// load and parse config file
	conf, err = config.MustLoad(flags.confFile)
	if err != nil {
		panic(err)
	}

	// setup log
	err = logger.MustSetup(conf.LogDir, conf.LogLevel, flags.BackTrackLevel)
	if err != nil {
		panic(err)
	}

	// setup postgresql
	err = model.MustSetDB(&conf.PostgreSQL)
	if err != nil {
		panic(err)
	}
}
