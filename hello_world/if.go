package main

import "fmt"

func mayor_de_edad(edad uint8) {
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}	
}

func main(){
	var edad uint8
	fmt.Println("¿Cuál es tu edad?")
	fmt.Scanln(&edad)
	mayor_de_edad(edad)
}	