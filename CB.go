package main

import "fmt"

// ==========================================
// BAB 1: STRUKTUR DATA
// ==========================================

type Capaian struct {
	NamaTugas          string
	TingkatKepentingan int
	TanggalSelesai     string
	IsSelesai          bool
	DeskripsiProgres   string
	SkorStres          int
	SkorMood           int
}

type ArrCapaian [1000]Capaian

// Fungsi pembantu murni ALPRO: Membaca string berspasi karakter demi karakter
func bacaString() string {
	var hasil string
	var char byte
	
	// Loop 1: Lewati sisa "Enter" (\n) atau spasi kosong dari input menu sebelumnya
	for {
		fmt.Scanf("%c", &char)
		if char != '\n' && char != '\r' && char != ' ' {
			hasil = hasil + string(char) // Simpan huruf pertama
			break
		}
	}
	
	// Loop 2: Baca sisa karakternya satu per satu sampai ditekan "Enter" lagi
	for {
		fmt.Scanf("%c", &char)
		if char == '\n' || char == '\r' {
			break
		}
		hasil = hasil + string(char)
	}
	
	return hasil
}

// ==========================================
// BAB 2.2: MANAJEMEN DATA (CRUD) - PENGEMBANG 1
// ==========================================

func TambahData(A *ArrCapaian, n *int) {
	if *n >= 1000 {
		fmt.Println("Kapasitas penyimpanan penuh!")
		return
	}
	
	fmt.Println("\n--- TAMBAH TUGAS BARU ---")
	fmt.Print("Nama Tugas           : ")
	A[*n].NamaTugas = bacaString()
	
	fmt.Print("Tingkat Kepentingan (1-5): ")
	fmt.Scan(&A[*n].TingkatKepentingan)
	
	fmt.Print("Target Tanggal (YYYYMMDD): ")
	fmt.Scan(&A[*n].TanggalSelesai)
	
	A[*n].IsSelesai = false
	*n = *n + 1
	fmt.Println("Data berhasil ditambahkan.")
}

func UbahData(A *ArrCapaian, n int) {
	fmt.Println("\n--- UBAH DATA TUGAS ---")
	fmt.Print("Masukkan Nama Tugas yang ingin diubah: ")
	target := bacaString()
	
	idx := CariBerdasarkanNama(*A, n, target)
	if idx == -1 {
		fmt.Println("Tugas tidak ditemukan.")
		return
	}
	
	fmt.Print("Nama Tugas Baru           : ")
	A[idx].NamaTugas = bacaString()
	
	fmt.Print("Tingkat Kepentingan (1-5) : ")
	fmt.Scan(&A[idx].TingkatKepentingan)
	
	fmt.Print("Target Tanggal (YYYYMMDD) : ")
	fmt.Scan(&A[idx].TanggalSelesai)
	fmt.Println("Data berhasil diubah.")
}

func SelesaikanTugas(A *ArrCapaian, n int) {
	fmt.Println("\n--- SELESAIKAN TUGAS ---")
	fmt.Print("Masukkan Nama Tugas: ")
	target := bacaString()
	
	idx := CariBerdasarkanNama(*A, n, target)
	if idx == -1 {
		fmt.Println("Tugas tidak ditemukan.")
		return
	}
	
	if A[idx].IsSelesai {
		fmt.Println("Tugas ini sudah diselesaikan sebelumnya.")
		return
	}
	
	fmt.Print("Deskripsi Progres : ")
	A[idx].DeskripsiProgres = bacaString()
	
	fmt.Print("Skor Stres (1-10) : ")
	fmt.Scan(&A[idx].SkorStres)
	
	fmt.Print("Skor Mood (1-10)  : ")
	fmt.Scan(&A[idx].SkorMood)
	
	A[idx].IsSelesai = true
	fmt.Println("Selamat! Tugas berhasil diselesaikan.")
}

func HapusData(A *ArrCapaian, n *int) {
	fmt.Println("\n--- HAPUS TUGAS ---")
	fmt.Print("Masukkan Nama Tugas yang akan dihapus: ")
	target := bacaString()
	
	idx := CariBerdasarkanNama(*A, *n, target)
	if idx == -1 {
		fmt.Println("Tugas tidak ditemukan.")
		return
	}
	
	// Menggeser array ke kiri untuk menutup celah
	for i := idx; i < *n-1; i++ {
		A[i] = A[i+1]
	}
	*n = *n - 1
	fmt.Println("Data berhasil dihapus.")
}

// ==========================================
// BAB 2.3 & 2.4 & 2.5: SEARCHING, SORTING, STATS - PENGEMBANG 2
// ==========================================

func CariBerdasarkanNama(A ArrCapaian, n int, target string) int {
	for i := 0; i < n; i++ {
		if A[i].NamaTugas == target {
			return i
		}
	}
	return -1
}

func urutkanTanggalInternal(A *ArrCapaian, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if A[j].TanggalSelesai > A[j+1].TanggalSelesai {
				temp := A[j]
				A[j] = A[j+1]
				A[j+1] = temp
			}
		}
	}
}

func CariBerdasarkanTanggal(A ArrCapaian, n int, target string) int {
	urutkanTanggalInternal(&A, n)
	low := 0
	high := n - 1
	
	for low <= high {
		mid := (low + high) / 2
		if A[mid].TanggalSelesai == target {
			return mid
		} else if A[mid].TanggalSelesai < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// Selection Sort (Descending)
func UrutkanKepentingan(A *ArrCapaian, n int) {
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i+1; j < n; j++ {
			if A[j].TingkatKepentingan > A[maxIdx].TingkatKepentingan {
				maxIdx = j
			}
		}
		temp := A[i]
		A[i] = A[maxIdx]
		A[maxIdx] = temp
	}
	fmt.Println("Data berhasil diurutkan berdasarkan tingkat kepentingan (Tertinggi ke Terendah).")
}

// Insertion Sort (Descending)
func UrutkanMood(A *ArrCapaian, n int) {
	for i := 1; i < n; i++ {
		key := A[i]
		j := i - 1
		for j >= 0 && A[j].SkorMood < key.SkorMood {
			A[j+1] = A[j]
			j = j - 1
		}
		A[j+1] = key
	}
	fmt.Println("Data berhasil diurutkan berdasarkan Skor Mood (Terbaik ke Terburuk).")
}

func TampilkanStatistikMingguan(A ArrCapaian, n int) {
	if n == 0 {
		fmt.Println("Belum ada data tugas.")
		return
	}

	var tanggalAwal, tanggalAkhir string
	fmt.Print("\nMasukkan Tanggal Awal Minggu (YYYYMMDD) : ")
	fmt.Scan(&tanggalAwal)
	fmt.Print("Masukkan Tanggal Akhir Minggu (YYYYMMDD): ")
	fmt.Scan(&tanggalAkhir)

	totalSkorStres := 0
	jumlahTugasSelesaiMingguIni := 0
	totalTugasMingguIni := 0

	for i := 0; i < n; i++ {
		if A[i].TanggalSelesai >= tanggalAwal && A[i].TanggalSelesai <= tanggalAkhir {
			totalTugasMingguIni++
			if A[i].IsSelesai {
				totalSkorStres += A[i].SkorStres
				jumlahTugasSelesaiMingguIni++
			}
		}
	}

	fmt.Printf("\n=== STATISTIK MINGGU (%s - %s) ===\n", tanggalAwal, tanggalAkhir)
	
	if jumlahTugasSelesaiMingguIni > 0 {
		// Menggunakan float64 agar pembagian menghasilkan angka desimal yang benar
		rataRataStres := float64(totalSkorStres) / float64(jumlahTugasSelesaiMingguIni)
		fmt.Printf("Rata-rata Skor Stres                 : %.2f\n", rataRataStres)
	} else {
		fmt.Println("Rata-rata Skor Stres                 : Tidak ada tugas selesai.")
	}

	if totalTugasMingguIni > 0 {
		persentaseSukses := (float64(jumlahTugasSelesaiMingguIni) / float64(totalTugasMingguIni)) * 100
		fmt.Printf("Persentase Keberhasilan              : %.2f%%\n", persentaseSukses)
	} else {
		fmt.Println("Persentase Keberhasilan              : 0%")
	}
}

func TampilkanSemuaData(A ArrCapaian, n int) {
	if n == 0 {
		fmt.Println("Data masih kosong.")
		return
	}
	fmt.Println("\n--- DAFTAR TUGAS ---")
	for i := 0; i < n; i++ {
		status := "Belum Selesai"
		if A[i].IsSelesai {
			status = "Selesai"
		}
		fmt.Printf("%d. [%s] %s | Tgl: %s | Prioritas: %d | Stres: %d | Mood: %d\n", 
			i+1, status, A[i].NamaTugas, A[i].TanggalSelesai, A[i].TingkatKepentingan, A[i].SkorStres, A[i].SkorMood)
	}
}

// ==========================================
// PROGRAM UTAMA (MAIN)
// ==========================================

func main() {
	var DataCapaian ArrCapaian
	var totalData int = 0
	var pilihan int

	for {
		fmt.Println("\n==================================")
		fmt.Println("      APLIKASI MINDSTONE")
		fmt.Println("==================================")
		fmt.Println("1. Tambah Tugas")
		fmt.Println("2. Ubah Tugas")
		fmt.Println("3. Selesaikan Tugas")
		fmt.Println("4. Hapus Tugas")
		fmt.Println("5. Cari Berdasarkan Nama (Sequential)")
		fmt.Println("6. Cari Berdasarkan Tanggal (Binary)")
		fmt.Println("7. Urutkan Berdasarkan Kepentingan")
		fmt.Println("8. Urutkan Berdasarkan Mood")
		fmt.Println("9. Tampilkan Statistik Mingguan")
		fmt.Println("10. Tampilkan Semua Data")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		// Rantai if - else if klasik, sesuai aturan dasar
		if pilihan == 1 {
			TambahData(&DataCapaian, &totalData)
		} else if pilihan == 2 {
			UbahData(&DataCapaian, totalData)
		} else if pilihan == 3 {
			SelesaikanTugas(&DataCapaian, totalData)
		} else if pilihan == 4 {
			HapusData(&DataCapaian, &totalData)
		} else if pilihan == 5 {
			fmt.Print("\nMasukkan nama tugas yang dicari: ")
			target := bacaString()
			idx := CariBerdasarkanNama(DataCapaian, totalData, target)
			if idx != -1 {
				fmt.Printf("Ditemukan pada indeks ke-%d (Tanggal Target: %s)\n", idx, DataCapaian[idx].TanggalSelesai)
			} else {
				fmt.Println("Data tidak ditemukan.")
			}
		} else if pilihan == 6 {
			fmt.Print("\nMasukkan tanggal (YYYYMMDD) yang dicari: ")
			var target string
			fmt.Scan(&target)
			idx := CariBerdasarkanTanggal(DataCapaian, totalData, target)
			if idx != -1 {
				fmt.Printf("Ditemukan! Salah satu tugas di tanggal tersebut adalah: %s\n", DataCapaian[idx].NamaTugas)
			} else {
				fmt.Println("Tidak ada tugas pada tanggal tersebut.")
			}
		} else if pilihan == 7 {
			UrutkanKepentingan(&DataCapaian, totalData)
			TampilkanSemuaData(DataCapaian, totalData)
		} else if pilihan == 8 {
			UrutkanMood(&DataCapaian, totalData)
			TampilkanSemuaData(DataCapaian, totalData)
		} else if pilihan == 9 {
			TampilkanStatistikMingguan(DataCapaian, totalData)
		} else if pilihan == 10 {
			TampilkanSemuaData(DataCapaian, totalData)
		} else if pilihan == 0 {
			fmt.Println("Terima kasih telah menggunakan MindStone.")
			break
		} else {
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}