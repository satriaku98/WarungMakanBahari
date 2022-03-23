package usecase

import (
	"WarungMakan/model"
	"WarungMakan/repository"
)

type SearchMakananUseCase interface {
	Search(productCode string) []model.Makanan
}
type searchMakananUseCase struct {
	repo repository.MakananRepo
}

func (a *searchMakananUseCase) Search(productCode string) []model.Makanan {
	if len(productCode) == 0 {
		return a.repo.GetAll()
	}
	result := a.repo.GetSelectedMakanan(productCode)
	if len(result.GetMakananCode()) == 0 {
		return nil
	} else {
		return []model.Makanan{result}
	}
}
func NewSearchMakananUseCase(repo repository.MakananRepo) SearchMakananUseCase {
	return &searchMakananUseCase{
		repo: repo,
	}
}
