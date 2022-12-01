package main

import (
	"go_gin/common"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r = CollectRouter(r)
	panic(r.Run(":8085"))
}
