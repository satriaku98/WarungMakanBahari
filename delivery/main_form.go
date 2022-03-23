package delivery

import (
	"fmt"
	"os"
	"strings"
)

func MainMenu() {
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println("\t\t\tWarung Makan Bahari")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println("|1. |Daftar Menu")
	fmt.Println("|2. |Halaman Pemesanan")
	fmt.Println("|3. |Halaman Pembayaran")
	fmt.Println("|4. |Keluar")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println("Pilih menu (1-3): ")
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
