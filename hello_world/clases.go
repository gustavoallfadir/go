package main

import "fmt"

type Persona struct {
	nombre string 
	apellido string 
	edad int 
	email string
}

func (p *Persona) get_full_name() string {
	return fmt.Sprintf("%s %s", p.nombre, p.apellido)
}

func (p *Persona) get_email() string {
	return p.email
}

func main(){
	gus := new(Persona)
	gus.nombre = "Gustavo"
	gus.apellido = "Robledo"
	gus.email = "guz.guitar@gmail.com"
	gus.edad = 33
	fmt.Println(gus.get_full_name())
	fmt.Println(gus.get_email())
}

