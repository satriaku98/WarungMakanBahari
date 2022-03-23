package model

type Makanan struct {
	Code  string `db:"kode_makanan"`
	Name  string `db:"nama_makanan"`
	Harga int    `db:"harga"`
}

func (m *Makanan) SetMakananCode(code string) {
	m.Code = code
}
func (m *Makanan) SetMakananName(name string) {
	m.Name = name
}
func (m *Makanan) GetMakananCode() string {
	return m.Code
}
func (m *Makanan) GetMakananName() string {
	return m.Name
}

func (m *Makanan) SetHarga(harga int) {
	m.Harga = harga
}
func (m *Makanan) GetHarga() int {
	return m.Harga
}
