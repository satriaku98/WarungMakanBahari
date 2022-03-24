package usecase

import (
	"WarungMakan/repository"
)

type TotalHargaPesananUseCase interface {
	Total(meja string) int
}

type totalHargaPesananUseCase struct {
	repo repository.PesananRepo
}

func (a *totalHargaPesananUseCase) Total(meja string) int {
	total := a.repo.Total(meja)
	return total
}
func NewTotalHargaPesananUseCase(repo repository.PesananRepo) TotalHargaPesananUseCase {
	return &totalHargaPesananUseCase{
		repo: repo,
	}
}
