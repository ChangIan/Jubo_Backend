package main

import (
	"fmt"

	utility "changian.com/jubo/utility"

	global "changian.com/jubo/startup"

	db "changian.com/jubo/helper"

	router "changian.com/jubo/router"
)

// 程式進入點
func main() {
	// 初始化設定檔
	global.Config = utility.GetConfig()
	// 初始化資料庫
	global.Postgresql = db.InitDB()

	// 設定路由器
	router := router.SetRouter()

	url := fmt.Sprintf("%s:%s", global.Config.WebSiteUrl, global.Config.WebSitePort)

	fmt.Println("app config : ", global.Config)
	fmt.Println("web site url : ", url)
	fmt.Println("app server start")

	// 執行
	router.Run(url)
}
