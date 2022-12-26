package clase01_parte2

import (
	"fmt"
)

func Solucion() {
	ejercicio1()
	ejercicio2()
}

// ej 1
func ejercicio1() {
	palabra := "teclado"
	imprimirCantidadDeLetras(palabra)
	imprimirLetras(palabra)
}

func imprimirCantidadDeLetras(palabra string) {
	fmt.Println(len(palabra))
}

func imprimirLetras(palabra string) {
	for i := 0; i < len(palabra); i++ {
		fmt.Println(i)
	}
}

// ej 2

func ejercicio2() {
	edad := 20
	esEmpleado := false
	antiguedad := 0
	sueldo := 120000.25
	puedeRecibirPrestamo(edad, esEmpleado, antiguedad, sueldo)
}

func puedeRecibirPrestamo(edad int, esEmpleado bool, antiguedad int, sueldo float64) {
	var restricciones []string
	if edad < 22 {
		restricciones = append(restricciones, "Debe ser mayor de 22 años")
	}
	if !esEmpleado {
		restricciones = append(restricciones, "El usuario no es empleado")
	}
	if antiguedad < 1 {
		restricciones = append(restricciones, "No tiene al menos un año de antiguedad")
	}
	if len(restricciones) > 0 {
		fmt.Println("No se pudo otorgar el préstamo")
		for _, restriccion := range restricciones {
			fmt.Println(restriccion)
		}
	} else {
		fmt.Println("Préstamo autorizado")
		if sueldo > 100000 {
			fmt.Println("Por su elevado sueldo no se le cobrarán intereses")
		}
	}
}
