package main

import (
	"blog/config"
	"blog/global"
	"blog/utils"
	"log"
)

//Init global variable
//初始化全局变量
func init() {
	global.Config = config.InitConf()

	dbConf := global.Config.DB
	global.DBClient = utils.InitDB(dbConf.Dialect, dbConf.Conn, dbConf.MaxOpenConnections, dbConf.MaxIdleConnections)

	logConf := global.Config.Log
	global.Logger = &utils.Logger{
		MailSender:         nil,
		Reminder:           nil,
		InfoPathFormatter:  logConf.InfoPath,
		WarnPathFormatter:  logConf.WarnPath,
		ErrorPathFormatter: logConf.ErrorPath,
		Level:              logConf.Level,
		Debug:              global.Config.Env == global.DEBUG_MODE,
		SplitType:          logConf.SplitLevel,
		SystemName:         global.Config.SystemName,
	}
	global.Logger.InitLogger()

	log.Println("Start Success.\n 系统启动成功")
}

func main() {
	global.Logger.LogInfo("这是一大段文字阿拉啦啦,%d", 1234)
}
