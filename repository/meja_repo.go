package repository

import "WarungMakan/model"

type MejaRepo interface {
	GetAll() []model.Meja
	Update(newStatus model.Meja)
}
