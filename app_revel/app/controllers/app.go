package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Saludo() revel.Result {
	return c.Render()
}

func (c App) Imagen() revel.Result {
	return c.Render()
}

func (c App) Formulario() revel.Result {
	return c.Render()
}

func (c App) Tabla() revel.Result {
	return c.Render()
}