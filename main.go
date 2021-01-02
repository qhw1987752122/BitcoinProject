package main

import (
	"BitcoinProject/db_mysql"
	_ "BitcoinProject/routers"
	"github.com/astaxie/beego"
)

func main() {
	db_mysql.ConnectDB()
	//设置静态资源路径
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}
