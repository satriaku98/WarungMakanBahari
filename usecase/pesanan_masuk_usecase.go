package usecase

import (
	"WarungMakan/model"
	"WarungMakan/repository"
)

type PesananMasukUseCase interface {
	Register(nomerMeja, kodeMakanan, namaPelanggan string) model.Pesanan
}

type pesanananMasukUseCase struct {
	repo repository.PesananRepo
}

func (a *pesanananMasukUseCase) Register(nomerMeja, kodeMakanan, namaPelanggan string) model.Pesanan {
	newProduct := model.NewPesanan(nomerMeja, kodeMakanan, namaPelanggan)
	a.repo.Insert(newProduct)
	return newProduct
}

func NewPesananRegistrationUseCase(repo repository.PesananRepo) PesananMasukUseCase {
	return &pesanananMasukUseCase{
		repo: repo,
	}
}
