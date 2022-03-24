package repository

import (
	"WarungMakan/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type makananRepoImpl struct {
	makananDb *sqlx.DB
}

func (m *makananRepoImpl) GetAll() []model.Makanan {
	//return *p.productDb
	var dataProduct []model.Makanan
	sql := `select * from makanan`
	m.makananDb.Select(&dataProduct, sql) // Get dengan jumlah data yang banyak
	//fmt.Println(dataProduct)
	return dataProduct
}
func (m *makananRepoImpl) GetSelectedMakanan(kodeMakanan string) model.Makanan {
	var selectedProduct model.Makanan
	//m.makananDb.Get(&selectedProduct, "select * from makanan where kode_makanan = $1", kodeMakanan) // Get seperti first pada silverstripe
	//input := fmt.Println("'%%s%%'" + kodeMakanan)
	//query := fmt.Sprintf("")
	query := `select * from makanan where nama_makanan like $1 or kode_makanan like $1`
	m.makananDb.Get(&selectedProduct, query, fmt.Sprintf("%s%%", kodeMakanan)) // Get seperti first pada silverstripe
	//fmt.Println(selectedProduct)
	//fmt.Println(kodeMakanan)
	return selectedProduct
}

func NewMakananRepo(makananDb *sqlx.DB) MakananRepo {
	makananRepo := makananRepoImpl{
		makananDb: makananDb,
	}
	return &makananRepo
}
