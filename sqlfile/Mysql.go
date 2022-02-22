package sqlfile

import (
	"fmt"
	Model "ginGO/Models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init() {
	host := "139.155.88.241"
	port := "3388"
	database := "oneweb"
	username := "root"
	password := "123456"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	dataDase, err := gorm.Open("mysql", dsn)
	//db, err := gorm.Open("mysql", "root:123456@/(139.155.88.241:3388)/testname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("有误", err)
		return
	}
	Db = dataDase
	Db.AutoMigrate(Model.Login{}, Model.WebTitle{}, Model.History{})
	//Db.Create(&st)
}
