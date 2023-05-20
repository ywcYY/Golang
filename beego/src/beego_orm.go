package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}
type Post struct {
	Id    int    `orm:"auto"`
	Title string `orm:"size(100)"`
	User  *User  `orm:"rel(fk)"`
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123789@tcp(127.0.0.1:3306)/beego_db?charset=utf8&loc=Local")

	// register model
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Post))
	// create table
	orm.RunSyncdb("default", false, true)
}

func insert(o orm.Ormer) {
	// insert
	user := User{Name: "slene2"}
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
}

func update(o orm.Ormer) {
	// update
	user := User{Name: "slene"}
	user.Name = "astaxie"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

}
func read(o orm.Ormer) {
	// read one
	user := User{Name: "slene"}
	u := User{Id: user.Id}
	err := o.Read(&u)
	fmt.Printf("ERR: %v\n", err)
}
func delete(o orm.Ormer) {

	// delete
	user := User{Name: "slene"}
	u := User{Id: user.Id}
	num, err := o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
func main() {

	o := orm.NewOrm()
	queryTable(o)

}
func queryTable(o orm.Ormer) {

	user := User{Name: "slene"}
	u := User{Id: user.Id}
	o.Insert(&Post{Id: 1, Title: "ywc", User: &u})
	var posts []*Post
	qs := o.QueryTable("post")
	num, err := qs.Filter("User__Name", "slene").All(&posts)
	fmt.Println(num, err)

}
