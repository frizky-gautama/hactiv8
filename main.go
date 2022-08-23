package main

import "fmt"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i, "Genap")
	// 	} else {
	// 		fmt.Println(i, "Ganjil")
	// 	}
	// }

	temen := []string{"Sigit Setiawan", "Yosef Brahmantyo", "Bayu", "Satrio", "Agus", "Yudha", "Peter", "Aulia", "Thalia", "Giva"}

	for _, s := range temen {
		fmt.Println("Halo Assalamualaikum", s+", Salam Kenal Yak")
		fmt.Println("=============================================")
	}
}
