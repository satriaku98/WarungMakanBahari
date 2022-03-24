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
	var dataProduct []model.Makanan
	sql := `select * from makanan`
	m.makananDb.Select(&dataProduct, sql) // Get dengan jumlah data yang banyak
	return dataProduct
}
func (m *makananRepoImpl) GetSelectedMakanan(kodeMakanan string) model.Makanan {
	var selectedProduct model.Makanan
	query := `select * from makanan where nama_makanan like $1 or kode_makanan like $1`
	m.makananDb.Get(&selectedProduct, query, fmt.Sprintf("%s%%", kodeMakanan)) // Get seperti first pada silverstripe
	return selectedProduct
}

func NewMakananRepo(makananDb *sqlx.DB) MakananRepo {
	makananRepo := makananRepoImpl{
		makananDb: makananDb,
	}
	return &makananRepo
}
