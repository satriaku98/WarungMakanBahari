package manager

import "WarungMakan/usecase"

type UseCaseManager interface {
	ListMakananUseCase() usecase.ListMakananUseCase
	ListMejaUseCase() usecase.ListMejaUseCase
	PesananMasukUseCase() usecase.PesananMasukUseCase
	UpdateStatusMeja() usecase.UpdateStatusMeja
	ListPesananUseCase() usecase.ListPesananUseCase
	HapusPesananUseCase() usecase.HapusPesananUseCase
	GetSelectedMakanan() usecase.SearchMakananUseCase
}
type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) ListMakananUseCase() usecase.ListMakananUseCase {
	return usecase.NewListMakananUseCase(u.repo.MakananRepo())
}
func (u *useCaseManager) ListMejaUseCase() usecase.ListMejaUseCase {
	return usecase.NewListMejaUseCase(u.repo.MejaRepo())
}
func (u *useCaseManager) PesananMasukUseCase() usecase.PesananMasukUseCase {
	return usecase.NewPesananRegistrationUseCase(u.repo.PesananRepo())
}
func (u *useCaseManager) UpdateStatusMeja() usecase.UpdateStatusMeja {
	return usecase.NewUpdateStatusMejaUseCase(u.repo.MejaRepo())
}
func (u *useCaseManager) ListPesananUseCase() usecase.ListPesananUseCase {
	return usecase.NewListPesananUseCase(u.repo.PesananRepo())
}
func (u *useCaseManager) HapusPesananUseCase() usecase.HapusPesananUseCase {
	return usecase.NewHapusPesananUseCase(u.repo.PesananRepo())
}
func (u *useCaseManager) GetSelectedMakanan() usecase.SearchMakananUseCase {
	return usecase.NewSearchMakananUseCase(u.repo.MakananRepo())
}
func NewUseCaseManager(manager RepoManager) UseCaseManager {
	return &useCaseManager{
		repo: manager,
	}
}
