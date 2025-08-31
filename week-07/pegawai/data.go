package pegawai

import "fmt"

type Pegawai struct {
	ID          uint `gorm:"primaryKey"`
	Nama        string
	Posisi      string
	GajiBulanan float64
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
	fmt.Printf("Gaji Bulanan: %.2f\n", p.GajiBulanan)
	fmt.Printf("Gaji Tahunan: %.2f\n", p.HitungGajiTahunan())
	fmt.Println("+--------------------------------+")
}
