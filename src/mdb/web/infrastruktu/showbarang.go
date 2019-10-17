package main

import(
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/sql"
)

type databarang struct{
	ID string 
	NamaBarang string
	Harga int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/training")
	if err != nil {
		return nil, err
	}

	return db, nil
}
//show data barang
func sqlQuerydatabarang() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from databarang")
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []databarang

	for rows.Next() {
		var each = databarang{}
		var err = rows.Scan(&each.ID, &each.NamaBarang, &each.Harga)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.ID,each.NamaBarang,each.Harga)
	}
}

func main()  {
	sqlQuerydatabarang()
}
	