package main

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/sql"
)
type transaksi struct{
	ID int
	IDbarang string
	NamaBarang string
	Harga int
	Jumlah int
	Total int
	Tanggal int
	Bayar int
	Kembalian int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/training")
	if err != nil {
		return nil, err
	}

	return db, nil
}
//show data transaksi
func sqlQuerytransaksi() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from simpanbarang")
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []transaksi

	for rows.Next() {
		var each = transaksi{}
		var err = rows.Scan(&each.ID, &each.IDbarang, &each.NamaBarang,&each.Harga, &each.Jumlah, &each.Total, &each.Tanggal,&each.Bayar, &each.Kembalian)

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
		fmt.Println(each.ID,each.IDbarang,each.NamaBarang,each.Harga,each.Jumlah,
					each.Total,each.Tanggal,each.Bayar,each.Kembalian)
	}
}
func main()  {
	sqlQuerytransaksi()
}