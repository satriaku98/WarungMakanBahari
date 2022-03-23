package main

import (
	"WarungMakan/config"
	"WarungMakan/delivery"
	"fmt"
)

func main() {
	appConfig := config.NewConfig()
	delivery.MainMenu()
	for {
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			delivery.MenuMakanan(appConfig.UseCaseManager.ListMakananUseCase(), appConfig.UseCaseManager.GetSelectedMakanan())
			break
		case "2":
			delivery.Pemesanan(appConfig.UseCaseManager.PesananMasukUseCase(), appConfig.UseCaseManager.ListMejaUseCase(), appConfig.UseCaseManager.UpdateStatusMeja())
			break
		case "3":
			delivery.Pembayaran(appConfig.UseCaseManager.ListPesananUseCase(), appConfig.UseCaseManager.HapusPesananUseCase())
			break
		case "4":
			delivery.ExitApp()
		}
	}
}
