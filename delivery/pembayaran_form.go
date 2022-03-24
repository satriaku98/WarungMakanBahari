package delivery

import (
	"WarungMakan/delivery/util"
	"WarungMakan/usecase"
	"fmt"
	"strings"
)

func Pembayaran(useCasePesanan usecase.ListPesananUseCase, useCaseHapus usecase.HapusPesananUseCase, useCaseTotal usecase.TotalHargaPesananUseCase) {
	util.GarisAtas()
	fmt.Println("\t\t\t\tWarung Makan Bahari")
	util.GarisAtas()
	fmt.Println("\t\t\t\tFORM PEMBAYARAN")
	util.GarisAtas()
	fmt.Printf(util.ListTableFormat, "Nomer", "NOMER MEJA", "KODE MAKANAN", "NAMA", "HARGA")
	fmt.Println(strings.Repeat("-", 90))
	for idx, meja := range useCasePesanan.GetAll() {
		fmt.Printf("%-10d%-20s%-20s%-20s%-20d\n", idx+1, meja.NomerMeja, meja.KodeMakanan, meja.NamaPelanggan, meja.Harga)
	}

	util.GarisAtas()
	var nomer_meja string
	var uang int
	var kembalian int
	var saveProductConfirmation string
	fmt.Println("Masukan Nomer Meja")
	fmt.Scanln(&nomer_meja)
	total := useCaseTotal.Total(nomer_meja)
	fmt.Printf("Total Tagihan Nomer Meja %s Adalah = Rp. %d,--\n", nomer_meja, total)
	fmt.Print("Uang yang diberikan = ")
	fmt.Scanln(&uang)
	kembalian = uang - total
	fmt.Printf("Total Kembalian = %d\n", kembalian)
	fmt.Printf("Nomer Meja : %s telah LUNAS (Y/N) ?", nomer_meja)
	fmt.Scanln(&saveProductConfirmation)
	if saveProductConfirmation == "Y" || saveProductConfirmation == "y" {
		useCaseHapus.Delete(nomer_meja, "", "")
	}
	ExitToMain()
}
