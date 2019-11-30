package global

import (
	"blog/config"
	"blog/utils"
	"github.com/jinzhu/gorm"
)

var RedisClient *utils.RedisClient

var DBClient *gorm.DB

var Config config.Config

var Logger *utils.Logger
