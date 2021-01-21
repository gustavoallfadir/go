package main

import "fmt"

type Frutas struct {
	lista_frutas []string
	conteo_frutas map[string]int
}

func (f *Frutas) hacer_conteo() {
	if f.conteo_frutas == nil {
        f.conteo_frutas = make(map[string]int)
	}
	for i := range f.lista_frutas {
		f.conteo_frutas[f.lista_frutas[i]] += 1
	}
}

func (f *Frutas) get_conteo() string {
	conteo := "\n"
	for key, value := range f.conteo_frutas{
		conteo += fmt.Sprintf("%s: %d\n", key, value)
	}
	return conteo
}

func main(){
	frutas := new(Frutas)
	frutas.lista_frutas = append(
		frutas.lista_frutas, 
		"manzana",
		"naranja",
		"manzana",
		"melon",
		"uva",
		"melon",
		"manzana",
		"limon",
		"lima",
		"melon",
		"naranja",
	)
	frutas.hacer_conteo()
	fmt.Println("-Lista original de frutas:", frutas.lista_frutas)
	fmt.Println("-Mapa de frutas con su conteo:", frutas.conteo_frutas)
	fmt.Println("-String dando resultados del conteo:", frutas.get_conteo())
}