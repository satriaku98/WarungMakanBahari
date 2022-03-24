package repository

import (
	"WarungMakan/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type pesananRepoImpl struct {
	pesananDb *sqlx.DB
}

func (p *pesananRepoImpl) Delete(newDelete model.Pesanan) {
	tx := p.pesananDb.MustBegin() // memulai transaksi ke database yang aman
	_, err := tx.Exec("delete from pesanan where nomer_meja=$1", newDelete.NomerMeja)
	_, err2 := tx.Exec("update meja_makan set status='KOSONG' where nomer_meja = $1", newDelete.NomerMeja)
	if err != nil && err2 != nil {
		fmt.Println(err, err2)
	}
	tx.Commit()
}
func (p *pesananRepoImpl) GetAll() []model.Pesanan {
	//return *p.productDb
	var dataProduct []model.Pesanan
	sql := `select p.nomer_meja,p.kode_makanan,p.nama_pelanggan,m.harga from pesanan p join makanan m on p.kode_makanan=m.kode_makanan`
	p.pesananDb.Select(&dataProduct, sql) // Get dengan jumlah data yang banyak
	//fmt.Println(dataProduct)
	return dataProduct
}
func (p *pesananRepoImpl) Insert(newPesanan model.Pesanan) {
	//*p.productDb = append(*p.productDb, newProduct)
	var status string
	sqlmd := `select status from meja_makan where nomer_meja = $1`
	p.pesananDb.Get(&status, sqlmd, newPesanan.NomerMeja)
	//fmt.Println(status)
	if status == "KOSONG" {
		tx := p.pesananDb.MustBegin() // memulai transaksi ke database yang aman
		_, err := tx.Exec("insert into pesanan(nomer_meja, kode_makanan, nama_pelanggan) values($1, $2, $3)", newPesanan.NomerMeja, newPesanan.KodeMakanan, newPesanan.NamaPelanggan)
		//fmt.Println(newProduct.Code, newProduct.Name)
		if err != nil {
			fmt.Println(err)
		}
		tx.Commit()
	} else if status == "ISI" {
		fmt.Println("Meja Telah Terisi")
	}

}
func (p *pesananRepoImpl) Total(meja string) int {
	//*p.productDb = append(*p.productDb, newProduct)
	var harga int
	//fmt.Println(meja)
	sqlmd := `select SUM(m.harga) from pesanan p join makanan m  on p.kode_makanan=m.kode_makanan where p.nomer_meja = $1`
	p.pesananDb.Get(&harga, sqlmd, meja)
	//fmt.Println(harga)
	return harga
}
func NewPesananRepo(pesananDb *sqlx.DB) PesananRepo {
	pesananRepo := pesananRepoImpl{
		pesananDb: pesananDb,
	}
	return &pesananRepo
}
