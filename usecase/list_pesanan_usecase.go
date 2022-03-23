package usecase

import (
	"WarungMakan/model"
	"WarungMakan/repository"
)

type ListPesananUseCase interface {
	GetAll() []model.Pesanan
}
type listPesananUseCase struct {
	repo repository.PesananRepo
}

func (a *listPesananUseCase) GetAll() []model.Pesanan {
	return a.repo.GetAll()
}

func NewListPesananUseCase(repo repository.PesananRepo) ListPesananUseCase {
	return &listPesananUseCase{
		repo: repo,
	}
}
