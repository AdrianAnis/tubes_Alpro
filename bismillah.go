package main

import "fmt"

const NMAX int = 100000

type dataTimTournament struct {
	namaTim   string
	jadwalTim string
	win       int
	lose      int
	skorTim   int
}
type dataTour [NMAX]dataTimTournament

var jumTimTour int
var dataTim dataTour

var timA, timB, JadwalMatch string
var skorA, skorB int
var namaTim int

var jumHasilPertandingan int

//========================================================================================================================================================================//

func main() {
	var pilihOpsiPeranFunc int
	var pilihOpsiMenuPanitia int
	var pilihOpsiMenuPeserta int
	var n int

	MenuUtama(&pilihOpsiPeranFunc)
	if pilihOpsiPeranFunc == 1 {
		menuPanitia(&pilihOpsiMenuPanitia)
		switch pilihOpsiMenuPanitia {
		case 1:
			nambahJumTim(n)
			masukanHasilPertandingan()
		}
	} else if pilihOpsiMenuPanitia == 2 {
		menuPeserta(&pilihOpsiMenuPeserta)
	}
}

//========================================================================================================================================================================//

func MenuUtama(pilihOpsiPeranFunc *int) {
	for *pilihOpsiPeranFunc != 1 && *pilihOpsiPeranFunc != 2 {
		fmt.Print("\n\t╔══════════════════════════════════════════════════╗\t\n")
		fmt.Print("\t║                                                  ║\t\n")
		fmt.Print("\t║   SELAMAT DATANG DI E-SPORTS TOURNAMENT TELKIF!  ║\t\n")
		fmt.Print("\t║                  Dibuat Oleh:                    ║\t\n")
		fmt.Print("\t║        Adrian Anis Pratama (103012400098)        ║\t\n")
		fmt.Print("\t║        Agiel Ghazi Abizar (103012400037)         ║\t\n")
		fmt.Print("\t║                                                  ║\t\n")
		fmt.Print("\t║==================================================║\t\n")
		fmt.Print("\t║ Silahkan Pilih Peran Anda Dibawah                ║\t\n")
		fmt.Print("\t║                                                  ║\t\n")
		fmt.Print("\t║ 1. Panitia Tournament                            ║\t\n")
		fmt.Print("\t║ 2. Pemain/Peserta Tournament                     ║\t\n")
		fmt.Print("\t║                                    ╔═════════════║\t\n")
		fmt.Print("\t║                                    ║HALAMAN UTAMA║\t\n")
		fmt.Print("\t╚════════════════════════════════════╩═════════════╝\t\n")
		fmt.Print("\tMasukkan Pilihan Anda: ")
		fmt.Scan(pilihOpsiPeranFunc)
		fmt.Println()

		if *pilihOpsiPeranFunc != 1 && *pilihOpsiPeranFunc != 2 {
			fmt.Print("\n\tPilihan Tidak Valid, Coba lagi!\n\n")
		}
	}
}

func menuPanitia(pilihOpsiMenuPanitia *int) {
	for *pilihOpsiMenuPanitia != 1 && *pilihOpsiMenuPanitia != 2 {
		fmt.Print("\n\t╔══════════════════════════════════════════════════╗\t\n")
		fmt.Print("\t║                                                  ║\t\n")
		fmt.Print("\t║          SELAMAT DATANG SEBAGAI PANITIA          ║\t\n")
		fmt.Print("\t║          DI E-SPORTS TOURNAMENT TELKIF!          ║\t\n")
		fmt.Print("\t║                                                  ║\t\n")
		fmt.Print("\t║==================================================║\t\n")
		fmt.Print("\t║ Silahkan Pilih Opsi Yang Ada Dibawah             ║\t\n")
		fmt.Print("\t║                                                  ║\t\n")
		fmt.Print("\t║ 1. Menambah Data TIM                             ║\t\n")
		fmt.Print("\t║ 2. Mengubah/Mengedit Data TIM                    ║\t\n")
		fmt.Print("\t║ 3. Menghapus Data TIM                            ║\t\n")
		fmt.Print("\t║ 4. Mencari TIM                                   ║\t\n")
		fmt.Print("\t║ 5. Mengurutkan TIM (Berdasarkan Poin )           ║\t\n")
		fmt.Print("\t║ 6. Klasemen                                      ║\t\n")
		fmt.Print("\t║ 7. Halaman Utama                                 ║\t\n")
		fmt.Print("\t║                                                  ║\t\n")
		fmt.Print("\t╚══════════════════════════════════════════════════╝\t\n")
		fmt.Print("\tMasukkan Pilihan Anda: ")
		fmt.Scan(pilihOpsiMenuPanitia)
		fmt.Println()

		if *pilihOpsiMenuPanitia < 1 && *pilihOpsiMenuPanitia > 6 {
			fmt.Print("\n\tPilihan Tidak Valid, Coba lagi!\n\n")
		}
	}
}
func menuPeserta(pilihOpsiMenuPeserta *int) {
	for *pilihOpsiMenuPeserta != 1 && *pilihOpsiMenuPeserta != 2 {
		fmt.Print("\n\t╔══════════════════════════════════════════════════╗\t\n")
		fmt.Print("\t║                                                  ║\t\n")
		fmt.Print("\t║          SELAMAT DATANG SEBAGAI PANITIA          ║\t\n")
		fmt.Print("\t║          DI E-SPORTS TOURNAMENT TELKIF!          ║\t\n")
		fmt.Print("\t║                                                  ║\t\n")
		fmt.Print("\t║==================================================║\t\n")
		fmt.Print("\t║ Silahkan Pilih Opsi Yang Ada Dibawah             ║\t\n")
		fmt.Print("\t║                                                  ║\t\n")
		fmt.Print("\t║ 1. Mencari TIM                                   ║\t\n")
		fmt.Print("\t║ 2. Mengurutkan TIM                               ║\t\n")
		fmt.Print("\t║ 3. Halaman Utama                                 ║\t\n")
		fmt.Print("\t║                                                  ║\t\n")
		fmt.Print("\t╚══════════════════════════════════════════════════╝\t\n")
		fmt.Print("\tMasukkan Pilihan Anda: ")
		fmt.Scan(pilihOpsiMenuPeserta)
		fmt.Println()

		if *pilihOpsiMenuPeserta < 1 && *pilihOpsiMenuPeserta > 3 {
			fmt.Print("\n\tPilihan Tidak Valid, Coba lagi!\n\n")
		}
	}
}

//========================================================================================================================================================================//
func nambahJumTim(n int) {

	fmt.Print("\n\t╔══════════════════════════════════════════════════╗\t\n")
	fmt.Print("\t║      Berapa total TIM yang akan Ditambahkan?     ║\t\n")
	fmt.Print("\t╚══════════════════════════════════════════════════╝\t\n")
	fmt.Print("\tMasukkan Pilihan Anda: ")
	fmt.Scan(&n)
	masukanTimTanggal(n)
}
func masukanTimTanggal(n int) {
	var i int

	fmt.Print("\n\t╔══════════════════════════════════════════════════╗\t\n")
	fmt.Print("\t║    Masukkan Nama Tim yang akan berpartisipasi    ║\t\n")
	fmt.Print("\t║    dan Jadwal pertandingan Dari TIM tersebut     ║\t\n")
	fmt.Print("\t╚══════════════════════════════════════════════════╝\t\n")

	for i = 1; i <= n; i++ {
		fmt.Printf("\tMasukkan nama tim ke-%d: ", i)
		fmt.Scan(&dataTim[jumTimTour].namaTim)

		fmt.Printf("\tMasukkan jadwal pertandingan tim %s: ", dataTim[jumTimTour].namaTim)
		fmt.Scan(&dataTim[jumTimTour].jadwalTim)
		jumTimTour++
	}
}
func masukanPertandingan() {
	fmt.Print("\n\t╔══════════════════════════════════════════════════╗\t\n")
	fmt.Print("\t║      Masukkan Nama Tim yang akan bertanding      ║\t\n")
	fmt.Print("\t║         dan Jadwal pertandingan Hari ini         ║\t\n")
	fmt.Print("\t╚══════════════════════════════════════════════════╝\t\n")

	fmt.Print("\t TIM A: ")
	fmt.Scan(&timA)
	fmt.Print("\t TIM B: ")
	fmt.Scan(&timB)
	fmt.Print("\t Jadwal Pertandingan Hari ini: ")
	fmt.Scan(&JadwalMatch)
}
func masukanHasilPertandingan() {
	var i int
	var ditemukanA, ditemukanB bool
	var idxTimA, idxTimB int

	masukanPertandingan()

	for i = 0; i < jumTimTour; i++ {
		if dataTim[i].namaTim == timA && dataTim[i].jadwalTim == JadwalMatch {
			ditemukanA = true
		}
		if dataTim[i].namaTim == timB && dataTim[i].jadwalTim == JadwalMatch {
			ditemukanB = true
		}
	}
	if ditemukanA && ditemukanB {
		idxTimA = cariIdxTim(timA)
		idxTimB = cariIdxTim(timB)

		masukanSkor()

		dataTim[idxTimA].win = skorA
		dataTim[idxTimA].lose = skorB

		dataTim[idxTimB].lose = skorA
		dataTim[idxTimB].win = skorB

		if skorA > skorB {
			dataTim[idxTimA].skorTim = dataTim[idxTimA].skorTim + 3
		} else if skorB > skorA {
			dataTim[idxTimB].skorTim = dataTim[idxTimA].skorTim + 3
		}
	} else if ditemukanA && !ditemukanB {
		fmt.Println("\n\tTim B tidak ditemukan atau jadwalnya salah.")
	} else if !ditemukanA && ditemukanB {
		fmt.Println("\n\tTim A tidak ditemukan atau jadwalnya salah.")
	} else {
		fmt.Println("\n\tKedua tim tidak ditemukan atau jadwalnya salah.")
	}
}
func cariIdxTim(tim string) int {
	var ketemu int
	var k int

	ketemu = -1
	k = 0
	for ketemu == -1 && k < jumTimTour {
		if dataTim[k].namaTim == tim {
			ketemu = k
		}
		k++
	}
	return ketemu
}
func masukanSkor() {
	fmt.Printf("\t Masukkan Skor Tim %s: ", timA)
	fmt.Scan(&skorA)
	fmt.Printf("\t Masukkan Skor Tim %s: ", timB)
	fmt.Scan(&skorB)
}

//========================================================================================================================================================================//

func ubahDataTim() {
	var pilihOpsiUbah int
	pilihUbah(&pilihOpsiUbah)
	switch pilihOpsiUbah{
	case 1:

	}
}

func pilihUbah(pilihOpsiUbah *int) {
	fmt.Print("\n\t╔══════════════════════════════════════════════════╗\t\n")
	fmt.Print("\t║                                                  ║\t\n")
	fmt.Print("\t║         MASUKKAN DATA APA YANG INGIN DIUBAH      ║\t\n")
	fmt.Print("\t║                                                  ║\t\n")
	fmt.Print("\t║==================================================║\t\n")
	fmt.Print("\t║ Silahkan Pilih Opsi Yang Ada Dibawah             ║\t\n")
	fmt.Print("\t║                                                  ║\t\n")
	fmt.Print("\t║ 1. Nama TIM                                      ║\t\n")
	fmt.Print("\t║ 2. Jadwal TIM                                    ║\t\n")
	fmt.Print("\t║ 3. Hasil Pertandingan                            ║\t\n")
	fmt.Print("\t║                                                  ║\t\n")
	fmt.Print("\t╚══════════════════════════════════════════════════╝\t\n")
	fmt.Print("\tMasukkan Pilihan Anda: ")
	fmt.Scan(pilihOpsiUbah)
}
func cekUbahData(){
	fmt.Print("\n\t Masukkan Nama Tim yang Ingin Anda Ubah: ")
	fmt.Print(namaTim)
}
