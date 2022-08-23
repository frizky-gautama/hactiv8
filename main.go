package main

import "fmt"

func main() {
	tugas1()
	tugas2()
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

func tugas2() {
	temen := []string{"Sigit Setiawan", "Yosef Brahmantyo", "Bayu", "Satrio", "Agus", "Yudha", "Peter", "Aulia", "Thalia", "Giva"}

	for _, nama := range temen {
		fmt.Println("Halo Assalamualaikum", nama+", Salam Kenal Yak")
		fmt.Println("=============================================")
	}
}
