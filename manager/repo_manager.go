package manager

import (
	"WarungMakan/repository"
)

type RepoManager interface {
	MakananRepo() repository.MakananRepo
	MejaRepo() repository.MejaRepo
	PesananRepo() repository.PesananRepo
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) MakananRepo() repository.MakananRepo {
	return repository.NewMakananRepo(r.infra.SqlDb())
}

func (r *repoManager) MejaRepo() repository.MejaRepo {
	return repository.NewMejaRepo(r.infra.SqlDb())
}
func (r *repoManager) PesananRepo() repository.PesananRepo {
	return repository.NewPesananRepo(r.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
