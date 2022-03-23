package usecase

import (
	"WarungMakan/model"
	"WarungMakan/repository"
)

type UpdateStatusMeja interface {
	Update(nomerMeja, status string) model.Meja
}
type updateStatusMeja struct {
	repo repository.MejaRepo
}

func (a *updateStatusMeja) Update(nomerMeja, status string) model.Meja {
	newStatus := model.NewMeja(nomerMeja, status)
	a.repo.Update(newStatus)
	return newStatus
}
func NewUpdateStatusMejaUseCase(repo repository.MejaRepo) UpdateStatusMeja {
	return &updateStatusMeja{
		repo: repo,
	}
}
