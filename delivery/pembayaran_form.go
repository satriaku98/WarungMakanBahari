package delivery

import (
	"WarungMakan/usecase"
	"fmt"
	"strings"
)

func Pembayaran(useCasePesanan usecase.ListPesananUseCase, useCaseHapus usecase.HapusPesananUseCase, useCaseTotal usecase.TotalHargaPesananUseCase) {
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println("\t\t\tWarung Makan Bahari")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Printf("%-10s%-20s%-20s%-20s%-20s\n", "Nomer", "NOMER MEJA", "KODE MAKANAN", "NAMA", "HARGA")
	//appConfig := config.NewConfig()
	for idx, meja := range useCasePesanan.GetAll() {
		fmt.Printf("%-10d%-20s%-20s%-20s%-20d\n", idx+1, meja.NomerMeja, meja.KodeMakanan, meja.NamaPelanggan, meja.Harga)
	}

	fmt.Println(strings.Repeat("=", 70))
	var nomer_meja string
	var saveProductConfirmation string
	fmt.Println("Masukan Nomer Meja")
	fmt.Scanln(&nomer_meja)
	total := useCaseTotal.Total(nomer_meja)
	fmt.Printf("Total Tagihan Nomer Meja %s Adalah = Rp. %d,--\n", nomer_meja, total)
	fmt.Printf("Nomer Meja : %s telah LUNAS (Y/N) ?", nomer_meja)
	fmt.Scanln(&saveProductConfirmation)
	if saveProductConfirmation == "Y" || saveProductConfirmation == "y" {
		//useCasePesanan.Register(nomer_meja, kode_makanan)
		//fmt.Println("Apakah Ada Tambahan (Y/N) ?")
		//fmt.Scanln(&repeat)
		//if repeat == "Y" || repeat == "y" {
		//	Pemesanan(appConfig.UseCaseManager.PesananMasukUseCase(), appConfig.UseCaseManager.ListMejaUseCase(), appConfig.UseCaseManager.UpdateStatusMeja())
		//}
		//useCaseUpdateMeja.Update(nomer_meja, "")
		useCaseHapus.Delete(nomer_meja, "", "")
		//fmt.Println("Oke")
	}
	ExitToMain()
}
