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

var jumHasilPertandingan int

var cekIdx int
var namaTimTour string
var jadwalTimTour string

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
		case 2:
			ubahDataTim()
		case 3:
			hapusDataTim()
		case 4:
			Klasemen()
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
	for *pilihOpsiMenuPanitia < 1 || *pilihOpsiMenuPanitia > 6 {
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
		fmt.Print("\t║ 5. Klasemen                                      ║\t\n")
		fmt.Print("\t║ 6. Halaman Utama                                 ║\t\n")
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
	switch pilihOpsiUbah {
	case 1:
		ubahNamaTim()
	case 2:
		ubahJadwalTim()
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
func ubahNamaTim() {

	fmt.Print("\n\t Masukkan Nama Tim yang Ingin Anda Ubah: ")
	fmt.Print(&namaTimTour)
	cekIdx = cariIdxTim(namaTimTour)
	if cekIdx != -1 {
		dataTim[cekIdx].namaTim = namaTimTour
	} else {
		fmt.Print("\n\t Nama Tim Tidak Berada Di Tournament")
	}
}
func ubahJadwalTim() {

	fmt.Print("\n\t Masukkan Jadwal Tim yang Ingin Anda Ubah: ")
	fmt.Print(&jadwalTimTour)
	cekIdx = cariIdxTim(jadwalTimTour)
	if cekIdx != -1 {
		dataTim[cekIdx].namaTim = jadwalTimTour
	} else {
		fmt.Print("\n\t Jadwal Tim Tidak Berada Di Tournament")
	}
}
func ubahHasilPertandingan() {

}

//========================================================================================================================================================================//

func hapusDataTim() {
	var pilihOpsiHapus int
	pilihHapus(&pilihOpsiHapus)
	switch pilihOpsiHapus {
	case 1:
		hapusNamaTim()
	case 2:
		hapusJadwalTim()
	}
}
func pilihHapus(pilihOpsiHapus *int) {
	fmt.Print("\n\t╔══════════════════════════════════════════════════╗\t\n")
	fmt.Print("\t║                                                  ║\t\n")
	fmt.Print("\t║        MASUKKAN DATA APA YANG INGIN DIHAPUS      ║\t\n")
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
	fmt.Scan(pilihOpsiHapus)
}
func hapusNamaTim() {
	var i int

	fmt.Print("\n\t Masukkan Nama Tim yang Ingin Anda Hapus: ")
	fmt.Print(&namaTimTour)
	cekIdx = cariIdxTim(namaTimTour)

	if cekIdx != -1 {
		for i = cekIdx; i < jumTimTour; i++ {
			dataTim[i].namaTim = dataTim[i+1].namaTim
		}
		jumTimTour--
	}
}
func hapusJadwalTim() {
	var i int

	fmt.Print("\n\t Masukkan Jadwal Tim yang Ingin Anda Hapus: ")
	fmt.Print(&jadwalTimTour)
	cekIdx = cariIdxTim(jadwalTimTour)

	if cekIdx != -1 {
		for i = cekIdx; i < jumTimTour; i++ {
			dataTim[i].jadwalTim = dataTim[i+1].jadwalTim
		}
		jumTimTour--
	}
}

//========================================================================================================================================================================//













//========================================================================================================================================================================//

func Klasemen() {
	var i int

	fmt.Print("\n\t╔═════════════╦══════╦══════╦═══════╗\t\n")
	fmt.Printf("\t║     %-8s║ %4s ║ %4s ║ %5s ║\n", "Team", "Win", "Lose", "Score")
	fmt.Println("\t╠═════════════╬══════╬══════╬═══════╣")

	for i = 0; i < jumTimTour; i++ {
		fmt.Printf("║     %-8s║ %4d ║ %4d ║ %5d ║\n",
			dataTim[i].namaTim, dataTim[i].win, dataTim[i].lose, dataTim[i].skorTim)
	}
	fmt.Println("\t╚═════════════╩══════╩══════╩═══════╝")
}
func selectionSort() {
	var pass,idx,i,temp int

	pass = 1
	for pass < jumTimTour{
		idx = pass - 1
		i = pass

		for i < jumTimTour + 1{
			if dataTim[idx].skorTim < dataTim[idx].skorTim {
				idx = i
			}
			i++
		}
		temp = dataTim[pass-1].skorTim
		dataTim[pass - 1].skorTim = dataTim[idx].skorTim
		dataTim[idx].skorTim = temp
		pass++
	}
}
func insertSort(){
	var pass,i,temp int

	pass = 1
	for pass <= jumTimTour{
		i = pass
		temp = dataTim[pass].skorTim
		for i > 0 && temp < dataTim[i-1].skorTim{
			dataTim[i].skorTim = dataTim[i-1].skorTim
			i--
		}
		dataTim[i].skorTim = temp
		pass++
	}
}
