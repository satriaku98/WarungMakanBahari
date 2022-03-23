package repository

import (
	"WarungMakan/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type mejaRepoImpl struct {
	mejaDb *sqlx.DB
}

func (m *mejaRepoImpl) GetAll() []model.Meja {
	//return *p.productDb
	var dataProduct []model.Meja
	sql := `select * from meja_makan where status!='ISI' order by nomer_meja`
	m.mejaDb.Select(&dataProduct, sql) // Get dengan jumlah data yang banyak
	//fmt.Println(dataProduct)
	return dataProduct
}

func (m *mejaRepoImpl) Update(newMeja model.Meja) {
	//return *p.productDb
	tx := m.mejaDb.MustBegin() // memulai transaksi ke database yang aman
	_, err := tx.Exec("update meja_makan set status='ISI' where nomer_meja = $1", newMeja.Nomer)
	//fmt.Println(newProduct.Code, newProduct.Name)
	if err != nil {
		fmt.Println(err)
	}
	tx.Commit()
}

func NewMejaRepo(mejaDb *sqlx.DB) MejaRepo {
	mejaRepo := mejaRepoImpl{
		mejaDb: mejaDb,
	}
	return &mejaRepo
}
