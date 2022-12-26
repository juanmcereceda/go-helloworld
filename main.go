package main

import (
	"fmt"
	//clase01_parte1 "github.com/juanmcereceda/go-helloworld/clase01_parte1"
	//clase01_parte2 "github.com/juanmcereceda/go-helloworld/clase01_parte2"
	//clase02parte1 "github.com/juanmcereceda/go-helloworld/clase02_parte1"
)

func main() {
	//ejemploSwitch()
	//clase01_parte1.Solucion()
	//clase01_parte2.Solucion()
	//clase02parte1.Resolucion()
}

func ejemploSwitch() {
	listaVerduleria := []string{"Manzana", "Zapallo", "Sandía", "Papa", "Lechuga", "Tomate"}
	for _, item := range listaVerduleria {
		switch item {
		case "Manzana", "Sandía":
			fmt.Println("Fruta")
		case "Zapallo", "Papa", "Lechuga":
			fmt.Println("Verdura")
		default:
			fmt.Println("Ni fruta ni verdura")
		}
	}
}
