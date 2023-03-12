package main

type kendaraan struct {
	totalRoda       int
	kecepatanPerJam int
}

type mobil struct {
	kendaraan
}

// tambahkan pointer pada parameter fungsi untuk mengubah struct kendaraan secara langsung
func (m *mobil) berjalan() {
	m.tambahKecepatan(10)
}

// tambahkan nama field pada parameter untuk lebih jelas
func (m *mobil) tambahKecepatan(kecepatanBaru int) {
	m.kecepatanPerJam += kecepatanBaru
}

func main() {
	mobilCepat := mobil{}
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()

	mobilLambat := mobil{}
	mobilLambat.berjalan()
}
