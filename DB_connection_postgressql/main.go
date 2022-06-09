package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	f_name    string
	l_name    string
	insrt_qry string
)

func Checkerror(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//connStr := "postgres://uday:uday@localhost:5432/hr?sslmode=disable"
	connStr := "port=5432 host=localhost user=uday password=uday dbname=hr sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	Checkerror(err)

	err = db.Ping()
	Checkerror(err)
	fmt.Println("Connected")
	rows, err := db.Query("Select first_name,last_name from employees")

	insrt_qry = "Insert into emp values "
	// the below commented code is used for single line insertion
	// for rows.Next() {
	// 	err := rows.Scan(&f_name, &l_name)
	// 	Checkerror(err)
	// 	stmt := `insert into emp values ($1,$2)`
	// 	res, err := db.Exec(stmt, f_name, l_name)
	// 	fmt.Println(res.LastInsertId())
	// }

	// below is for multi line insert with single command
	cnt := 1
	for rows.Next() {
		err = rows.Scan(&f_name, &l_name)
		if cnt == 1 {
			insrt_qry = insrt_qry + "\n('" + f_name + "','" + l_name + "')"
			cnt++
		} else {
			insrt_qry = insrt_qry + "\n,('" + f_name + "','" + l_name + "')"
		}
	}
	fmt.Println(insrt_qry)
	_, err = db.Exec(insrt_qry)
	Checkerror(err)
	defer db.Close()

}
