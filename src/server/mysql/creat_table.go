package main

import (
	_ "encoding/binary"
	"fmt"
	"frame"
	"frame/def"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "io/ioutil"
	_ "math"
	_ "math/rand"
	_ "runtime"
	_ "time"
)

var gDB *gorm.DB

func init() {
	var err error
	gDB, err = connectMysql()
	if err != nil {
		panic(fmt.Sprintf("Connect mysql error = %v\n", err))
	}
	fmt.Println("Connect success")
	gDB.LogMode(false)
	gDB.DB().SetMaxIdleConns(10)
}

func connectMysql() (db *gorm.DB, err error) {
	user := frame.GetDbUser()
	pswd := frame.GetDbPswd()
	host := frame.GetDbHost()
	port := frame.GetDbPort()
	database := "game"
	login := user + ":" + pswd + "@(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
    fmt.Println(login)
	db, err = gorm.Open("mysql", login)
	return
}

func initTable() {
	if ok := gDB.HasTable(&def.User{}); ok {
        gDB.DropTable(&def.User{})
		gDB.CreateTable(&def.User{})
	}
	if ok := gDB.HasTable(&def.NormalCraft{}); ok {
        gDB.DropTable(&def.NormalCraft{})
		gDB.CreateTable(&def.NormalCraft{})
	}
	if ok := gDB.HasTable(&def.GoodCraft{}); ok {
        gDB.DropTable(&def.GoodCraft{})
		gDB.CreateTable(&def.GoodCraft{})
	}
}

func initData() {
	for i := 0; i < 10; i++ {
		user := def.User{
            Gold: 0,
            Info: "",
		}
		gDB.Create(&user)
	}
}

func main() {
    fmt.Println("this mysql tool")
	initTable()
	//initData()
}
