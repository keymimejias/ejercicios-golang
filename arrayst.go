package main

import "fmt"

func main() {

	//Creamos un array con cap de 8, para guardar la letras contiguas.
	igual := [8]string{}
	//Arrays Contiguos: Definimos el arrays con 9 elemento.
	e := []string{"a", "a", "c", "c", "s", "x", "x", "x", "x", "a"}
	//Creamos una variable b que la iniciamos en 0.
	b := 0
	//Creamos el for para recorrer el array (e).
	for i := 0; i <= len(e); i++ {
		//Creamos una condición cuando i sea < 9 se salga del for.
		if i < 9 {
			//Creamos una condición para coparar lo que tenemos en el array (e) y el indice (i) sea igual al array (e) y en cada vuelta se aumente 1.
			if e[i] == e[i+1] {

				//si es verdadero guardame el array (e) son su indice en el array (igual) su valor.
				igual[b] = e[i]
				//guardame del array (ei) su valor en (b) y a b sumale 1 en cada recorrido.
				igual[b+1] = e[i]

				//Y al indice i en cada recorrido sumale 1.
				i++
				//Y al indice b en cada recorrido sumale 2.
				b = b + 2

			}
		}

	}
	fmt.Println(igual)
}
