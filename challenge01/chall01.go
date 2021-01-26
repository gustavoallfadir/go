/* Crear una función que retorne la suma de los valores que sean 
multiplos de 3 o de 5 y que sean menores que un número dado. */
package main

import "fmt"

func resultado(max int) int {
	var multiplos []int
	for i := 0; i < max; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			multiplos = append(multiplos, i);
		}
	}	
	var suma = 0;
	for i := range multiplos {
		suma += multiplos[i];
	}
	return suma;
}

func main()  {
	var res = resultado(10);
	fmt.Println(res);
}