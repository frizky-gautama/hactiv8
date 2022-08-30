package main

import (
	"fmt"
	"hactiv8/Service"
)

type Person struct {
	Name      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	var db []*Service.User
	serviceUser := Service.HandlerFunc(db)
	temanHactive := []*Service.User{
		{Nama: "Sigit Setiawan", Age: "37"},
		{Nama: "Yosef Brahmantyo", Age: "42"},
	}

	for _, val := range temanHactive {
		insert := serviceUser.CreateUser(val)
		fmt.Println(insert)
	}

	dataShow := serviceUser.GetAllUser()
	fmt.Println("=== Show Users ===")
	for _, val := range dataShow {
		printUser(val.Nama)
	}

	// tugas1()
	// tugas2()
	// tugas3()
	// args, _ := strconv.Atoi(os.Args[1])
	// tugas_CLI(args)
}

func printUser(nama string) {
	fmt.Println(nama)
}

// func tugas_CLI(os int) {
// 	temanHactive := []*Person{
// 		{Name: "Sigit Setiawan", Alamat: "Jakarta", Pekerjaan: "IT Security", Alasan: "Gabut"},
// 		{Name: "Yosef Brahmantyo", Alamat: "Surabaya", Pekerjaan: "Middleware Programmer", Alasan: "Iseng"},
// 	}

// 	loop := func(persons []*Person, os int) {
// 		for key, i := range persons {
// 			if key == os {
// 				fmt.Println("Nama: "+i.Name+",", "Alamat: "+i.Alamat+",", "Pekerjaan: "+i.Pekerjaan+",", "Alasan: "+i.Alasan)
// 			}
// 		}
// 	}
// 	loop(temanHactive, os)
// }

// func tugas_closure() {
// 	temanHactive := []*Person{
// 		{Name: "Sigit Setiawan", Gender: "Pria"},
// 		{Name: "Yosef Brahmantyo", Gender: "Pria"},
// 		{Name: "Bayu", Gender: "Pria"},
// 		{Name: "Satrio", Gender: "Pria"},
// 		{Name: "Agus", Gender: "Pria"},
// 		{Name: "Yudha", Gender: "Pria"},
// 		{Name: "Peter", Gender: "Pria"},
// 		{Name: "Aulia", Gender: "Wanita"},
// 		{Name: "Thalia", Gender: "Wanita"},
// 		{Name: "Giva", Gender: "Wanita"},
// 	}

// 	loop := func(persons []*Person) {
// 		for _, i := range persons {
// 			fmt.Println("- "+i.Name+",", i.Gender)
// 		}
// 	}

// 	loop(temanHactive)
// }

func tugas3() {
	binatang := map[string]interface{}{
		"nama":         "Ayam",
		"jumlah_telur": 3,
	}

	for key, value := range binatang {
		fmt.Println(key, ":", value)
	}
}

func tugas2() {
	temen := []string{"Sigit Setiawan", "Yosef Brahmantyo", "Bayu", "Satrio", "Agus", "Yudha", "Peter", "Aulia", "Thalia", "Giva"}

	for _, nama := range temen {
		fmt.Println("Halo Assalamualaikum", nama+", Salam Kenal Yak")
		fmt.Println("=============================================")
	}
}

func tugas1() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "Genap")
		} else {
			fmt.Println(i, "Ganjil")
		}
	}
}
