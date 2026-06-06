================================================================================
                    README - DOKUMENTASI RESMI APLIKASI MINDSTONE
         (PEMANTAU CAPAIAN KERJA DAN KESEHATAN MENTAL BERBASIS TERMINAL)
================================================================================

Aplikasi MindStone dirancang sebagai solusi terintegrasi untuk membantu pekerja
profesional dan mahasiswa dalam mendokumentasikan target capaian kerja (milestone)
sekaligus memantau kondisi kesehatan mental secara berkala. Aplikasi ini dibangun
menggunakan bahasa pemrograman Go (Golang) dengan mematuhi standar arsitektur 
komputer, efisiensi memori statis, serta modularitas fungsi dan prosedur.

Dokumentasi ini disusun secara rinci sebagai panduan operasional bagi pengguna
dan panduan teknis bagi Inspektur Program / Dosen Penguji.

--------------------------------------------------------------------------------
1. SPESIFIKASI STRUKTUR DATA & MANAJEMEN MEMORI
--------------------------------------------------------------------------------
Untuk menjamin stabilitas eksekusi runtime dan mencegah kebocoran memori (memory
leak), aplikasi menggunakan skema alokasi memori statis terpusat:

A. Tipe Bentukan Utama (Struct Capaian)
   Setiap data target kerja diisolasi di dalam sebuah struct dengan atribut:
   - NamaTugas (string)          : Judul atau nama target capaian kerja.
   - TingkatKepentingan (int)    : Skala prioritas tugas dengan rentang nilai 1-5.
   - TanggalSelesai (string)     : Format standar YYYYMMDD (komparasi leksikografis).
   - IsSelesai (boolean)         : Indikator status (False = Belum, True = Selesai).
   - DeskripsiProgres (string)   : Catatan atau evaluasi penyelesaian tugas.
   - SkorStres (int)             : Skala tingkat stres saat pengerjaan (1-10).
   - SkorMood (int)              : Skala suasana hati setelah tugas selesai (1-10).

B. Array Penyimpanan Statis (ArrCapaian)
   Data disimpan dalam tipe array bentukan `[1000]Capaian`. Alokasi memori sebesar
   1000 elemen dideklarasikan sejak awal eksekusi program.

C. Manajemen Pointer (Pass-by-Reference)
   Seluruh prosedur manipulasi data (Tambah, Ubah, Hapus, Selesaikan) menggunakan
   pointer terhadap array utama (`*ArrCapaian`) dan pointer terhadap counter data
   (`*int` untuk variabel totalData). Hal ini memastikan manipulasi data dilakukan
   secara langsung pada alamat memori asli tanpa melakukan kloning array yang
   memboroskan sumber daya.

--------------------------------------------------------------------------------
2. FITUR UTAMA DAN IMPLEMENTASI ALGORITMA
--------------------------------------------------------------------------------
A. Modul Manajemen Data (CRUD)
   - Tambah Data: Memasukkan elemen tugas baru langsung pada indeks ke-n di dalam
     array, kemudian melakukan increment nilai counter (n = n + 1).
   - Ubah Data: Memungkinkan pengguna mengoreksi kesalahan ketik nama, tanggal,
     maupun tingkat kepentingan sebelum tugas ditandai selesai.
   - Selesaikan Tugas: Mengubah status `IsSelesai` menjadi True serta merekam
     indikator psikologis pengguna (Deskripsi Progres, Skor Stres, Skor Mood).
   - Hapus Data (Logika Pergeseran): Menghapus elemen pada indeks tertentu, lalu
     menjalankan loop untuk menggeser seluruh sisa elemen di sebelah kanannya ke
     arah kiri (A[i] = A[i+1]). Celah memori tertutup rapat, dan counter dikurangi (n = n - 1).

B. Modul Pencarian Data (Searching)
   - Pencarian Berdasarkan Nama (Sequential Search):
     Melakukan iterasi linear satu per satu dari indeks 0 hingga n-1. Kompleksitas
     waktu O(n). Digunakan karena data nama tugas masuk secara acak dan tidak terurut.
   - Pencarian Berdasarkan Tanggal (Binary Search):
     Menggunakan algoritma pembelahan area pencarian (low, mid, high) dengan 
     kompleksitas O(log n). Agar Binary Search bekerja valid, modul ini secara
     internal memanggil fungsi pengurutan tanggal terlebih dahulu sebelum membelah data.

C. Modul Pengurutan Data (Sorting)
   - Urutkan Kepentingan (Selection Sort - Descending):
     Memindai tingkat kepentingan tertinggi dari sisa array yang belum terurut,
     lakukan operasi tukar posisi (swap) ke area depan. Berguna untuk memunculkan prioritas utama.
   - Urutkan Mood (Insertion Sort - Descending):
     Mengambil elemen satu per satu dan menyisipkannya mundur ke posisi yang tepat
     di antara elemen yang sudah terurut (seperti menyusun kartu remi).

D. Modul Analitik (Statistik Mingguan)
   - Menghitung rata-rata tingkat stres pengguna dan persentase keberhasilan tugas
     khusus pada rentang minggu tertentu. Memanfaatkan perbandingan string leksikografis
     (operator >= dan <= pada pola YYYYMMDD) untuk memfilter rentang tanggal tanpa
     bergantung pada library eksternal.

--------------------------------------------------------------------------------
3. PRASYARAT DAN CARA MENJALANKAN PROGRAM
--------------------------------------------------------------------------------
Prasyarat Sistem:
- Perangkat keras telah terinstal Go Compiler (Golang) versi 1.13 atau yang lebih baru.

Langkah-Langkah Menjalankan:
1. Buka Terminal / Command Prompt / PowerShell.
2. Navigasikan direktori aktif ke folder tempat file `main.go` disimpan.
   Contoh: cd /path/to/project/mindstone
3. Jalankan perintah kompilasi dan eksekusi langsung berikut:
   go run main.go
4. Program akan langsung memunculkan antarmuka menu utama MindStone di terminal.

--------------------------------------------------------------------------------
4. PANDUAN PENGGUNAAN INTERAKTIF (LANGKAH DEMI LANGKAH)
--------------------------------------------------------------------------------
Ketika program berjalan, Anda akan disajikan 10 menu operasional utama:

[Menu 1] Tambah Tugas
- Ketik `1` lalu tekan Enter.
- Masukkan Nama Tugas (Mendukung spasi, contoh: "Revisi Perhitungan Matriks").
- Masukkan Skala Kepentingan (Angka 1 sampai 5).
- Masukkan Target Tanggal dengan pola 8 digit angka YYYYMMDD (Contoh: "20260601").

[Menu 2] Ubah Tugas
- Ketik `2` lalu tekan Enter.
- Masukkan nama tugas lama yang ingin diperbaiki. Jika ditemukan, sistem akan
  meminta input nama baru, skala kepentingan baru, dan tanggal target baru.

[Menu 3] Selesaikan Tugas
- Ketik `3` lalu tekan Enter.
- Masukkan nama tugas yang telah Anda kerjakan.
- Masukkan Deskripsi Progres (Catatan penutup pengerjaan tugas).
- Masukkan Skor Stres Anda saat mengerjakan tugas tersebut (Skala 1-10).
- Masukkan Skor Mood Anda setelah tugas berhasil diselesaikan (Skala 1-10).
- Atribut `IsSelesai` otomatis berubah dari False menjadi True.

[Menu 4] Hapus Tugas
- Ketik `4` lalu tekan Enter.
- Masukkan nama tugas yang ingin dihapus secara permanen dari sistem memori.

[Menu 5] Cari Berdasarkan Nama (Sequential Search)
- Ketik `5` lalu tekan Enter.
- Masukkan nama tugas secara spesifik. Sistem mengembalikan informasi indeks memori
  dan tanggal target dari tugas tersebut.

[Menu 6] Cari Berdasarkan Tanggal (Binary Search)
- Ketik `6` lalu tekan Enter.
- Masukkan target tanggal (YYYYMMDD). Sistem secara otomatis mengurutkan memori internal
  terlebih dahulu, membelah pencarian, dan menampilkan nama tugas yang berada pada tanggal tersebut.

[Menu 7] Urutkan Berdasarkan Kepentingan
- Ketik `7` lalu tekan Enter.
- Sistem mengeksekusi Selection Sort dan langsung menyajikan tabel daftar tugas
  terurut dari skala kepentingan tertinggi (Prioritas 5) hingga terendah (Prioritas 1).

[Menu 8] Urutkan Berdasarkan Mood
- Ketik `8` lalu tekan Enter.
- Sistem mengeksekusi Insertion Sort dan menyajikan rekam jejak tugas berdasarkan
  skor suasana hati terbaik (Skor 10) hingga terburuk (Skor 1).

[Menu 9] Tampilkan Statistik Mingguan
- Ketik `9` lalu tekan Enter.
- Masukkan Tanggal Awal rentang minggu (YYYYMMDD).
- Masukkan Tanggal Akhir rentang minggu (YYYYMMDD).
- Sistem memfilter data dan menampilkan rata-rata skor stres dari tugas yang selesai,
  serta persentase rasio keberhasilan target kerja spesifik pada minggu tersebut.

[Menu 10] Tampilkan Semua Data
- Ketik `10` lalu tekan Enter untuk melihat kondisi mentah seluruh baris data di dalam array.

[Menu 0] Keluar
- Ketik `0` untuk menghentikan loop program dan keluar dari aplikasi secara aman.

--------------------------------------------------------------------------------
5. CATATAN TEKNIS UNTUK INSPEKTUR PROGRAM (PANDUAN DEFLEKSI BUG)
--------------------------------------------------------------------------------
Bila Inspektur Program menanyakan aspek teknis penanganan bug masukan pada kode ini:
1. Penanganan Input Buffer Terlewat (Input Skip Bug):
   Program ini telah dilengkapi dengan fungsi pembantu `bacaString(scanner)`. Fungsi
   ini menggunakan loop `for text == ""` yang secara cerdas mendeteksi dan membuang
   sisa karakter newline (`
`) yang tertinggal di input buffer terminal setelah
   pengguna menginput angka pilihan menu menggunakan `fmt.Scan`. Hal ini menjamin
   bahwa input string Nama Tugas tidak akan pernah terlewati (jump/skip).
2. Validasi Batas Array (Array Boundary Guard):
   Pada modul `TambahData`, program secara eksplisit memeriksa kondisi `if *n >= 1000`.
   Jika batasan ini tercapai, program menolak input baru secara halus guna menghindari
   terjadinya error fatal `index out of range` pada memori statis komputer.
3. Penanganan Division by Zero:
   Pada kalkulasi statistik, pembagian nilai float64 dilindungi oleh pengkondisian
   `if jumlahTugasSelesaiMingguIni > 0` dan `if totalTugasMingguIni > 0`, sehingga
   program terbebas dari crash matematika runtime saat data minggu terkait kosong.
================================================================================
