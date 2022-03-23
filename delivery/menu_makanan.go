package delivery

import (
	"WarungMakan/config"
	"WarungMakan/usecase"
	"fmt"
	"strings"
)

func MenuMakanan(useCase usecase.ListMakananUseCase) {
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println("\t\t\tWarung Makan Bahari")
	fmt.Println(strings.Repeat("=", 70))
	config.NewConfig()

	for idx, product := range useCase.GetAll() {
		fmt.Printf("%-3d%-20s%-20sRp. %-20d\n", idx+1, product.Code, product.Name, product.Harga)
	}
	fmt.Println(strings.Repeat("=", 70))
	ExitToMain()
}
