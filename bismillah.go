package main

import "fmt"

const NMAX int = 100000

type DataUtamaEsport struct {
	daftarTim          string
	jadwalPertandingan int
	hasilPertandingan  int
}
type DataUtama [NMAX]DataUtamaEsport

func main() {
	var namaTim DataUtama
	var pilihOpsiPeran, pilihOpsiMenuPanitia, pilihOpsiMenuPeserta,n int

	MenuUtama(&pilihOpsiPeran)
	if pilihOpsiPeran == 1 {
		menuPanitia(&pilihOpsiMenuPanitia)
		switch pilihOpsiMenuPanitia {
		case 1:
			menambahData(&n,&namaTim)
			jadwalPertandingan(&namaTim)
		case 2:

		case 4:

		case 5:

		case 6:
			MenuUtama(&pilihOpsiPeran)
		}
	} else if pilihOpsiPeran == 2 {
		menuPeserta(&pilihOpsiMenuPeserta)
		switch pilihOpsiMenuPeserta {
		case 1:

		case 2:

		case 3:

		}
	}

}

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

func menambahData(n *int, namaTim *DataUtama) {

	masukanJumlahTim(n)
	masukanNamaTim(*n, namaTim)
}
func masukanJumlahTim(n *int) {
	fmt.Print("\n\t╔══════════════════════════════════════════════════╗\t\n")
	fmt.Print("\t║    Berapa total TIM yang akan berpartisipasi?    ║\t\n")
	fmt.Print("\t╚══════════════════════════════════════════════════╝\t\n")
	fmt.Print("\tMasukkan Pilihan Anda: ")
	fmt.Scan(n)
}
func masukanNamaTim(n int, namaTim *DataUtama) {
	var i int

	fmt.Print("\n\t╔══════════════════════════════════════════════════╗\t\n")
	fmt.Print("\t║    Masukkan Nama Tim yang akan berpartisipasi    ║\t\n")
	fmt.Print("\t║    dan Jadwal pertandingan Dari TIM tersebut     ║\t\n")
	fmt.Print("\t╚══════════════════════════════════════════════════╝\t\n")

	for i = 0; i < n; i++ {
		fmt.Printf("\tMasukkan nama tim ke-%d: ", i+1)
		fmt.Scan(&namaTim[i].daftarTim)

		fmt.Printf("\tMasukkan jadwal pertandingan tim %s: ",namaTim[i].daftarTim)
		fmt.Scan(&namaTim[i].jadwalPertandingan)
	}
}
func jadwalPertandingan(namaTim *DataUtama){
	var timA,timB string
	var jadwal,n int
	masukanPertandingan(*namaTim,&timA,&timB,&jadwal)
	cekTimDanJadwal(timA,timB,jadwal,namaTim,n)
}
func masukanPertandingan(namaTim DataUtama, timA,timB *string, jadwal *int){
	fmt.Print("\n\t╔══════════════════════════════════════════════════╗\t\n")
	fmt.Print("\t║      Masukkan Nama Tim yang akan bertanding      ║\t\n")
	fmt.Print("\t║         dan Jadwal pertandingan Hari ini         ║\t\n")
	fmt.Print("\t╚══════════════════════════════════════════════════╝\t\n")

	fmt.Print("\t TIM A: ")
	fmt.Scan(timA)
	fmt.Print("\t TIM B: ")
	fmt.Scan(timB)
	fmt.Print("\t Jadwal Pertandingan Hari ini: ")
	fmt.Scan(jadwal)
}
func cekTimDanJadwal(timA, timB string, jadwal int, namaTim *DataUtama, n int) {
	var ditemukanA, ditemukanB bool

	for i := 0; i < n; i++ {
		if namaTim[i].daftarTim == timA && namaTim[i].jadwalPertandingan == jadwal {
			ditemukanA = true
		}
		if namaTim[i].daftarTim == timB && namaTim[i].jadwalPertandingan == jadwal {
			ditemukanB = true
		}
	}
	if ditemukanA && ditemukanB {
		fmt.Println("\n Kedua tim ditemukan dan memiliki jadwal yang sama.")
	} else if !ditemukanA && !ditemukanB {
		fmt.Println("\n Kedua tim tidak ditemukan atau jadwalnya salah.")
	} else if !ditemukanA {
		fmt.Println("\n Tim A tidak ditemukan atau jadwalnya salah.")
	} else {
		fmt.Println("\nRESAN GACOR")
	}
}

