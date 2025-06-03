package main

import "fmt"

const MAX int = 10

type peminjam struct {
	nama           string
	statusLunas    bool
	tenor          int
	bunga          float64
	cicilanBulanan float64
	jumlahPinjaman float64
}

type tabPeminjam [MAX]peminjam

func main() {
	var pilihan, jumlahPeminjam int
	var dataPeminjam tabPeminjam
	jumlahPeminjam = 0

	for pilihan != 9 {
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

		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahPeminjam(&dataPeminjam, &jumlahPeminjam)
		} else if pilihan == 2 {
			ubahPeminjam(&dataPeminjam, jumlahPeminjam)
		} else if pilihan == 3 {
			cariNamaSequential(dataPeminjam, jumlahPeminjam)
		} else if pilihan == 4 {
			cariNamaBinary(dataPeminjam, jumlahPeminjam)
		} else if pilihan == 5 {
			selectionSortByJumlah(dataPeminjam, jumlahPeminjam)
			fmt.Println("Data berhasil diurutkan berdasarkan jumlah pinjaman.")
		} else if pilihan == 6 {
			insertionSortByTenor(dataPeminjam, jumlahPeminjam)
			fmt.Println("Data berhasil diurutkan berdasarkan tenor.")
		} else if pilihan == 7 {
			hapusPeminjam(&dataPeminjam, &jumlahPeminjam)
		} else if pilihan == 8 {
			tampilkanLaporan(dataPeminjam, jumlahPeminjam)
		} else if pilihan == 9 {
			fmt.Println("Terima kasih telah menggunakan program ini")
		} else {
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

func tambahPeminjam(A *tabPeminjam, n *int) {
	var p peminjam
	if *n >= MAX {
		fmt.Println("Data penuh, tidak bisa menambahkan lagi.")
		return
	}

	fmt.Print("Nama: ")
	fmt.Scan(&p.nama)
	fmt.Print("Jumlah Pinjaman: ")
	fmt.Scan(&p.jumlahPinjaman)
	fmt.Print("Tenor (bulan): ")
	fmt.Scan(&p.tenor)
	fmt.Print("Bunga (% per tahun): ")
	fmt.Scan(&p.bunga)

	hitungCicilan(&p)
	p.statusLunas = false

	A[*n] = p
	*n++

	fmt.Println("Data peminjam berhasil ditambahkan!")
}

func ubahPeminjam(A *tabPeminjam, n int) {
	var nama string
	var idx int
	fmt.Print("Masukkan nama peminjam yang mau diubah: ")
	fmt.Scan(&nama)
	idx = cariPeminjamSequential(nama, *A, n)

	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	fmt.Println("Masukkan data baru:")
	fmt.Print("Jumlah Pinjaman: ")
	fmt.Scan(&A[idx].jumlahPinjaman)
	fmt.Print("Tenor (bulan): ")
	fmt.Scan(&A[idx].tenor)
	fmt.Print("Bunga (% per tahun): ")
	fmt.Scan(&A[idx].bunga)

	hitungCicilan(&A[idx])
	fmt.Println("Data berhasil diubah!")
}

func hapusPeminjam(A *tabPeminjam, n *int) {
	var nama string
	var i, idx int

	tampilkanLaporan(*A, *n)
	fmt.Print("Masukkan nama peminjam yang mau dihapus: ")
	fmt.Scan(&nama)

	idx = cariPeminjamSequential(nama, *A, *n)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	for i = idx; i < *n-1; i++ {
		A[i] = A[i+1]
	}
	*n--

	fmt.Println("Data berhasil dihapus!")
}

func hitungCicilan(p *peminjam) {
	var bungaPerBulan, totalBunga, totalPinjaman float64

	bungaPerBulan = p.bunga / 12 / 100
	totalBunga = p.jumlahPinjaman * bungaPerBulan * float64(p.tenor)
	totalPinjaman = p.jumlahPinjaman + totalBunga
	p.cicilanBulanan = totalPinjaman / float64(p.tenor)
}

func cariPeminjamSequential(nama string, A tabPeminjam, n int) int {
	var i int
	for i = 0; i < n; i++ {
		if A[i].nama == nama {
			return i
		}
	}
	return -1
}

func cariNamaSequential(A tabPeminjam, n int) {
	var nama string
	var idx int
	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scan(&nama)

	idx = cariPeminjamSequential(nama, A, n)
	if idx != -1 {
		tampilkanDetail(A[idx])
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func cariPeminjamBinary(nama string, A tabPeminjam, n int) int {
	var left, right, mid int
	left = 0
	right = n - 1

	for left <= right {
		mid = (left + right) / 2

		if A[mid].nama == nama {
			return mid
		} else if A[mid].nama < nama {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func cariNamaBinary(A tabPeminjam, n int) {
	var nama string
	var idx int
	fmt.Print("Masukkan nama yang dicari (pastikan data sudah diurutkan): ")
	fmt.Scan(&nama)

	idx = cariPeminjamBinary(nama, A, n)
	if idx != -1 {
		tampilkanDetail(A[idx])
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func selectionSortByJumlah(A tabPeminjam, n int) {
	var i, idx, pass int
	var temp peminjam

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if A[idx].jumlahPinjaman > A[i].jumlahPinjaman {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
	tampilkanLaporan(A, n)
}

func insertionSortByTenor(A tabPeminjam, n int) {
	var pass, i int
	var temp peminjam

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.tenor < A[i-1].tenor {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}

	tampilkanLaporan(A, n)
}

func tampilkanLaporan(A tabPeminjam, n int) {
	var i int
	fmt.Println("===== Laporan Pinjaman =====")
	for i = 0; i < n; i++ {
		tampilkanDetail(A[i])
	}
}

func tampilkanDetail(p peminjam) {
	fmt.Printf("\nNama: %s\nPinjaman: %.2f\nTenor: %d bulan\nBunga: %.2f%%\nCicilan: %.2f\nStatus Lunas: %v\n",
		p.nama, p.jumlahPinjaman, p.tenor, p.bunga, p.cicilanBulanan, p.statusLunas)
}
