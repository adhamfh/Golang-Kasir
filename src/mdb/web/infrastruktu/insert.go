package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type barang struct {
	IDBarang   string
	NamaBarang string
	Harga      int
	Jumlah     int
	Total      int
	Bayar      int
	Kembalian  int
	menu       string
	// TotalBayar int
}
type menuawal struct {
	IDBarang   string
	NamaBarang string
	Harga      int
	menu       string
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
func TambahBarang() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	var tambah = barang{}
	fmt.Print("Masukan ID Barang :")
	fmt.Scan(&tambah.IDBarang)
	fmt.Print("Masukan Nama Barang :")
	fmt.Scan(&tambah.NamaBarang)
	fmt.Print("Masukan Harga Barang :")
	fmt.Scan(&tambah.Harga)
	_, err = db.Query("insert into databarang(ID,NamaBarang, Harga) values (?,?,?)", tambah.IDBarang, tambah.NamaBarang, tambah.Harga)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success!")
	fmt.Printf("Apakah ada barang lain ? (Y/T)")
	fmt.Scan(&tambah.menu)
	switch tambah.menu {
	case "Y", "y":
		sqlExec()
	case "T", "t":
		break
	}
	defer db.Close()

}
func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	var result = barang{}
	fmt.Print("masukan ID: ")
	fmt.Scan(&result.IDBarang)
	// fmt.Printf("harga total: %d", result.total)
	err = db.QueryRow("select NamaBarang, Harga from databarang where id = ?", &result.IDBarang).
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
	fmt.Printf("Apakah ada barang lain ? (Y/T)")
	fmt.Scan(&result.menu)
	switch result.menu {
	case "Y", "y":
		sqlExec()
	case "T", "t":

		// fmt.Printf("total bayar :%d\n", result.TotalBayar)
		break
	default:
		fmt.Print("pilihan salah")
	}

	fmt.Printf("masukan jumlah uang: ")
	fmt.Scan(&result.Bayar)
	result.Kembalian = result.Bayar - result.Total
	fmt.Printf("Kembali :%d\n", result.Kembalian)
	_, err = db.Query("insert into simpanbarang(IDBarang,NamaBarang, Harga,Jumlah,Total,Tanggal,Bayar,Kembalian) values (?,?,?,?,?,?,?,?)", result.IDBarang, result.NamaBarang, result.Harga, result.Jumlah, result.Total, t, result.Bayar, result.Kembalian)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success!")
	defer db.Close()
}

func main() {
	var awal = menuawal{}
	fmt.Println("===Menu Utama===")
	fmt.Println("1. Tambah Barang")
	fmt.Println("2. Transaksi")
	fmt.Print("masukan pilihan :")
	fmt.Scan(&awal.menu)
	switch awal.menu {
	case "1":
		TambahBarang()
	case "2":
		sqlExec()
	default:
		fmt.Print("salah")
	}

}
