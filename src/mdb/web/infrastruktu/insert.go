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
type login struct {
	User string
	Pass string
}
type databarang struct {
	ID         string
	NamaBarang string
	Harga      int
}
type transaksi struct {
	ID         int
	IDbarang   string
	NamaBarang string
	Harga      int
	Jumlah     int
	Total      int
	Tanggal    int
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
		var err = rows.Scan(&each.ID, &each.IDbarang, &each.NamaBarang, &each.Harga, &each.Jumlah, &each.Total, &each.Tanggal, &each.Bayar, &each.Kembalian)

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
		fmt.Println(each.ID, each.IDbarang, each.NamaBarang, each.Harga, each.Jumlah,
			each.Total, each.Tanggal, each.Bayar, each.Kembalian)
	}
}
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
		fmt.Println(each.ID, each.NamaBarang, each.Harga)
	}
}
func loginn() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	var l = login{}
	fmt.Print("Masukan username: ")
	fmt.Scan(&l.User)
	err = db.QueryRow("select ID,passw from login where uname=?", &l.User).Scan(&l.User, &l.Pass)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	fmt.Print("Masukan password: ")
	fmt.Scan(&l.Pass)
	err = db.QueryRow("select ID, uname from login where passw=?", &l.Pass).Scan(&l.User, &l.Pass)

	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	// rows, err := db.Query("select passw from login")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	defer db.Close()

	// var resultt []login

	// for rows.Next() {
	// 	var eachh = login{}
	// 	var err = rows.Scan(&eachh.User, &eachh.Pass)

	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		return
	// 	}

	// 	resultt = append(resultt, eachh)
	// }

	// if err = rows.Err(); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
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
	// w := t.Format("15.04")
	fmt.Printf("nama barang: %s\nharga: %d\ntotal harga: %d\n", result.NamaBarang, result.Harga, result.Total)
	// fmt.Printf("Apakah ada barang lain ? (Y/T)")
	// fmt.Scan(&result.menu)
	// switch result.menu {
	// case "Y", "y":
	// 	sqlExec()
	// case "T", "t":

	// 	// fmt.Printf("total bayar :%d\n", result.TotalBayar)
	// 	break
	// default:
	// 	fmt.Print("pilihan salah")
	// }

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
	// var l = login{}
	var awal = menuawal{}

	fmt.Println("===Menu Utama===")
	fmt.Println("1. Tambah Barang")
	fmt.Println("2. Transaksi")
	fmt.Println("3. Lihat Daftar Barang")
	fmt.Println("4. Lihat Daftar Transaksi")
	fmt.Print("masukan pilihan :")
	fmt.Scan(&awal.menu)
	switch awal.menu {
	case "1":
		loginn()
		// fmt.Println("===Login===")
		// fmt.Print("Username :")
		// fmt.Scan(&l.User)
		// fmt.Print("Password :")
		// fmt.Scan(&l.Pass)
		// if l.User == "admin" && l.Pass == "admin" {
		TambahBarang()
		// }
	case "2":
		loginn()
		// fmt.Println("===Login===")
		// fmt.Print("Username :")
		// fmt.Scan(&l.User)
		// fmt.Print("Password :")
		// fmt.Scan(&l.Pass)
		// if l.User == "admin" && l.Pass == "admin" {
		sqlExec()
		// }
	case "3":
		loginn()
		// fmt.Println("===Login===")
		// fmt.Print("Username :")
		// fmt.Scan(&l.User)
		// fmt.Print("Password :")
		// fmt.Scan(&l.Pass)
		// if l.User == "admin" && l.Pass == "admin" {
		sqlQuerydatabarang()
		// }
	case "4":
		loginn()
		// fmt.Println("===Login===")
		// fmt.Print("Username :")
		// fmt.Scan(&l.User)
		// fmt.Print("Password :")
		// fmt.Scan(&l.Pass)
		// if l.User == "admin" && l.Pass == "admin" {
		sqlQuerytransaksi()
		// }
	default:
		fmt.Print("salah")
	}
}
