package main

import (
	"fmt"
	"math"
)

type login struct {
	user, passUser   string
	admin, passAdmin string
}
type arrlogin [21]login

func daftarLogin(T *arrlogin) {
	//ADMIN
	T[0].admin = "kanjeng"
	T[0].passAdmin = "bismillah"
	T[1].admin = "fadjar"
	T[1].passAdmin = "fadjar123"
	//USERSw
	T[0].user = "user"
	T[0].passUser = "user"
	T[1].user = "Yan"
	T[1].passUser = "123"
}
func signinUser(T arrlogin, user, pass string) int {
	for i := 0; i < 21; i++ {
		if T[i].user == user && T[i].passUser == pass {
			return i
		}
	}
	return -1
}
func signinAdmin(T arrlogin, user, pass string) int {
	for i := 0; i < 21; i++ {
		if T[i].admin == user && T[i].passAdmin == pass {
			return i
		}
	}
	return -1
}

func userUser(T arrlogin, user, pass string) int {
	fmt.Print("Username : ")
	fmt.Scan(&user)
	fmt.Print("Password : ")
	fmt.Scan(&pass)
	found := signinUser(T, user, pass)
	for found == -1 {
		fmt.Println("--------------------------------------")
		fmt.Println("Username/Password tidak ditemukan")
		fmt.Print("Username : ")
		fmt.Scan(&user)
		fmt.Print("Password : ")
		fmt.Scan(&pass)
		found = signinUser(T, user, pass)
	}
	return found
}
func userAdmin(T arrlogin, user, pass string) int {
	fmt.Print("Username : ")
	fmt.Scan(&user)
	fmt.Print("Password : ")
	fmt.Scan(&pass)
	found := signinAdmin(T, user, pass)
	for found == -1 {
		fmt.Println("--------------------------------------")
		fmt.Println("Username/Password tidak tersedia")
		fmt.Print("Username : ")
		fmt.Scan(&user)
		fmt.Print("Password : ")
		fmt.Scan(&pass)
		found = signinAdmin(T, user, pass)
	}
	return found
}

type pariwisata struct {
	name, katakunci, kategori string
	hargatiket                int
	pusat, jarak              float64
	fitur                     [100]fiture
	tfasilitas, twahan        int
	wahan                     [100]wahana
}
type fiture struct {
	name  string
	biaya int
}
type wahana struct {
	name      string
	deskripsi [100]string
}

type arrPariwisata [100]pariwisata

const wisata int = 3

// ADMIN
func TempatWisata(T *arrPariwisata) {
	T[0].kategori = "Hiburan"
	T[0].katakunci = "JatimPark1"
	T[0].name = "JawaTimurPark1"
	T[0].hargatiket = 50000
	T[0].pusat = 10

	T[1].kategori = "Satwa"
	T[1].katakunci = "JatimPark2"
	T[1].name = "JawaTimurPark2"
	T[1].hargatiket = 60000
	T[1].pusat = 20

	T[2].kategori = "Sejarah"
	T[2].katakunci = "JatimPark3"
	T[2].name = "JawaTimurPark3"
	T[2].hargatiket = 70000
	T[2].pusat = 30
}

func Caristring(T arrPariwisata, n int, x string) int {
	var found int = -1
	i := 0
	for i < n && found == -1 {
		if x == T[i].name || x == T[i].katakunci {
			found = i
		}
		i++
	}
	return found
}

func tambahWahan(T *arrPariwisata, d *int, X string) {
	var input wahana
	var i int = Caristring(*T, wisata, X)
	fmt.Println("--------------------------------------------")
	fmt.Println("              MENAMBAH WAHANA               ")
	fmt.Println("--------------------------------------------")
	fmt.Println("Ketik (Nama Wahana) STOP bila berhenti")
	fmt.Scan(&input.name)
	for input.name != "STOP" {
		T[i].wahan[T[i].twahan].name = input.name
		deskripsi(T, d, i)
		T[i].twahan++
		fmt.Println("--------------------------------------------")
		fmt.Println("Ketik (Nama Wahana) STOP bila berhenti")
		fmt.Scan(&input.name)
	}
	fmt.Println("--------------------------------------------")
}
func deskripsi(T *arrPariwisata, d *int, i int) {
	var input string
	*d = 0
	fmt.Println("Menambahkan Kolom DESKRIPSI :")
	fmt.Println("Ketik # untuk selesai ")
	fmt.Scan(&input)
	for input != "#" {
		T[i].wahan[T[i].twahan].deskripsi[*d] = input
		*d++
		fmt.Scan(&input)
	}
}

func tambahFiture(T *arrPariwisata, X string) {
	var input fiture
	var i int = Caristring(*T, wisata, X)
	fmt.Println("--------------------------------------------")
	fmt.Println("             MENAMBAH FASILITAS             ")
	fmt.Println("--------------------------------------------")
	fmt.Println("Ketik (Nama Fasilitas) STOP bila berhenti")
	fmt.Scan(&input.name)
	for input.name != "STOP" {
		fmt.Scan(&input.biaya)
		T[i].fitur[T[i].tfasilitas] = input
		T[i].tfasilitas++
		fmt.Println("--------------------------------------------")
		fmt.Println("Ketik (Nama Fasilitas) STOP bila berhenti")
		fmt.Scan(&input.name)
	}
	fmt.Println("--------------------------------------------")
}

func renameWahana(T *arrPariwisata, X string) {
	fmt.Println("--------------------------------------------")
	fmt.Println("               RENAME WAHANA                ")
	fmt.Println("--------------------------------------------")
	var cari2 string
	W := Caristring(*T, wisata, X)
	fmt.Print("Nama Wahana : ")
	fmt.Scan(&cari2)
	var found int = -1
	i := 0
	for i < T[W].twahan && found == -1 {
		if cari2 == T[W].wahan[i].name {
			found = i
		}
		i++
	}
	if found == -1 {
		fmt.Println("Wahana tidak ditemukan ")
	}
	fmt.Println("Rename dengan : namawahana")
	fmt.Scan(&T[W].wahan[found].name)
}

func renameFiture(T *arrPariwisata, X string) {
	fmt.Println("RENAME FASILITAS")
	var cari2 string
	W := Caristring(*T, wisata, X)
	fmt.Println("Nama FASILITAS : ")
	fmt.Scan(&cari2)
	var found int = -1
	i := 0
	for i < T[W].tfasilitas && found == -1 {
		if cari2 == T[W].fitur[i].name {
			found = i
		}
		i++
	}
	if found == -1 {
		fmt.Println("FASILITAS tidak ditemukan ")
	}
	fmt.Println("Rename dengan : namafasilitas <space> biaya")
	fmt.Scan(&T[W].fitur[found].name, &T[W].fitur[found].biaya)
}

func deleteWahana(T *arrPariwisata, X string) {
	fmt.Println("DELETE WAHANA")
	var cari2 string
	var i int = 0
	W := Caristring(*T, wisata, X)
	fmt.Print("Nama Wahana yang di HAPUS :")
	fmt.Scan(&cari2)
	var found int = -1
	for i < T[W].twahan && found == -1 {
		if cari2 == T[W].wahan[i].name {
			found = i
		}
		i++
	}
	if W == -1 && found == -1 {
		fmt.Println("Wahana tidak ditemukan")
	} else {
		i = found
		for i <= T[W].twahan-2 {
			T[W].wahan[i] = T[W].wahan[i+1]
			i++
		}
		T[W].twahan--
	}
}

func deleteFitur(T *arrPariwisata, X string) {
	fmt.Println("DELETE FASILITAS")
	var cari2 string
	var i int = 0
	W := Caristring(*T, wisata, X)
	fmt.Println("Nama FASILITAS yang di HAPUS")
	fmt.Scan(&cari2)
	var found int = -1
	for i < T[W].tfasilitas && found == -1 {
		if cari2 == T[W].fitur[i].name {
			found = i
		}
		i++
	}
	if W == -1 && found == -1 {
		fmt.Println("FASILITAS tidak ditemukan ")
	} else {
		i = found
		for i <= T[W].tfasilitas-2 {
			T[W].fitur[i] = T[W].fitur[i+1]
			i++
		}
		T[W].tfasilitas--
	}
}

func edit(T *arrPariwisata, d *int) {
	TempatWisata(T)
	var cari, status1, status2 string
	fmt.Println("Nama Tempat Wisata yang akan diedit :")
	fmt.Scan(&cari)
	fmt.Println()
	fmt.Println("--------------------------------------------")
	fmt.Println("Lakukan Tambah ke :")
	fmt.Println("1. Wahana")
	fmt.Println("2. Fasilitas")
	fmt.Println("LANJUT")
	fmt.Print("JAWABANMU : ")
	fmt.Scan(&status1)
	for status1 != "LANJUT" {
		if status1 == "1" || status1 == "Wahana" {
			tambahWahan(T, d, cari)
		} else if status1 == "2" || status1 == "Fasilitas" {
			tambahFiture(T, cari)
		}
		fmt.Println()
		fmt.Println("--------------------------------------------")
		fmt.Println("Lakukan Tambah ke :")
		fmt.Println("1. Wahana")
		fmt.Println("2. Fasilitas")
		fmt.Println("LANJUT")
		fmt.Print("JAWABANMU : ")
		fmt.Scan(&status1)
	}
	fmt.Println()

	print(*T, *d, cari)

	fmt.Println()
	fmt.Println("--------------------------------------------")
	fmt.Println("Pilih akan melakukan edit :")
	fmt.Println("1. Rename Wahana")
	fmt.Println("2. DELETE Wahana")
	fmt.Println("3. Rename Fitur")
	fmt.Println("4. DELETE Fitur")
	fmt.Println("LANJUT")
	fmt.Print("JAWABANMU : ")
	fmt.Scan(&status2)
	fmt.Println()
	for status2 != "LANJUT" {
		if status2 == "1" || status2 == "Rename Wahana" {
			renameWahana(T, cari)
		} else if status2 == "2" || status2 == "DELETE Wahana" {
			deleteWahana(T, cari)
		} else if status2 == "3" || status2 == "Rename Fasilita" {
			renameFiture(T, cari)
		} else if status2 == "4" || status2 == "DELETE Fasilita" {
			deleteFitur(T, cari)
		}
		fmt.Println()
		fmt.Println("--------------------------------------------")
		fmt.Println("Pilih akan melakukan edit :")
		fmt.Println("1. Rename Wahana")
		fmt.Println("2. DELETE Wahana")
		fmt.Println("3. Rename Fasilitas")
		fmt.Println("4. DELETE Fasilitas")
		fmt.Println("LANJUT")
		fmt.Print("JAWABANMU : ")
		fmt.Scan(&status2)
	}
	print(*T, *d, cari)
}

func print(T arrPariwisata, d int, X string) {
	idx := Caristring(T, wisata, X)
	fmt.Println("--------------------------------------------")
	fmt.Println("Wahana di", T[idx].name)
	fmt.Println()
	for i := 0; i < T[idx].twahan; i++ {
		fmt.Println(T[idx].wahan[i].name)
		for j := 0; j < d; j += 4 {
			fmt.Printf("%s %s %s %s \n", T[idx].wahan[i+0].deskripsi[j], T[idx].wahan[i].deskripsi[j+1], T[idx].wahan[i].deskripsi[j+2], T[idx].wahan[i].deskripsi[j+3])
		}
		fmt.Println()
	}
	fmt.Println("Fasilitas di", T[idx].name)
	for j := 0; j < T[idx].tfasilitas; j++ {
		fmt.Println(T[idx].fitur[j].name, T[idx].fitur[j].biaya)
	}
	fmt.Println("--------------------------------------------")
}

func printAdmin(T arrPariwisata) {
	fmt.Println("--------------------------------------------")
	for i := 0; i < wisata; i++ {
		fmt.Println("--------------------------------------------")
		fmt.Printf("%s		 %s \n", T[i].katakunci, T[i].kategori)
		fmt.Printf("%s \n", T[i].name)
		fmt.Printf("Harga Tiket : %d \n", T[i].hargatiket)
		fmt.Printf("Jarak Dengan Kota Malang : %0.2f km\n", T[i].pusat)
	}
	fmt.Println("--------------------------------------------")
	fmt.Println("--------------------------------------------")
}

// User
func ukurJarak(T *arrPariwisata) {
	var km float64
	fmt.Print("Jarak Anda ke Kota Malang (km) : ")
	fmt.Scan(&km)
	for i := 0; i < wisata; i++ {
		T[i].jarak = math.Sqrt(math.Pow(km, 2) + math.Pow(T[i].pusat, 2))
	}
}

func printUser(T arrPariwisata) {
	fmt.Println("--------------------------------------------")
	for i := 0; i < wisata; i++ {
		fmt.Println("--------------------------------------------")
		fmt.Printf("%s		 %s \n", T[i].katakunci, T[i].kategori)
		fmt.Printf("%s \n", T[i].name)
		fmt.Printf("Harga Tiket : %d \n", T[i].hargatiket)
		fmt.Printf("Jarak Yang Akan Ditempuh : %0.2f km\n", T[i].jarak)
	}
	fmt.Println("--------------------------------------------")
	fmt.Println("--------------------------------------------")
}

func printChose(T arrPariwisata, d int) {
	var pilih string
	fmt.Print("Pilih Nama Wisata Keinginanmu :")
	fmt.Scan(&pilih)
	idx := Caristring(T, wisata, pilih)
	for idx == -1 {
		fmt.Println()
		fmt.Println("Nama Termpat Wisata Tidak Ditemukan")
		fmt.Print("Pilih Nama Wisata Keinginanmu :")
		fmt.Scan(&pilih)
		idx = Caristring(T, wisata, pilih)
	}
	fmt.Println()
	fmt.Println("--------------------------------------------")
	fmt.Printf("Wisata di %s(%s)\n", T[idx].katakunci, T[idx].name)
	fmt.Println("Dengan Kategori :", T[idx].kategori)
	fmt.Println("Memiliki ", T[idx].twahan, " Wahana")
	fmt.Println()
	for i := 0; i < T[idx].twahan; i++ {
		fmt.Println(T[idx].wahan[i].name)
		for j := 0; j < d; j += 4 {
			fmt.Printf("%s %s %s %s \n", T[idx].wahan[i+0].deskripsi[j], T[idx].wahan[i].deskripsi[j+1], T[idx].wahan[i].deskripsi[j+2], T[idx].wahan[i].deskripsi[j+3])
		}
		fmt.Println()
	}
	fmt.Println("Memiliki ", T[idx].tfasilitas, " Fasilitas")
	fmt.Println("dengan harga yang beragam")
	for j := 0; j < T[idx].tfasilitas; j++ {
		fmt.Println(T[idx].fitur[j].name, T[idx].fitur[j].biaya)
	}
	fmt.Println("--------------------------------------------")
}

func printFitur(T arrPariwisata) {
	fmt.Println("CARI DENGAN FASILITAS : ")
	var cari1 string
	fmt.Scan(&cari1)
	fmt.Println("Wisata yang memiliki FASILITAS ", cari1)
	var found int = -1
	fmt.Println("--------------------------------------------")
	for i := 0; i < wisata; i++ {
		for j := 0; j < T[i].tfasilitas; j++ {
			if cari1 == T[i].fitur[j].name {
				fmt.Println("--------------------------------------------")
				fmt.Printf("%s		 %s \n", T[i].katakunci, T[i].kategori)
				fmt.Printf("%s \n", T[i].name)
				fmt.Printf("Harga Tiket : %d \n", T[i].hargatiket)
				fmt.Printf("Jarak Yang Akan Ditempuh : %0.2f km\n", T[i].jarak)
				found = j
			}
		}
	}
	for found == -1 {
		fmt.Println("FASILITAS tidak ditemukan")
		fmt.Print("Input Lagi : ")
		fmt.Scan(&cari1)
		fmt.Println("--------------------------------------------")
		for i := 0; i < wisata; i++ {
			for j := 0; j < T[i].tfasilitas; j++ {
				if cari1 == T[i].fitur[j].name {
					fmt.Println("--------------------------------------------")
					fmt.Printf("%s		 %s \n", T[i].katakunci, T[i].kategori)
					fmt.Printf("%s \n", T[i].name)
					fmt.Printf("Harga Tiket : %d \n", T[i].hargatiket)
					fmt.Printf("Jarak Yang Akan Ditempuh : %0.2f km\n", T[i].jarak)
					found = j
				}
			}
		}
	}
	fmt.Println("--------------------------------------------")
	fmt.Println("--------------------------------------------")
}

func printKategori(T arrPariwisata, n int) {
	var stat string
	fmt.Println()
	fmt.Println("Berdasarkan kategori seperti")
	fmt.Println("Hiburan / Satwa / Sejarah")
	fmt.Print("JAWABANMU :")
	fmt.Scan(&stat)
	fmt.Println("--------------------------------------------")
	var found int = -1
	for i := 0; i < n; i++ {
		if stat == T[i].kategori {
			fmt.Println("--------------------------------------------")
			fmt.Printf("%s		 %s \n", T[i].katakunci, T[i].kategori)
			fmt.Printf("%s \n", T[i].name)
			fmt.Printf("Harga Tiket : %d \n", T[i].hargatiket)
			fmt.Printf("Jarak Yang Akan Ditempuh : %0.2f km\n", T[i].jarak)
			found = i
		}
	}
	for found == -1 {
		fmt.Println("Kategori tidak ditemukan")
		fmt.Print("Input Lagi : ")
		fmt.Scan(&stat)
		fmt.Println("--------------------------------------------")
		for i := 0; i < n; i++ {
			if stat == T[i].kategori {
				fmt.Println("--------------------------------------------")
				fmt.Printf("%s		 %s \n", T[i].katakunci, T[i].kategori)
				fmt.Printf("%s \n", T[i].name)
				fmt.Printf("Harga Tiket : %d \n", T[i].hargatiket)
				fmt.Printf("Jarak Yang Akan Ditempuh : %0.2f km\n", T[i].jarak)
				found = i
			}
		}
	}
	fmt.Println("--------------------------------------------")
	fmt.Println("--------------------------------------------")
}

func jarak(T *arrPariwisata, n int) {
	var X string
	var pass, i int
	var temp pariwisata
	fmt.Scan(&X)
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = T[pass]
		if X == "Tertinggi" {
			for i > 0 && temp.jarak > T[i-1].jarak {
				T[i] = T[i-1]
				i--
			}
		} else if X == "Terendah" {
			for i > 0 && temp.jarak < T[i-1].jarak {
				T[i] = T[i-1]
				i--
			}
		}
		T[i] = temp
		pass++
	}
}
func printJarak(T arrPariwisata, n int, X string) {
	fmt.Println("Jarak", X)
	fmt.Println("--------------------------------------------")
	for i := 0; i < wisata; i++ {
		fmt.Println("--------------------------------------------")
		fmt.Printf("%s		 %s \n", T[i].katakunci, T[i].kategori)
		fmt.Printf("%s \n", T[i].name)
		fmt.Printf("Harga Tiket : %d \n", T[i].hargatiket)
		fmt.Printf("Jarak Yang Akan Ditempuh : %0.2f km\n", T[i].jarak)
	}
	fmt.Println("--------------------------------------------")
	fmt.Println("--------------------------------------------")
}

func Tiket(T *arrPariwisata, n int) {
	var X string
	var pass, i int
	var temp pariwisata
	fmt.Scan(&X)
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = T[pass]
		if X == "Termahal" {
			for i > 0 && temp.jarak > T[i-1].jarak {
				T[i] = T[i-1]
				i--
			}
		} else if X == "Termurah" {
			for i > 0 && temp.jarak < T[i-1].jarak {
				T[i] = T[i-1]
				i--
			}
		}
		T[i] = temp
		pass++
	}
}
func printTiket(T arrPariwisata, n int, X string) {
	fmt.Println("Harga Tiket", X)
	fmt.Println("--------------------------------------------")
	for i := 0; i < wisata; i++ {
		fmt.Println("--------------------------------------------")
		fmt.Printf("%s		 %s \n", T[i].katakunci, T[i].kategori)
		fmt.Printf("%s \n", T[i].name)
		fmt.Printf("Harga Tiket : %d \n", T[i].hargatiket)
		fmt.Printf("Jarak Yang Akan Ditempuh : %0.2f km\n", T[i].jarak)
	}
	fmt.Println("--------------------------------------------")
	fmt.Println("--------------------------------------------")
}

func Template(x string) {
	if x == "utama" {
		fmt.Println("--------------------------------------------")
		fmt.Println("***      SELAMAT DATANG DITRAVELKUY      ***")
		fmt.Println("***     TRAVELKUY MERUPAKAN APLIKASI     ***")
		fmt.Println("***      PARIWISATA UNTUK MENEMUKAN      ***")
		fmt.Println("***           WISATA LIBURANMU           ***")
		fmt.Println("--------------------------------------------")
	} else if x == "exit" {
		fmt.Println("               TERIMA KASIH                 ")
	}
}

func main() {
	var tab arrlogin
	var T arrPariwisata
	var d int
	var user, pass, user1, pass1, num, inapk, inapk1 string
	daftarLogin(&tab)
	TempatWisata(&T)
	Template("utama")
	fmt.Println("***            HALAMAN LOGIN             ***")
	fmt.Println("--------------------------------------------")
	fmt.Println("Masuk sebagai ADMIN/USER ? ")
	fmt.Println("1. ADMIN")
	fmt.Println("2. USER")
	fmt.Println("EXIT")
	fmt.Print("LOGIN SEBAGAI : ")
	fmt.Scan(&num)
	fmt.Println()

	for num != "EXIT" {
		if num == "1" || num == "ADMIN" {
			fmt.Println("--------------------------------------------")
			fmt.Println("***      SELAMAT DATANG DITRAVELKUY      ***")
			fmt.Println("***         LOGIN SEBAGAI ADMIN          ***")
			fmt.Println("--------------------------------------------")
			idx := userAdmin(tab, user, pass)
			fmt.Println("--------------------------------------------")
			for inapk != "LOGOUT" {
				fmt.Println()
				fmt.Println("SELAMAT DATANG ", tab[idx].admin)
				printAdmin(T)
				fmt.Println()
				fmt.Println("INGIN MELAKUKAN EDITING? YA/TIDAK")
				fmt.Println("PRINT")
				fmt.Println("LOGOUT")
				fmt.Print("Jawabanmu : ")
				fmt.Scan(&inapk)
				fmt.Println()
				if inapk == "YA" || inapk == "EDITING" {
					edit(&T, &d)
					fmt.Print("Pilih HOME/LOGOUT : ")
					fmt.Scan(&inapk)
				} else if inapk == "TIDAK" || inapk == "PRINT" {
					printAdmin(T)
					inapk = "LOGOUT"
				}
			}

		}
		if num == "2" || num == "USER" {
			fmt.Println("--------------------------------------------")
			fmt.Println("***      SELAMAT DATANG DITRAVELKUY      ***")
			fmt.Println("***         LOGIN SEBAGAI USER           ***")
			fmt.Println("--------------------------------------------")
			idx1 := userUser(tab, user1, pass1)
			fmt.Println("--------------------------------------------")
			for inapk1 != "LOGOUT" {
				fmt.Println()
				fmt.Println("SELAMAT DATANG ", tab[idx1].user)
				ukurJarak(&T)
				fmt.Println()
				printUser(T)
				fmt.Println()
				fmt.Println("INGIN MELAKUKAN APA")
				fmt.Println("1. Pilih Wisata (Berdasarkan NAMA)")
				fmt.Println("2. Lakukan FILTER")
				fmt.Println("LOGOUT")
				fmt.Print("JAWABANMU : ")
				fmt.Scan(&inapk1)
				fmt.Println()
				if inapk1 == "1" {
					printChose(T, d)
					fmt.Print("Pilih HOME/LOGOUT :")
					fmt.Scan(&inapk1)
				} else if inapk1 == "2" {
					fmt.Println("Lakukan FILTER Berdasarkan")
					fmt.Println("1. KATEGORI")
					fmt.Println("2. FASILITAS")
					fmt.Println("3. JARAK")
					fmt.Println("4. HARGA TIKET")
					fmt.Println("LOGOUT")
					fmt.Print("JAWABANMU : ")
					fmt.Scan(&inapk1)
					if inapk1 == "1" || inapk1 == "KATEGORI" {
						printKategori(T, wisata)
						printChose(T, d)
						fmt.Print("Pilih HOME/LOGOUT : ")
						fmt.Scan(&inapk1)
					} else if inapk1 == "2" || inapk1 == "FASILITAS" {
						printFitur(T)
						printChose(T, d)
						fmt.Print("Pilih HOME/LOGOUT : ")
						fmt.Scan(&inapk1)
					} else if inapk1 == "3" || inapk1 == "JARAK" {
						fmt.Println("Melakukan Filter JARAK Tertinggi/Terendah")
						jarak(&T, wisata)
						printJarak(T, wisata, inapk1)
						fmt.Print("Pilih HOME/LOGOUT :")
						fmt.Scan(&inapk1)
					} else if inapk1 == "4" || inapk1 == "HARGA TIKET" {
						fmt.Println("Melakukan Filter HARGA TIKET Tertinggi/Terendah")
						Tiket(&T, wisata)
						printTiket(T, wisata, inapk1)
						fmt.Print("Pilih HOME/LOGOUT :")
						fmt.Scan(&inapk1)
					}
				}
			}
		}
		fmt.Println()
		fmt.Println("--------------------------------------------")
		fmt.Println("***            HALAMAN LOGIN             ***")
		fmt.Println("--------------------------------------------")
		fmt.Println("Masuk sebagai ADMIN/USER ? ")
		fmt.Println("1. ADMIN")
		fmt.Println("2. USER")
		fmt.Println("EXIT")
		fmt.Print("LOGIN SEBAGAI : ")
		fmt.Scan(&num)
		fmt.Println("--------------------------------------------")
		fmt.Println()
	}
	Template("exit")
}
