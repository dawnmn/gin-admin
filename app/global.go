package app

import (
	"gin-admin/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Config config.Server
var Log *logrus.Logger
var Router *gin.Engine
