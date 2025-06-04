package main
import "fmt"

const N = 1000

type Peminjam struct {
	nama      string
	pinjaman  float64
	bunga     float64
	lamaBulan int
	angsuran  float64
	totalBayar float64
}

var dataPeminjam [N]Peminjam
var totalPeminjam int

func hitungAngsuranFlat(pinjaman float64, bunga float64, lamaBulan int) float64 {
	var bungaTotal, totalPembayaran, angsuranFLat float64
	
	bungaTotal = pinjaman * (bunga / 100) * float64(lamaBulan) / 12
	totalPembayaran = pinjaman + bungaTotal
	angsuranFLat = totalPembayaran / float64(lamaBulan)
	
	return angsuranFLat
}

func tambahData() {
	var angsuran, totalBayar float64
	
	if totalPeminjam >= N {
		fmt.Println("Data peminjam sudah penuh!")
		return
	}

	var nama string
	var pinjaman, bunga float64
	var lamaBulan int

	fmt.Print("Masukkan nama peminjam: ")
	fmt.Scanln(&nama)

	// validasi input (pinjaman yang dimau) 
	pinjaman = -1
	for pinjaman < 0 {
		fmt.Print("Masukkan jumlah pinjaman (Rp): ")
		fmt.Scanln(&pinjaman)
		if pinjaman < 0 {
			fmt.Println("Jumlah pinjaman tidak boleh negatif!")
		}
	}

	// validasi input (bunga)
	bunga = -1
	for bunga < 0 {
		fmt.Print("Masukkan suku bunga per tahun (%): ")
		fmt.Scanln(&bunga)
		if bunga < 0 {
			fmt.Println("Suku bunga tidak boleh negatif!")
		}
	}

	// validasi input (lama pinjaman)
	lamaBulan = 0
	for lamaBulan <= 0 {
		fmt.Print("Masukkan lama pinjaman (bulan): ")
		fmt.Scanln(&lamaBulan)
		if lamaBulan <= 0 {
			fmt.Println("Lama pinjaman harus lebih dari 0 bulan!")
		}
	}

	angsuran = hitungAngsuranFlat(pinjaman, bunga, lamaBulan)
	totalBayar = angsuran * float64(lamaBulan)

	dataPeminjam[totalPeminjam] = Peminjam{
		nama:      nama,
		pinjaman:  pinjaman,
		bunga:     bunga,
		lamaBulan: lamaBulan,
		angsuran:  angsuran,
		totalBayar: totalBayar,
	}
	totalPeminjam++

	fmt.Println("Data berhasil ditambahkan!")
	tampilkanHasil(dataPeminjam[totalPeminjam-1])
}

func tampilkanHasil(p Peminjam) {
	fmt.Println("\n=== Hasil Simulasi Kredit ===")
	fmt.Println("Nama Peminjam        :", p.nama)
	fmt.Printf("Jumlah Pinjaman      : Rp%.0f\n", p.pinjaman)
	fmt.Printf("Suku Bunga Tahunan   : %.2f%%\n", p.bunga)
	fmt.Println("Lama Pinjaman        :", p.lamaBulan, "bulan")
	fmt.Printf("Angsuran per Bulan   : Rp%.0f\n", p.angsuran)
	fmt.Printf("Total Pembayaran     : Rp%.0f\n", p.totalBayar)
}

func tampilkanSemua() {
	if totalPeminjam == 0 {
		fmt.Println("Belum ada data.")
	}
	fmt.Println("\n=== Data Kredit Nasabah ===")
	for i := 0; i < totalPeminjam; i++ {
		p := dataPeminjam[i]
		fmt.Printf("%d. %s | Pinjaman: Rp%.0f | Bunga: %.2f%% | Lama: %d bulan | Angsuran: Rp%.0f | Total Bayar: Rp%.0f\n",
			i+1, p.nama, p.pinjaman, p.bunga, p.lamaBulan, p.angsuran, p.totalBayar)
	}
}

func cariNama() {
	var cari string
	var idxDitemukan = -1

	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scanln(&cari)

	for i := 0; i < totalPeminjam; i++ {
		if dataPeminjam[i].nama == cari {
			idxDitemukan = i
		}
	}

	if idxDitemukan != -1 {
		p := dataPeminjam[idxDitemukan]
		fmt.Printf("Ditemukan: %s\n", p.nama)
		fmt.Printf("Total bayar: Rp%.0f\n", p.totalBayar)
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func urutkanTotalBayar() {
	
	var i, j int
	
	if totalPeminjam == 0 {
		fmt.Println("Belum ada data untuk diurutkan.")
		return
	}
	// Bubble sort sederhana berdasarkan totalBayar ascending
	for i = 0; i < totalPeminjam-1; i++ {
		for j = 0; j < totalPeminjam-i-1; j++ {
			if dataPeminjam[j].totalBayar > dataPeminjam[j+1].totalBayar {
				dataPeminjam[j], dataPeminjam[j+1] = dataPeminjam[j+1], dataPeminjam[j]
			}
		}
	}
	fmt.Println("Data berhasil diurutkan berdasarkan total pembayaran.")
}

func cariEkstrim() {
	
	var i int 
	
	if totalPeminjam == 0 {
		fmt.Println("Belum ada data.")
	}

	min := dataPeminjam[0]
	max := dataPeminjam[0]

	i = 1
	for i < totalPeminjam {
		if dataPeminjam[i].totalBayar < min.totalBayar {
			min = dataPeminjam[i]
		}
		if dataPeminjam[i].totalBayar > max.totalBayar {
			max = dataPeminjam[i]
		}
		i++
	}

	fmt.Printf("Total Bayar TERKECIL: %s = Rp%.0f\n", min.nama, min.totalBayar)
	fmt.Printf("Total Bayar TERBESAR: %s = Rp%.0f\n", max.nama, max.totalBayar)
}

func Menu() {
	var pilih int
	for pilih != 6 {
		fmt.Println("\n=== MENU UTAMA SIMULASI KREDIT SOLO ===")
		fmt.Println("1. Tambah Data Kredit")
		fmt.Println("2. Tampilkan Semua Data")
		fmt.Println("3. Cari Data berdasarkan Nama")
		fmt.Println("4. Urutkan berdasarkan Total Bayar")
		fmt.Println("5. Tampilkan Nilai Ekstrim (Min/Max)")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilih)

		if pilih == 1 {
			tambahData()
		} else if pilih == 2 {
			tampilkanSemua()
		} else if pilih == 3 {
			cariNama()
		} else if pilih == 4 {
			urutkanTotalBayar()
		} else if pilih == 5 {
			cariEkstrim()
		} else if pilih == 6 {
			fmt.Println("Terima kasih!")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func main(){
	Menu()
}
