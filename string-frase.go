package main

import (
	"fmt"
)

func main() {

	palaFrase := []string{
		"programa",
		"programa",
	}

	/*for _, i := range palaFrase {
	fmt.Println(i, palaFrase)
	if palaFrase[0] == palaFrase[1] {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
	b := 0*/
	//Creamos el for para recorrer el array (e).
	for i := 0; i <= len(palaFrase); i++ {
		//Creamos una condición cuando i sea < 9 se salga del for.
		if i < 1 {
			//Creamos una condición para coparar lo que tenemos en el array (e) y el indice (i) sea igual al array (e) y en cada vuelta se aumente 1.
			if palaFrase[i] == palaFrase[i+1] {
				fmt.Println(true)
			} else {
				fmt.Println(false)
			}
			//si es verdadero guardame el array (e) son su indice en el array (igual) su valor.
			//igual[b] = e[i]
			//guardame del array (ei) su valor en (b) y a b sumale 1 en cada recorrido.
			//igual[b+1] = e[i]

			//Y al indice i en cada recorrido sumale 1.
			//i++
			//Y al indice b en cada recorrido sumale 2.
			//b = b + 2

		}

	}
	//fmt.Println(igual)
}

/*var a, b []string

a = palaFrase[0:1]
for i := 0; i < len(a); {
	fmt.Println(a[i])
	break
}
b = palaFrase[1:]
for i := 0; i < len(b); {
	fmt.Println(b[i])
	break
}*/

//}
//}
