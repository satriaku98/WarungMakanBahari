package usecase

import (
	"WarungMakan/model"
	"WarungMakan/repository"
)

type HapusPesananUseCase interface {
	Delete(nomerMeja, kodeMakanan, namaPelanggan string) model.Pesanan
}
type hapusPesananUseCase struct {
	repo repository.PesananRepo
}

func (a *hapusPesananUseCase) Delete(nomerMeja, kodeMakanan, namaPelanggan string) model.Pesanan {
	newDelete := model.NewPesanan(nomerMeja, kodeMakanan, namaPelanggan, 0)
	a.repo.Delete(newDelete)
	return newDelete
}
func NewHapusPesananUseCase(repo repository.PesananRepo) HapusPesananUseCase {
	return &hapusPesananUseCase{
		repo: repo,
	}
}
