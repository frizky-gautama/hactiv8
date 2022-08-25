package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Name      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {

	// tugas1()
	// tugas2()
	// tugas3()
	args, _ := strconv.Atoi(os.Args[1])
	tugas_CLI(args)
}

func tugas_CLI(os int) {
	temanHactive := []*Person{
		{Name: "Sigit Setiawan", Alamat: "Jakarta", Pekerjaan: "IT Security", Alasan: "Gabut"},
		{Name: "Yosef Brahmantyo", Alamat: "Surabaya", Pekerjaan: "Middleware Programmer", Alasan: "Iseng"},
	}

	loop := func(persons []*Person, os int) {
		for key, i := range persons {
			if key == os {
				fmt.Println("Nama: "+i.Name+",", "Alamat: "+i.Alamat+",", "Pekerjaan: "+i.Pekerjaan+",", "Alasan: "+i.Alasan)
			}
		}
	}
	loop(temanHactive, os)
}

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
