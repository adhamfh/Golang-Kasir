package views

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

// func main() {
func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "form.html")
		return
	}
	var Barang barang

	Barang.ID = r.FormValue("ID")
	Barang.NamaBarang = r.FormValue("NamaBarang")
	Barang.Harga, _ = strconv.ParseInt(r.FormValue("Harga"), 0, 64)
	Barang.Jumlah = r.FormValue("Jumlah")
	Barang.Total = r.FormValue("Total")
	Barang.Bayar = r.FormValue("Bayar")
	Barang.Kembalian = r.FormValue("Kembalian")

	b := &barang{}

	simpanbarang(b, Barang)
}
func (b *barang) simpanbarang(Barang barang) barang {
	db := infastruktu.connect()
	b.ID = Barang.ID
	b.NamaBarang = Barang.NamaBarang
	b.Harga = Barang.Harga
	b.Jumlah = Barang.Jumlah
	b.Total = Barang.Total
	b.Bayar = Barang.Bayar
	b.Kembalian = Barang.Kembalian
	t := time.Now()
	_, err := db.Exec("INSERT INTO simpanbarang(simpanbarang(ID,NamaBarang, Harga,Jumlah,Total,Tanggal,Bayar,Kembalian)values (?,?,?,?,?,?,?,?)", b.ID, b.NamaBarang, b.Harga, b.Jumlah, b.Total, t, b.Bayar, b.Kembalian)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	return *b
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Selamat Datand Di website saya gans!")
	})
	http.HandleFunc("/form", Form)
	fmt.Println("Server running on port :8000")

	http.ListenAndServe(":8000", nil)
}

// http.HandleFunc("/", handlerIndex)
// http.HandleFunc("/index", handlerIndex)
