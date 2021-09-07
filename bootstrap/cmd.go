package bootstrap

import (
	"flag"
	"fmt"

	"github.com/bh-qt/alist/conf"
	serv "github.com/bh-qt/alist/server"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	flag.BoolVar(&conf.Debug, "debug", false, "use debug mode")
	flag.BoolVar(&conf.Help, "help", false, "show usage help")
	flag.BoolVar(&conf.Version, "version", false, "show version info")
	flag.StringVar(&conf.ConfigFile, "conf", "conf.yml", "config file")
	flag.BoolVar(&conf.SkipUpdate, "skip-update", false, "skip update")
}

// bootstrap run
func Run() {
	flag.Parse()
	if conf.Help {
		flag.Usage()
		return
	}
	if conf.Version {
		fmt.Println("Current version:" + conf.VERSION)
		return
	}
	start()
}

// print asc
func printASC() {
	log.Info(`
 ________  ___       ___  ________  _________   
|\   __  \|\  \     |\  \|\   ____\|\___   ___\ 
\ \  \|\  \ \  \    \ \  \ \  \___|\|___ \  \_| 
 \ \   __  \ \  \    \ \  \ \_____  \   \ \  \  
  \ \  \ \  \ \  \____\ \  \|____|\  \   \ \  \ 
   \ \__\ \__\ \_______\ \__\____\_\  \   \ \__\
    \|__|\|__|\|_______|\|__|\_________\   \|__|
                            \|_________|
`)
}

// start server
func start() {
	InitLog()
	printASC()
	if !conf.SkipUpdate {
		CheckUpdate()
	}
	if !ReadConf(conf.ConfigFile) {
		log.Errorf("读取配置文件时出现错误,启动失败.")
		return
	}
	InitClient()
	if !InitAliDrive() {
		log.Errorf("初始化阿里云盘出现错误,启动失败.")
		return
	}
	if !InitModel() {
		log.Errorf("初始化数据库出现错误,启动失败.")
		return
	}
	InitCron()
	server()
}

// start http server
func server() {
	baseServer := conf.Conf.Server.Address + ":" + conf.Conf.Server.Port
	r := gin.Default()
	serv.InitRouter(r)
	log.Infof("Starting server @ %s", baseServer)
	err := r.Run(baseServer)
	if err != nil {
		log.Errorf("Server failed start:%s", err.Error())
	}
}
