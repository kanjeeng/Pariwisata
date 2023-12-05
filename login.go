package main

import "fmt"

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
func signup(T *arrlogin, i *int) {
	var user, pass string
	fmt.Println("Sign Up")
	fmt.Println("Username :")
	fmt.Scan(&user)
	fmt.Println("Password :")
	fmt.Scan(&pass)
	var found int = signinUser(*T, user, pass)
	if found == -1 {
		for i := 0; i < 100; i++ {
			if T[i].user == "" && T[i].passUser == "" {
				T[i].user = user
				T[i].passUser = pass
			}
		}
	}
}

func userUser(T arrlogin, user, pass string) int {
	fmt.Println("LOGIN SEBAGAI USER")
	fmt.Println("Username")
	fmt.Scan(&user)
	fmt.Println("Password")
	fmt.Scan(&pass)
	found := signinUser(T, user, pass)
	for found == -1 {
		fmt.Println("Username/Password tidak ditemukan")
		fmt.Println("Username")
		fmt.Scan(&user)
		fmt.Println("Password")
		fmt.Scan(&pass)
		found = signinUser(T, user, pass)
	}
	return found
}
func userAdmin(T arrlogin, user, pass string) int {
	fmt.Println("LOGIN SEBAGAI ADMIN")
	fmt.Println("Username")
	fmt.Scan(&user)
	fmt.Println("Password")
	fmt.Scan(&pass)
	found := signinAdmin(T, user, pass)
	for found == -1 {
		fmt.Println("Username/Password tidak ditemukan")
		fmt.Println("Username")
		fmt.Scan(&user)
		fmt.Println("Password")
		fmt.Scan(&pass)
		found = signinAdmin(T, user, pass)
	}
	return found
}

func main() {
	var tab arrlogin
	var user, pass, num, inapk string
	daftarLogin(&tab)
	fmt.Println("Selamat datang TRAVELKUY")

	fmt.Println("1. ADMIN")
	fmt.Println("2. USER")
	fmt.Println("EXIT")
	fmt.Print("LOGIN SEBAGAI : ")
	fmt.Scan(&num)
	for num != "EXIT" {
		if num == "1" || num == "ADMIN" {
			for inapk != "LOGOUT" {
				idx := userAdmin(tab, user, pass)
				fmt.Println("Selamat Datang", tab[idx].admin)

			}

		} else {
			idx := userUser(tab, user, pass)
			fmt.Println("Selamat Datang", tab[idx].user)
		}
	}

}
