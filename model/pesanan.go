package model

type Pesanan struct {
	NomerMeja     string `db:"nomer_meja"`
	KodeMakanan   string `db:"kode_makanan"`
	NamaPelanggan string `db:"nama_pelanggan"`
}

func (m *Pesanan) SetPesananNomerMeja(nomerMeja string) {
	m.NomerMeja = nomerMeja
}
func (m *Pesanan) SetPesananKodeMekanan(kodeMakanan string) {
	m.KodeMakanan = kodeMakanan
}
func (m *Pesanan) SetPesananNamaPelanggan(namaPelanggan string) {
	m.NamaPelanggan = namaPelanggan
}
func (m *Pesanan) GetPesananNomerMeja() string {
	return m.NomerMeja
}
func (m *Pesanan) GetPesananKodeMakanan() string {
	return m.KodeMakanan
}
func (m *Pesanan) GetPesananNamaPelanggan() string {
	return m.KodeMakanan
}
func NewPesanan(nomerMeja, kodeMakanan, namaPelanggan string) Pesanan {
	return Pesanan{
		NomerMeja:     nomerMeja,
		KodeMakanan:   kodeMakanan,
		NamaPelanggan: namaPelanggan,
	}
}
