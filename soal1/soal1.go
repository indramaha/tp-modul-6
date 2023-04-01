//jawaban dari chat GPT, nomor 1 belum tau cara pastinya

package main

import "fmt"

func abc_1304221045(x int) string {
	var s string
    if x == 0 {
        return ""
    } else {
        s = abc_1304221045(x / 2)
        if x % 2 == 0 {
            return s + "0"
        } else {
            return s + "1"
        }
    }
}

func main() {
	var x int
	var hasil string

	fmt.Scan(&x)
	hasil = abc_1304221045(x)
	fmt.Println(hasil)
}