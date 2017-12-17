package main

import (
	"fmt"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func testMysqlDriver() {
	fmt.Println("====== Test mysql driver ======")

	//Connect mysql
	db, err := sql.Open("mysql", "root:rangwojinqu@tcp(172.16.101.128:3306)/test")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	//Test insert
	/*result, err := db.Exec("INSERT INTO mytable VALUES (?, ?, ?, ?)", 0, "techie", 1, 99.99)

	if err != nil {
		fmt.Println(err)
		return
	}

	id, _ := result.LastInsertId()
	fmt.Println("last insert id: ", id)*/

	//Test select
	rows, err := db.Query("SELECT * FROM mytable")

	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next(){
		var id int
		var name string
		var sex int
		var degree float64

		rows.Scan(&id, &name, &sex, &degree)

		fmt.Printf("id:%v, name:%v, sex:%v, degree:%v\r\n", id, name, sex, degree)
	}

	//Test update
	result, err := db.Exec("UPDATE mytable SET name = 'york' WHERE id = 2")

	if err != nil {
		fmt.Println(err)
		return
	}

	count, err := result.RowsAffected()
	fmt.Println("affected rows: ", count)

	//Test delete
	result, err = db.Exec("DELETE FROM mytable WHERE id = 3")

	if err != nil {
		fmt.Println(err)
		return
	}

	count, err = result.RowsAffected()
	fmt.Println("affected rows: ", count)
}

type User struct {
	Id int `xorm:"Int pk autoincr 'id'"`
	Name string `xorm:"char(20) not null 'name'"`
	Sex int `xorm:"Int not null default 0 'sex'"`
	Degree float32 `xorm:"DOUBLE 'degree'"`
}

func (user *User) TableName() string{
	return "mytable"
}

func testXorm() {
	fmt.Println("====== Test mysql xorm ======")

	//New ORM engine using mysql
	engine, err := xorm.NewEngine("mysql", "root:rangwojinqu@tcp(172.16.101.128:3306)/test")

	if err != nil {
		fmt.Println(err)
		return
	}

	//Print all the sql the engine generated
	engine.ShowSQL(true)

	//Sync the struct and db table scheme
	engine.Sync2(new(User))

	//fmt.Println(engine.DBMetas())
	//fmt.Println(engine.IsTableExist("mytable"))
	//fmt.Println(engine.IsTableEmpty(new(User)))

	//Test insert
	fmt.Println("====== Test insert ======")
	user := &User{Name:"liuchao", Sex:1, Degree:98.88}
	affected, err := engine.Insert(user)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Test insert success! affected rows:", affected)

	//Test update
	fmt.Println("====== Test update ======")
	user.Degree = 88.88
	affected, err = engine.Id(2).Update(user)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Test update success! affected rows:", affected)

	//Test get
	fmt.Println("====== Test get ======")
	user = new(User)
	has, err := engine.Id(2).Get(user)

	if err != nil {
		fmt.Println(err)
		return
	}

	if !has {
		fmt.Println("Test get failed! ID Not found")
	} else {
		fmt.Println("Test get success! data:", user)
	}

	//Test delete
	fmt.Println("====== Test delete ======")
	affected, err = engine.Id(3).Delete(user)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Test delete success! affected rows:", affected)
}

func main()  {
	//testMysqlDriver()
	testXorm()
}
