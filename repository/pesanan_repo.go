package repository

import "WarungMakan/model"

type PesananRepo interface {
	Insert(newPesanan model.Pesanan)
	Delete(newDelete model.Pesanan)
	GetAll() []model.Pesanan
}
