package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type DaftarTugas struct {
	nama             string
	jenisTugas       string
	deskripsiTugas   string
	tingkatKesulitan int
	estimasiWaktu    int
	KategoriRuangan  int
}

var dataSelesai []DaftarTugas
var reader = bufio.NewReader(os.Stdin)

func main() {
	var data []DaftarTugas
	fmt.Print("\n\n==== TASK-MATE ====\n")
	menu(data)
}

func tampilkanData(data []DaftarTugas) {
	var inTD int
	var nama string
	var stat bool

	fmt.Print("\n~~~~~ LIST TUGAS RUMAH ~~~~~~\n")
	fmt.Println("\n====================")
	for i := 0; i < len(data); i++ {
		fmt.Println("Nama					:", data[i].nama)
		fmt.Println("Jenis tugas				:", data[i].jenisTugas)
		fmt.Println("Deskripsi tugas				:", data[i].deskripsiTugas)
		fmt.Println("Tingkat kesulitan			:", data[i].tingkatKesulitan, "/5")
		fmt.Println("Estimasi waktu pengerjaan		:", data[i].estimasiWaktu, "Menit")
		fmt.Println("Kategori ruangan			:", data[i].KategoriRuangan)
		fmt.Print("\n\n====================\n")
	}

	fmt.Println("======= MENU =======")
	fmt.Printf("1. Tambahkan tugas rumah\n2. Urut berdasarkan tingkat kesulitan\n3. Urut berdasarkan estimasi waktu\n4. Hapus pekerjaan\n5. Tugas Selesai\n6. Update pekerjaan\n7. kembali\n\n")
	fmt.Printf("Input = ")
	fmt.Scan(&inTD)
	fmt.Println("")
	switch inTD {
	case 1:
		inputData(&data, stat, nama)
	case 2:
		selectionSort(data)
	case 3:
		insertionSort(data)
	case 4:
		var target string
		fmt.Print("Nama = ")
		fmt.Scan(&nama)
		fmt.Print("Nama tugas yang ingin dihapus = ")
		fmt.Scan(&target)
		hapusData(&data, target, nama, true)
		fmt.Print("\n~~~~~ Tugas telah di hapus ~~~~~\n")
		tampilkanData(data)
	case 5:
		var target, nama string
		var inp int
		// for {
		// 	fmt.Print("Masukkan nama tugas yang telah di selesaikan (ketik stop untuk berhenti) = ")
		// 	fmt.Scan(&target)
		// 	if target == "stop" || target == "Stop" {
		// 		break
		// 	} else {
		// 		tugasSelesai(&data, target)
		// 	}
		// }
		fmt.Print("Nama = ")
		fmt.Scan(&nama)
		fmt.Print("Nama Tugas = ")
		fmt.Scan(&target)
		sequentialSearch(data, target, nama)

		fmt.Print("\n==== MENU ====\n")
		fmt.Print("1. Data tugas yang sudah selesai\n2. Data tugas yang belum selesai\n3. Hasil statistik\n4. Kembali\n\nInput = ")
		fmt.Scan(&inp)
		switch inp {
		case 1:
			tugasSelesai(&data, target)
			menu(data)
		case 2:
			tampilkanData(data)
		case 3:
			tugasSelesai(&data, target)
			statistik(dataSelesai)
		case 4:
			menu(data)
		default:
			fmt.Print("tidak validd!!!")
			menu(data)
		}
	case 6:
		updateData(&data)
		fmt.Print("~~~~~~ Data berhasil di update ~~~~~~\n")
	case 7:
		menu(data)
	default:
		fmt.Println("Input tidak valid!!!")
	}
}

func statistik(data []DaftarTugas) {
	totwak := 0

	for i := 0; i < len(dataSelesai); i++ {
		totwak = totwak + dataSelesai[i].estimasiWaktu
	}

	if len(dataSelesai) <= 0 {
		fmt.Println("Belum ada tugas yang selesai!!")
		menu(data)
	} else {
		fmt.Print("\n~~~~~~~~ RESULT ~~~~~~~~\n\n")
		fmt.Println("Tugas yang sudah selesai sebanyak		:", len(dataSelesai), "tugas")
		fmt.Println("Rata-rata waktu mengerjakan adalah		:", totwak/len(dataSelesai), "menit")
		fmt.Print("\n~~~~~~~~~~~~~~~~~\n\n")
		menu(data)
	}

}

func menu(data []DaftarTugas) {
	var Menu int
	var nama string
	var stat bool
	fmt.Print("\n----- MENU -----")
	fmt.Printf("\n1. Tambah tugas rumah\n2. Tampilkan daftar tugas\n3. Hasil Statistik tugas\n4. Cari\n5. Keluar\n\n")
	fmt.Print("Pilih Menu = ")
	fmt.Scan(&Menu)
	fmt.Println(" ")
	switch Menu {
	case 1:
		nama = ""
		stat = true
		inputData(&data, stat, nama)
	case 2:
		nama = ""
		stat = true
		if len(data) == 0 {
			fmt.Print("Belum ada daftar tugas rumah!!!!\n")
			menu(data)
		} else {
			tampilkanData(data)
		}
	case 3:
		statistik(data)
	case 4:
		var cari int
		fmt.Print("1. Berdasarkan nama dan tugas\n2. berdasarkan ruangan\n\nInput = ")
		fmt.Scan(&cari)
		switch cari {
		case 1:
			var target string
			fmt.Print("Nama = ")
			fmt.Scan(&nama)
			fmt.Print("Tugas = ")
			fmt.Scan(&target)
			i := sequentialSearch(data, target, nama)
			if i == -1 {
				fmt.Println("tugas rumah tidak ditemukan!")
				menu(data)
			} else {
				fmt.Println("Nama					:", data[i].nama)
				fmt.Println("Jenis tugas				:", data[i].jenisTugas)
				fmt.Println("Deskripsi tugas				:", data[i].deskripsiTugas)
				fmt.Println("Tingkat kesulitan			:", data[i].tingkatKesulitan, "/5")
				fmt.Println("Estimasi waktu pengerjaan		:", data[i].estimasiWaktu, "Menit")
				fmt.Println("Kategori ruangan			:", data[i].KategoriRuangan)
				fmt.Print("\n\n====================\n")
				menu(data)
			}
		case 2:
			var target int
			datasortnew := NewSortruangan(data)
			fmt.Print("Input nomor ruangan = ")
			fmt.Scan(&target)
			idx := binarySearch(datasortnew, target)
			if idx == -1 {
				fmt.Println("Data tidak ditemukan")
				menu(data)
				return
			}
			kiri := idx
			for kiri > 0 && datasortnew[kiri-1].KategoriRuangan == target {
				kiri--
			}
			kanan := idx
			for kanan < len(datasortnew)-1 &&
				datasortnew[kanan+1].KategoriRuangan == target {
				kanan++
			}
			for i := kiri; i <= kanan; i++ {
				fmt.Println("Nama\t\t\t\t:", datasortnew[i].nama)
				fmt.Println("Jenis tugas\t\t\t:", datasortnew[i].jenisTugas)
				fmt.Println("Deskripsi tugas\t\t\t:", datasortnew[i].deskripsiTugas)
				fmt.Println("Tingkat kesulitan\t\t:", datasortnew[i].tingkatKesulitan, "/5")
				fmt.Println("Estimasi waktu pengerjaan\t:", datasortnew[i].estimasiWaktu, "Menit")
				fmt.Println("Kategori ruangan\t\t:", datasortnew[i].KategoriRuangan)
				fmt.Println("==============================")
			}
		}
		menu(data)
	case 5:
		fmt.Print("\n===== Terima Kasih =====\n\n")

	default:
		fmt.Print("Input tidak sesuai, mohon input kembali!!!")
	}
}

func inputData(data *[]DaftarTugas, stat bool, nama string) {
	var Tugas DaftarTugas
	var quest int
	if stat {

		reader.ReadString('\n')
		fmt.Print("Masukkan nama = ")
		nama, _ = reader.ReadString('\n')
		nama = strings.TrimSpace(nama)
		Tugas.nama = nama

		fmt.Print("Masukkan jenis tugas = ")
		fmt.Scan(&Tugas.jenisTugas)

		reader.ReadString('\n')
		fmt.Print("Deskripsi Tugas = ")
		Tugas.deskripsiTugas, _ = reader.ReadString('\n')
		Tugas.deskripsiTugas = strings.TrimSpace(Tugas.deskripsiTugas)

		fmt.Print("Kategori Ruangan = ")
		fmt.Scan(&Tugas.KategoriRuangan)

		for {
			fmt.Print("Tingkat kesulitan 1-5 (Mudah - Sulit) = ")
			fmt.Scan(&Tugas.tingkatKesulitan)
			if Tugas.tingkatKesulitan <= 5 && Tugas.tingkatKesulitan > 0 {
				break
			} else {
				fmt.Println("Hanya menerima input dari 1-5, mohon input ulangi")
			}
		}

		fmt.Print("Estimasi waktu dalam menit = ")
		fmt.Scan(&Tugas.estimasiWaktu)

		*data = append(*data, Tugas)

		fmt.Println("\n~~~~ Berhasil ditambahkan ~~~~")
	} else {
		Tugas.nama = nama
		reader.ReadString('\n')
		fmt.Print("Masukkan jenis tugas = ")
		fmt.Scan(&Tugas.jenisTugas)

		reader.ReadString('\n')
		fmt.Print("Deskripsi Tugas = ")
		Tugas.deskripsiTugas, _ = reader.ReadString('\n')
		Tugas.deskripsiTugas = strings.TrimSpace(Tugas.deskripsiTugas)

		fmt.Print("Kategori Ruangan = ")
		fmt.Scan(&Tugas.KategoriRuangan)
		for {
			fmt.Print("Tingkat kesulitan 1-5 (Mudah - Sulit) = ")
			fmt.Scan(&Tugas.tingkatKesulitan)
			if Tugas.tingkatKesulitan <= 5 && Tugas.tingkatKesulitan > 0 {
				break
			} else {
				fmt.Println("Hanya menerima input dari 1-5, mohon input ulangi")
			}
		}
		fmt.Print("Estimasi waktu dalam menit = ")
		fmt.Scan(&Tugas.estimasiWaktu)

		*data = append(*data, Tugas)
		fmt.Print("\n~~~~ Berhasil ditambahkan ~~~~\n")
	}
	fmt.Print("\n----- MENU -----\n1. Tambahkan tugas rumah\n2. Ganti nama\n3. Tampilkan daftar pekerjaan\n4. Kembali\n\n")
	fmt.Print("Input = ")

	for {
		fmt.Scan(&quest)
		fmt.Println(" ")
		if quest == 1 {
			stat = false
			inputData(data, stat, nama)
			break
		} else if quest == 2 {
			stat = true
			nama = ""
			inputData(data, stat, nama)
			break
		} else if quest == 3 {
			tampilkanData(*data)
		} else if quest == 4 {
			menu(*data)
			break
		} else {
			fmt.Print("Input tidak sesuai, mohon input kembali!!!\n\n")
			fmt.Println("Input = ")
			fmt.Scan(&quest)
			fmt.Println(" ")
		}
	}
}

func hapusData(data *[]DaftarTugas, target, nama string, stat bool) []DaftarTugas {
	// if target == "stop" || target == "Stop" {
	// 	return *data
	// }
	var idx int
	idx = sequentialSearch(*data, target, nama)

	if idx != -1 {
		*data = append((*data)[:idx], (*data)[idx+1:]...)
	}
	return *data
}

func tugasSelesai(data *[]DaftarTugas, target string) {
	var nama string
	idx := sequentialSearch(*data, target, nama)
	if idx != -1 {
		dataSelesai = append(dataSelesai, (*data)[idx])
		hapusData(data, target, nama, true)
	}
	if len(dataSelesai) > 0 {
		fmt.Print("\n~~~~~~ Tugas yang sudah selesai ~~~~~~\n")
		fmt.Println("\n====================")
		for i := 0; i < len(dataSelesai); i++ {
			fmt.Println("Nama					:", dataSelesai[i].nama)
			fmt.Println("Jenis tugas				:", dataSelesai[i].jenisTugas)
			fmt.Println("Deskripsi tugas				:", dataSelesai[i].deskripsiTugas)
			fmt.Println("Tingkat kesulitan			:", dataSelesai[i].tingkatKesulitan, "/5")
			fmt.Println("Estimasi waktu pengerjaan		:", dataSelesai[i].estimasiWaktu, "Menit")
			fmt.Println("Kategori ruangan			:", dataSelesai[i].KategoriRuangan)
			fmt.Print("\n\n====================\n")
		}
	}
}

func updateData(data *[]DaftarTugas) {
	var target, nama string
	fmt.Print("Nama = ")
	fmt.Scan(&nama)
	fmt.Print("Nama tugas yang ingin di update = ")
	fmt.Scan(&target)
	idx := sequentialSearch(*data, target, nama)

	if idx == -1 {
		fmt.Print("Data tidak ditemukan!!!\n")
		updateData(data)
	} else {
		var stat string
		for i := 0; i < len(*data); i++ {
			if i == idx {
				idx = i

				fmt.Print("Apakah ingin ubah nama? (y/n)")
				fmt.Scan(&stat)
				if stat == "y" || stat == "Y" {
					fmt.Print("Masukkan nama baru = ")
					fmt.Scan(&(*data)[idx].nama)
				}

				reader.ReadString('\n')
				fmt.Print("Masukkan tugas baru = ")
				(*data)[idx].jenisTugas, _ = reader.ReadString('\n')
				(*data)[idx].jenisTugas = strings.TrimSpace((*data)[idx].jenisTugas)

				fmt.Print("Deskripsi Tugas = ")
				(*data)[idx].deskripsiTugas, _ = reader.ReadString('\n')
				(*data)[idx].deskripsiTugas = strings.TrimSpace((*data)[idx].deskripsiTugas)

				fmt.Print("Kateori Ruangan = ")
				fmt.Scan(&(*data)[idx].KategoriRuangan)

				fmt.Print("Tingkat kesulitan 1-5 = ")
				fmt.Scan(&(*data)[idx].tingkatKesulitan)

				fmt.Print("Estimasi waktu dalam menit = ")
				fmt.Scan(&(*data)[idx].estimasiWaktu)

				fmt.Print("\n~~~~~~ Tugas berhasil di update ~~~~~~\n")
				tampilkanData(*data)
			}
		}
	}

}

func binarySearch(data []DaftarTugas, target int) int {
	var mid int
	n := len(data)
	low := 0
	high := n - 1

	for low <= high {
		mid = low + (high-low)/2
		if data[mid].KategoriRuangan == target {
			return mid
		} else if data[mid].KategoriRuangan < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func sequentialSearch(data []DaftarTugas, target, nama string) int {
	n := len(data)
	hasil := -1

	for i := 0; i < n; i++ {
		if data[i].jenisTugas == target && data[i].nama == nama {
			hasil = i
		}
	}
	return hasil
}

func NewSortruangan(data []DaftarTugas) []DaftarTugas {
	n := len(data)
	for i := 0; i < n-1; i++ {
		idx_min := i
		for j := i + 1; j < n; j++ {
			if data[j].KategoriRuangan < data[idx_min].KategoriRuangan {
				idx_min = j
			}
		}
		data[i], data[idx_min] = data[idx_min], data[i]
	}
	return data
}

func selectionSort(data []DaftarTugas) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		idx_min := i
		for j := i + 1; j < n; j++ {
			if data[j].tingkatKesulitan < data[idx_min].tingkatKesulitan {
				idx_min = j
			}
		}
		data[i], data[idx_min] = data[idx_min], data[i]
	}

	fmt.Println("\n~~~~ Sort berdasarkan tingkat kesulitan dari mudah ke tersulit ~~~~")
	tampilkanData(data)
}

func insertionSort(data []DaftarTugas) {
	n := len(data)
	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].estimasiWaktu > key.estimasiWaktu {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
	println("\n~~~~ Urutan berdasar estimasi waktu selesai tercepat ~~~~")
	tampilkanData(data)
}
