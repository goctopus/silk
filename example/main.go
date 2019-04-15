package main

import (
	"fmt"
	"github.com/goctopus/silk"
	_ "github.com/goctopus/silk/drivers/sqlite"
	"github.com/goctopus/silk/example/models"
)

func main() {
	silk.Open("sqlite3", "test.db")

	// models下的模型文件需由命令行工具生成
	user := models.Users()
	user.Id = 1
	user.Name = "张三"
	user.Save()

	newUser := models.Users().WhereId(1).First()
	fmt.Println(newUser.Id, newUser.Name)

	user.Id = 2
	user.Name = "张三"
	user.Save()
	newUser = models.Users().WhereId(2).First()
	fmt.Println(newUser.Id, newUser.Name)

	//coll := models.Users().WhereName("张三").Collection()
	//fmt.Println(coll.Take(1).ToJson())

	allUsers := models.Users().WhereName("张三").All()
	fmt.Println(allUsers)

	models.Users().WhereName("张三").Delete()
}
