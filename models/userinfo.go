package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"time"
)

type User struct {
	Id int
	Name string
	Logininfo string
	Date string	
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:password@tcp(127.0.0.1:3306)/clockin?charset=utf8")
	orm.RegisterModel(new(User))
}

func GetUserInfo() []User {
	o := orm.NewOrm()
	var users [] User
	t := time.Now()
	date := t.Format("2006-01-02") + " 00:00:00"
	_,err := o.Raw("select id, name, logininfo, date from user where date >= ?", date).QueryRows(&users)	
	fmt.Println(err)
	return users
}

func InsertRec(name string) int {
	o := orm.NewOrm()
	var user User
	user.Name = name
	err := o.Read(&user, "name")
	fmt.Println(err)
	if user.Logininfo == "" {
		user.Name = name
		user.Logininfo = "y"
		t := time.Now()
		user.Date = t.Format("2006-01-02 15:04:05")
		_,err := o.Insert(&user)
		if err == nil {
			return 1
		} else {
			return 0
		}
	}
	return 2			
}
