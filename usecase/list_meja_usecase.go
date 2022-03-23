package usecase

import (
	"WarungMakan/model"
	"WarungMakan/repository"
)

type ListMejaUseCase interface {
	GetAll() []model.Meja
}
type listMejaUseCase struct {
	repo repository.MejaRepo
}

func (a *listMejaUseCase) GetAll() []model.Meja {
	return a.repo.GetAll()
}

func NewListMejaUseCase(repo repository.MejaRepo) ListMejaUseCase {
	return &listMejaUseCase{
		repo: repo,
	}
}
