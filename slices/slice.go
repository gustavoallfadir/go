package main

import (
	"fmt"
	"sort"
)

func main()  {
	//Creamos un slice con un solo elemento
	var slice = make([]string, 1)
	slice[0] = "volkswagen"
	fmt.Println(slice)
	//Creamos un nuevo slice que incluye el primero y nuevos elementos
	nuevo_slice := append(slice, "ford","nissan","mitsubishi","toyota","volvo")
	fmt.Println(nuevo_slice)	
	//Creamos un nuevo slice con elementos
	frutas := []string{"manzana","naranja","limon","piña"}
	fmt.Println(frutas)
	//Un nuevo slice de números 
	numeros := []uint8{1,8,4,5,6,2,3,9}
	fmt.Println(numeros.order())
}
