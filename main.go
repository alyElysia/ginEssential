package main

import (
	"ginEssential/common"
	"ginEssential/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = routes.CollectRoutes(r)
	panic(r.Run(":8080"))
}
