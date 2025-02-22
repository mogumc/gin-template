// 启动文件
// @author MoGuQAQ
// @version 1.0.0

package main

import (
	"gin-template/router"
)

func main() {
	router := router.InitRouter()
	address := "127.0.0.1:1088"
	router.Run(address)
}
