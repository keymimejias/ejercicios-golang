package main

import "fmt"

func main() {

	pali := []string{"o", "t", "t", "o"}

	for i := 0; i < len(pali); i++ {
		if i < 4 {
			if pali[i] == pali[i+3] {

				fmt.Println(true)

				fmt.Println("La palabra", pali, "si es palindrome")
				//break

			} else {
				fmt.Println(false)

				fmt.Println("La palabra", pali, "no es palindrome")
				//break
			}
		}
	}

}
