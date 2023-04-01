// pilih salah satu jawaban agar tidak sama

// package main

// import (
// 	"fmt"
// )

// func pola_1304221045(n int, s *string) {
// 	var i int

// 	i = 0
// 	if n > i {
// 		pola_1304221045(n-1, s)
// 		*s += "*"
// 		fmt.Println(*s)
// 	}
// }

// func main() {
// 	var n int
// 	var s string

// 	fmt.Scan(&n)
// 	pola_1304221045(n, &s)
// }

package main

import (
	"fmt"
)

func pola_1304221045(i, n int, s *string) {
	if n > i {
		pola_1304221045(i+1, n, s)
		*s += "*"
		fmt.Println(*s)
	}
}

func main() {
	var i,n int
	var s string

	fmt.Scan(&n)
	i = 0
	pola_1304221045(i, n, &s)
}