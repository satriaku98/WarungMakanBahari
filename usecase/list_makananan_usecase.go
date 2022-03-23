package usecase

import (
	"WarungMakan/model"
	"WarungMakan/repository"
)

type ListMakananUseCase interface {
	GetAll() []model.Makanan
}
type listMakananUseCase struct {
	repo repository.MakananRepo
}

func (a *listMakananUseCase) GetAll() []model.Makanan {
	return a.repo.GetAll()
}

func NewListMakananUseCase(repo repository.MakananRepo) ListMakananUseCase {
	return &listMakananUseCase{
		repo: repo,
	}
}
