package delivery

import (
	"WarungMakan/config"
	"WarungMakan/usecase"
	"fmt"
	"strings"
)

func Pemesanan(useCasePesanan usecase.PesananMasukUseCase, useCaseMeja usecase.ListMejaUseCase, useCaseUpdateMeja usecase.UpdateStatusMeja) {
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println("\t\t\tWarung Makan Bahari")
	fmt.Println(strings.Repeat("=", 70))
	appConfig := config.NewConfig()

	for idx, meja := range useCaseMeja.GetAll() {
		fmt.Printf("%-3d%-20s%s\n", idx+1, meja.Nomer, meja.Status)
	}
	fmt.Println(strings.Repeat("=", 70))

	var nomer_meja string
	var kode_makanan string
	var nama_pelanggan string
	var saveProductConfirmation string
	var repeat string
	fmt.Println("Masukan Nomer Meja")
	fmt.Scanln(&nomer_meja)
	fmt.Println("Masukan Kode Makanan")
	fmt.Scanln(&kode_makanan)
	fmt.Println("Masukan Nama Pelanggan")
	fmt.Scanln(&nama_pelanggan)

	fmt.Printf("Nomer Meja : %s Kode Makanan : %s akan disimpan (Y/N) ?", nomer_meja, kode_makanan)
	fmt.Scanln(&saveProductConfirmation)

	if saveProductConfirmation == "Y" || saveProductConfirmation == "y" {
		useCasePesanan.Register(nomer_meja, kode_makanan, nama_pelanggan)
		fmt.Println("Apakah Ada Tambahan (Y/N) ?")
		fmt.Scanln(&repeat)
		if repeat == "Y" || repeat == "y" {
			Pemesanan(appConfig.UseCaseManager.PesananMasukUseCase(), appConfig.UseCaseManager.ListMejaUseCase(), appConfig.UseCaseManager.UpdateStatusMeja())
		}
		useCaseUpdateMeja.Update(nomer_meja, "")
	}
	ExitToMain()
}
