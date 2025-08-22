package main

import "hsiweek05/pegawai"

func main() {
	dataPegawai := pegawai.Pegawai{
		Nama: "Fulan bin Fulan",
		Posisi: "Go Developer",
		GajiBulanan: 8000000,
	}

	dataPegawai.TampilkanInformasi()
}