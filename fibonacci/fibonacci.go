package main

import "fmt"

func fibonacci(max int){
	var a = 0;
	var b = 1;
	var c = a + b;
	fmt.Println("Inicializando secuencia fibonacci del 0 al", max)
	fmt.Println(a);
	fmt.Println(b);
	for i := 0; i < max; i++ {
		fmt.Println(c);
		a = b;
		b = c;
		c = a+b;
		if c > max {
			break
		}		
	}	
}

func main(){
	fibonacci(500);
}