package pegawai

import "fmt"

type Pegawai struct {
	Nama   string
	Posisi	string
	GajiBulanan   float64
}

type InformasiPegawai interface {
	TampilkanInformasi()
}

func (p Pegawai) HitungGajiTahunan() float64 {
	return p.GajiBulanan * 12
}

func (p Pegawai) TampilkanInformasi() {
	fmt.Println("Nama:", p.Nama)
	fmt.Println("Posisi:", p.Posisi)
	fmt.Printf("Gaji Tahunan: %.2f\n", p.HitungGajiTahunan())
}