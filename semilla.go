package main

import "fmt"

func main() {

	//Arrays de semilla
	//var arrays [2]string
	ArrS := []string{"a", "b", "c", "y"}
	ArrP := []string{"x", "b", "c", "y", "a"}

	for i := 0; i < len(ArrS); i++ {
		for i := 0; i < len(ArrP); i++ {
			if ArrS[i] == ArrP[i] {
				resul1 := ArrS[i]
				resul2 := ArrP[i]
				fmt.Println(resul1, resul2, len(ArrP), len(ArrS), cap(ArrP), cap(ArrS))
			}

		}
	}

}

//}
