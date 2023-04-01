package main

import "fmt"

func beliBuku_1304221045(n, m int) {
	var hasil int

	hasil = n-m
	if hasil != 0 {
		fmt.Println("beli 1 buku, sisa", hasil-1)
		beliBuku_1304221045(n, m+1)
	} else if hasil == 0 {
		fmt.Println("beli 1 buku, rak buku penuh")
	}
}

func main() {
	var n, m, rakKosong int

	fmt.Scan(&n, &m)
	fmt.Println("Kapasitas rak dan jumlah buku saat ini?", n, m)
	rakKosong = n - m
	fmt.Println("Sisa rak kosong", rakKosong, "buku")
	beliBuku_1304221045(n,m)
}