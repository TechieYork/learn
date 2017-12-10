package main

import (
	"fmt"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
func main()  {
	fmt.Println("====== Test mysql ======")

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
