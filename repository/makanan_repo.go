package repository

import "WarungMakan/model"

type MakananRepo interface {
	GetAll() []model.Makanan
}
