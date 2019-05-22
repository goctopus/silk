package main

import (
	"fmt"
	"github.com/goctopus/silk"
	_ "github.com/goctopus/silk/drivers/sqlite"
	"github.com/goctopus/silk/example/models"
	"io/ioutil"
)

func main() {

	old, err := ioutil.ReadFile("test.db")
	if err != nil {
		panic(err)
	}

	_, err = silk.Open("sqlite3", "test.db")

	if err != nil {
		panic("invalid connection")
	}

	// models下的模型文件需由命令行工具生成
	user := models.Users()
	user.Name = "张三"
	user.Sex = 0
	user.Country = "中国"
	user.Save()

	newUser := models.Users().WhereName("张三").First()
	fmt.Println("newUser.Id", newUser.Id, "newUser.Name", newUser.Name)

	user.Name = "李四"
	user.Sex = 1
	user.Country = "中国"
	user.Save()
	newUser = models.Users().WhereName("李四").First()
	fmt.Println("newUser.Id", newUser.Id, "newUser.Name", newUser.Name)
	fmt.Println("model", models.Users().Find(newUser.Id))

	coll := models.Users().WhereCountry("中国").Collection()
	fmt.Println("collection", coll.Take(1).ToMapArray())

	allUsers := models.Users().WhereCountry("中国").All()
	fmt.Println("allUsers", allUsers)

	models.Users().WhereCountry("中国").Delete()

	if ioutil.WriteFile("test.db", old, 0644) != nil {
		panic(err)
	}
}
