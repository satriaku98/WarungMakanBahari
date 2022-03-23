package delivery

import (
	"WarungMakan/config"
	"WarungMakan/usecase"
	"fmt"
	"strings"
)

func MenuMakanan(useCaseList usecase.ListMakananUseCase, useCaseSearch usecase.SearchMakananUseCase) {
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println("\t\t\tWarung Makan Bahari")
	fmt.Println(strings.Repeat("=", 70))
	config.NewConfig()

	for idx, product := range useCaseList.GetAll() {
		fmt.Printf("%-3d%-20s%-20sRp. %-20d\n", idx+1, product.Code, product.Name, product.Harga)
	}
	fmt.Println(strings.Repeat("=", 70))
	var kodeMakanan string
	fmt.Print("Masukan Kode Makanan = ")
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
