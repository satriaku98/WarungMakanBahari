package repository

import (
	"WarungMakan/model"
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

func NewMakananRepo(makananDb *sqlx.DB) MakananRepo {
	makananRepo := makananRepoImpl{
		makananDb: makananDb,
	}
	return &makananRepo
}
