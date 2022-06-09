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
	for rows.Next() {
		err := rows.Scan(&f_name, &l_name)
		Checkerror(err)
		stmt := `insert into emp values ($1,$2)`
		res, err := db.Exec(stmt, f_name, l_name)
		fmt.Println(res.LastInsertId())
	}
	// stmt := `insert into emp(first_name,last_name) values ($1,$2)`
	// _, err = db.Exec(stmt, "uday", "salla")
	//fmt.Println(res.LastInsertId())
	defer db.Close()

}
