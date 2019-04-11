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
	user := models.Users()
	user.Id = 1
	user.Name = "张三"
	user.Save()
	fmt.Println(user.Id, user.Name)

	newUser := models.Users().WhereId(1).First()
	fmt.Println(newUser.Id, newUser.Name)

	user.Id = 2
	user.Name = "张三"
	user.Save()

	//coll := models.Users().WhereName("张三").Collection()
	//fmt.Println(coll.Take(1).ToJson())

	models.Users().WhereName("张三").Delete()
}
