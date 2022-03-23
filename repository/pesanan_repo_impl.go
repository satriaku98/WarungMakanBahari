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
	//fmt.Println(err)
	if err != nil && err2 != nil {
		fmt.Println(err, err2)
	}
	tx.Commit()

	//tx2 := p.pesananDb.MustBegin() // memulai transaksi ke database yang aman
	//_, err2 := tx.Exec("update meja_makan set status='KOSONG' where nomer_meja = $1", newDelete.NomerMeja)
	//fmt.Println(newProduct.Code, newProduct.Name)
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	//tx2.Commit()
}
func (p *pesananRepoImpl) GetAll() []model.Pesanan {
	//return *p.productDb
	var dataProduct []model.Pesanan
	sql := `select * from pesanan  ORDER BY nomer_meja`
	p.pesananDb.Select(&dataProduct, sql) // Get dengan jumlah data yang banyak
	//fmt.Println(dataProduct)
	return dataProduct
}
func (p *pesananRepoImpl) Insert(newPesanan model.Pesanan) {
	//*p.productDb = append(*p.productDb, newProduct)

	tx := p.pesananDb.MustBegin() // memulai transaksi ke database yang aman
	_, err := tx.Exec("insert into pesanan(nomer_meja, kode_makanan, nama_pelanggan) values($1, $2, $3)", newPesanan.NomerMeja, newPesanan.KodeMakanan, newPesanan.NamaPelanggan)
	//fmt.Println(newProduct.Code, newProduct.Name)
	if err != nil {
		fmt.Println(err)
	}
	tx.Commit()
}

func NewPesananRepo(pesananDb *sqlx.DB) PesananRepo {
	pesananRepo := pesananRepoImpl{
		pesananDb: pesananDb,
	}
	return &pesananRepo
}
