package main

import (
	"fmt"
	"os"
)

func main() {
	var size int
	fmt.Println("Введите размер шахматной доски:")
	fmt.Fscan(os.Stdin, &size)

	if size <= 0 || size%2 != 0 || size < 20 {
		fmt.Println("Размер доски должен быть больше нуля и меньше 20 и четным числом !!!")
	} else {
		for i := 1; i <= size; i++ {
			for a := 0; a <= size; a++ {
				if (i+a)%2 == 0 && a < size {
					fmt.Print(" ")
				} else if (i+a)%2 != 0 && a < size {
					fmt.Print("#")
				} else if a == size {
					fmt.Print("\n")
				}
			}
		}
	}
}
