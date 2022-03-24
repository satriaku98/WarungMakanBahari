package util

import (
	"fmt"
	"strings"
)

const (
	//TabelPembayaranHeader = "Nomer", "NOMER MEJA", "KODE MAKANAN", "NAMA", "HARGA"
	ListTableFormat = "%-10s%-20s%-20s%-20s%-20s\n"
)

func GarisAtas() {
	fmt.Println(strings.Repeat("=", 90))
}
