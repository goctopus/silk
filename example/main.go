package main

import (
	_ "github.com/goctopus/silk/dirvers/sqlite"
	"github.com/goctopus/silk"
	"github.com/goctopus/silk/example/models"
	"fmt"
)

func main() {
	silk.Open("sqlite3", "test.db")

	// models下的模型文件需由命令行工具生成
	user := models.NewUsers()
	user.Id = 1
	user.Name = "张三"
	user.Save()
	fmt.Println(user.Id, user.Name)

	newUser := models.NewUsers().WhereId(1).First()
	fmt.Println(newUser.Id, newUser.Name)
}
