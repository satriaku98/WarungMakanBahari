package model

type Meja struct {
	Nomer  string `db:"nomer_meja"`
	Status string `db:"status"`
}

func (m *Meja) SetMejaCode(code string) {
	m.Nomer = code
}
func (m *Meja) SetMejaName(name string) {
	m.Status = name
}
func (m *Meja) GetMejaCode() string {
	return m.Nomer
}
func (m *Meja) GetMejaName() string {
	return m.Status
}
func NewMeja(nomerMeja, status string) Meja {
	return Meja{
		Nomer:  nomerMeja,
		Status: status,
	}
}
