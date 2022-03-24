package delivery

import (
	"WarungMakan/config"
	"WarungMakan/delivery/util"
	"WarungMakan/usecase"
	"fmt"
)

func MenuMakanan(useCaseList usecase.ListMakananUseCase, useCaseSearch usecase.SearchMakananUseCase) {
	util.GarisAtas()
	fmt.Println("\t\t\t\tWarung Makan Bahari")
	util.GarisAtas()
	config.NewConfig()
	for idx, product := range useCaseList.GetAll() {
		fmt.Printf("%-3d%-20s%-20sRp. %-20d\n", idx+1, product.Code, product.Name, product.Harga)
	}
	util.GarisAtas()
	var kodeMakanan string
	fmt.Print("Masukan Pencarian (Kode atau Nama Makanan) = ")
	fmt.Scanln(&kodeMakanan)
	makanan := useCaseSearch.Search(kodeMakanan)
	if makanan != nil {
		p := makanan[0]
		fmt.Printf("%-3d%-20s%-20s\n", 1, p.GetMakananCode(), p.GetMakananName())
	} else {
		fmt.Println("Produk tidak ditemukan")
	}
	ExitToMain()
}
