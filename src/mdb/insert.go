package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type barang struct {
	ID         string
	NamaBarang string
	Harga      int
	Jumlah     int
	Total      int
	Bayar      int
	Kembalian  int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/Training")
	if err != nil {
		return nil, err
	}
	return db, nil
	// defer db.Close()
}

// func Routes(){
// 	http.HandleFunc("/register", register)
// }
func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	var result = barang{}
	fmt.Print("masukan ID: ")
	fmt.Scan(&result.ID)
	// fmt.Printf("harga total: %d", result.total)
	err = db.
		QueryRow("select NamaBarang, Harga from murid where id = ?", &result.ID).
		Scan(&result.NamaBarang, &result.Harga)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("masukan jumlah barang: ")
	fmt.Scan(&result.Jumlah)
	result.Total = result.Jumlah * result.Harga
	t := time.Now()
	fmt.Printf("nama barang: %s\nharga: %d\ntotal harga: %d\n", result.NamaBarang, result.Harga, result.Total)
	fmt.Printf("masukan jumlah uang: ")
	fmt.Scan(&result.Bayar)
	result.Kembalian = result.Bayar - result.Total
	fmt.Printf("Kembali :%d\n", result.Kembalian)
	_, err = db.Query("insert into simpanbarang(ID,NamaBarang, Harga,Jumlah,Total,Tanggal,Bayar,Kembalian) values (?,?,?,?,?,?,?,?)", result.ID, result.NamaBarang, result.Harga, result.Jumlah, result.Total, t, result.Bayar, result.Kembalian)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success!")
	defer db.Close()
}

func main() {
	sqlExec()
}
