package main

import "fmt"

const MAX = 100

type Peminjam struct {
	Nama string
	JumlahPinjaman float64
	Tenor int
	Bunga float64
	CicilanBulanan float64
	StatusLunas bool
}

var dataPeminjam [MAX]Peminjam
var jumlahData int = 0

func main() {
	for {
		fmt.Println("\n===== MENU APLIKASI PINJAMAN =====")
		fmt.Println("1. Tambah Peminjam")
		fmt.Println("2. Ubah Data Peminjam")
		fmt.Println("3. Cari Peminjam (Sequential)")
		fmt.Println("4. Cari Peminjam (Binary)")
		fmt.Println("5. Urutkan Berdasarkan Jumlah Pinjaman (Selection Sort)")
		fmt.Println("6. Urutkan Berdasarkan Tenor (Insertion Sort)")
		fmt.Println("7. Hapus Data Peminjam")
		fmt.Println("8. Tampilkan Laporan")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahPeminjam()
		case 2:
			ubahPeminjam()
		case 3:
			cariNamaSequential()
		case 4:
			cariNamaBinary()
		case 5:
			selectionSortByJumlah()
			fmt.Println("Data berhasil diurutkan berdasarkan jumlah pinjaman.")
		case 6:
			insertionSortByTenor()
			fmt.Println("Data berhasil diurutkan berdasarkan tenor.")
		case 7:
			hapusPeminjam()
		case 8:
			tampilkanLaporan()
		case 9:
			fmt.Println("Terima kasih! Program selesai.")
			return
		default:
		fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

func tambahPeminjam() {
	if jumlahData >= MAX {
		fmt.Println("Data penuh, tidak bisa menambahkan lagi.")
		return
	}

	var p Peminjam
	fmt.Print("Nama: ")
	fmt.Scanln(&p.Nama)
	fmt.Print("Jumlah Pinjaman: ")
	fmt.Scanln(&p.JumlahPinjaman)
	fmt.Print("Tenor (bulan): ")
	fmt.Scanln(&p.Tenor)
	fmt.Print("Bunga (% per tahun): ")
	fmt.Scanln(&p.Bunga)

	hitungCicilan(&p)
	p.StatusLunas = false

	dataPeminjam[jumlahData] = p
	jumlahData++

	fmt.Println("Data berhasil ditambahkan!")
}

func ubahPeminjam() {
	var nama string
	fmt.Print("Masukkan nama peminjam yang mau diubah: ")
	fmt.Scanln(&nama)

	var index int = cariPeminjamSequential(nama)
	if index == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	fmt.Println("Masukkan data baru:")
	fmt.Print("Jumlah Pinjaman: ")
	fmt.Scanln(&dataPeminjam[index].JumlahPinjaman)
	fmt.Print("Tenor (bulan): ")
	fmt.Scanln(&dataPeminjam[index].Tenor)
	fmt.Print("Bunga (% per tahun): ")
	fmt.Scanln(&dataPeminjam[index].Bunga)

	hitungCicilan(&dataPeminjam[index])
	fmt.Println("Data berhasil diubah!")
}

func hapusPeminjam() {
	var nama string
	fmt.Print("Masukkan nama peminjam yang mau dihapus: ")
	fmt.Scanln(&nama)

	var index int = cariPeminjamSequential(nama)
	if index == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	var i int 
	for i = index; i < jumlahData-1; i++ {
		dataPeminjam[i] = dataPeminjam[i+1]
	}
	jumlahData--

	fmt.Println("Data berhasil dihapus!")
}

func hitungCicilan(p *Peminjam) {
	var bungaPerBulan float64 = p.Bunga / 12 / 100
	var totalBunga float64 = p.JumlahPinjaman * bungaPerBulan * float64(p.Tenor)
	var totalPinjaman float64 = p.JumlahPinjaman + totalBunga
	p.CicilanBulanan = totalPinjaman / float64(p.Tenor)
}

func cariPeminjamSequential(nama string) int {
	var i int 
	for i = 0; i < jumlahData; i++ {
		if dataPeminjam[i].Nama == nama {
			return i
		}
	}
	return -1
}

func cariNamaSequential() {
	var nama string
	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scanln(&nama)

	var idx int = cariPeminjamSequential(nama)
	if idx != -1 {
		tampilkanDetail(dataPeminjam[idx])
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func cariPeminjamBinary(nama string) int {
	var left int = 0
	var right int = jumlahData - 1
	var mid int 

	for left <= right {
		mid = (left + right) / 2
		
		if dataPeminjam[mid].Nama == nama {
			return mid
		} else if dataPeminjam[mid].Nama < nama {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func cariNamaBinary() {
	var nama string
	fmt.Print("Masukkan nama yang dicari (pastikan data sudah diurutkan): ")
	fmt.Scanln(&nama)

	idx := cariPeminjamBinary(nama)
	if idx != -1 {
		tampilkanDetail(dataPeminjam[idx])
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func selectionSortByJumlah() {
	var i int 
	for i = 0; i < jumlahData-1; i++ {
		var min int = i
		var j int 
		for j = i + 1; j < jumlahData; j++ {
			if dataPeminjam[j].JumlahPinjaman < dataPeminjam[min].JumlahPinjaman {
				min = j
			}
		}
		var temp Peminjam = dataPeminjam[i]
		dataPeminjam[i] = dataPeminjam[min]
		dataPeminjam[min] = temp 
	}
}

func insertionSortByTenor() {
	var i int 
	for i = 1; i < jumlahData; i++ {
		var temp Peminjam 
		temp = dataPeminjam[i]
		
		var j int 
		j = i - 1
		for j >= 0 && dataPeminjam[j].Tenor > temp.Tenor {
			dataPeminjam[j+1] = dataPeminjam[j]
			j--
		}
		dataPeminjam[j+1] = temp
	}
}

func tampilkanLaporan() {
	fmt.Println("===== Laporan Pinjaman =====")
	var i int 
	for i = 0; i < jumlahData; i++ {
		tampilkanDetail(dataPeminjam[i])
	}
}

func tampilkanDetail(p Peminjam) {
	fmt.Printf("\nNama: %s\nPinjaman: %.2f\nTenor: %d bulan\nBunga: %.2f%%\nCicilan: %.2f\nStatus Lunas: %v\n",
		p.Nama, p.JumlahPinjaman, p.Tenor, p.Bunga, p.CicilanBulanan, p.StatusLunas)
}