package main

import "fmt"

func main() {
	//data uji
	data := []float64{275, 40, 430}

	//perhitungan besar tip dengan memanggil func hitung
	for i := range data {
		hitung(data[i])
	}
}

func hitung(tagihan float64){
	//deklarasi variabel
	var tip float64 
	
	//perhitungan menggunakan if/else 
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
	total := tagihan + tip

	//cetak hasil ke konsol dengan memanggila func cetak
	fmt.Printf("Tagihannya %.2f, tipnya %.2f, dan total nilainya %.2f\n", tagihan, tip, total)
}



