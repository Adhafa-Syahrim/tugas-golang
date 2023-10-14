package main

import "fmt"

func main() {
	//deklarasi variabel
	var tip, tagihan, total float64 	

	//menginput tagihan dari user jika manual
	fmt.Println("Program Penghitung TIP")
	fmt.Println("Masukkan nilai tagihan")
	fmt.Scanln(&tagihan)

	// perhitungan menggunakan if/else 
	if(tagihan >= 50 && tagihan <=300){
		fmt.Println("\nBesaran tip adalah 15%")
		tip = tagihan * 0.15
	}else{
		fmt.Println("\nBesaran tip adalah 20%")
		tip = tagihan * 0.2
	}

	//perhitungan menggunakan ternary operator tidak bisa dilakukan (tidak didukung) dalam bahasa golang
	//tip = (tagihan >= 50 && tagihan <=300) ? tagihan * 0.15 : tagihan * 0.2

	//total harga
	total = tagihan + tip

	//cetak hasil ke konsol dengan memanggila func cetak
	fmt.Printf("Tagihannya %.2f, tipnya %.2f, dan total nilainya %.2f\n", tagihan, tip, total)

}




