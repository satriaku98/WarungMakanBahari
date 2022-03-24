package delivery

import (
	"WarungMakan/delivery/util"
	"fmt"
	"os"
)

func MainMenu() {
	util.GarisAtas()
	fmt.Println("\t\t\tWarung Makan Bahari")
	util.GarisAtas()
	fmt.Println("|1. |Daftar Menu")
	fmt.Println("|2. |Halaman Pemesanan")
	fmt.Println("|3. |Halaman Pembayaran")
	fmt.Println("|4. |Keluar")
	util.GarisAtas()
	fmt.Println("Pilih menu (1-4): ")
}
func ExitToMain() {
	var mainMenuConfirmation string
	fmt.Printf("Apakah Anda ingin kembali ke Main Menu (y/n) ?")
	fmt.Scanln(&mainMenuConfirmation)
	if mainMenuConfirmation == "y" {
		MainMenu()
	} else if mainMenuConfirmation == "n" {
		ExitApp()
	} else {
		fmt.Println("Inputan Salah ! Kembali Ke Main Menu !")
		MainMenu()
	}
}

func ExitApp() {
	fmt.Println("Warung Makan Bahari JAYA JAYA JAYA")
	os.Exit(0)
}
