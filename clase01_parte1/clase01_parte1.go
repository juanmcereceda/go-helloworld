package clase01_parte1

import "fmt"

func Solucion() {
	// Ej 1
	ejercicio1()
	// Ej 2
	ejercicio2()
	ejercicio3()
	ejercicio4()
}

func ejercicio1() {
	var nombre string = "Juan"
	fmt.Println(nombre)
	var direccion string = "Sarasa 123"
	fmt.Println(direccion)
}

func ejercicio2() {
	var temperatura float32 = 32.0
	var humedad float32 = 0.28
	var presion float32 = 1009.0
	fmt.Println(temperatura)
	fmt.Println(humedad)
	fmt.Println(presion)
}
func ejercicio3() {
	boolean := "false"
	fmt.Println(boolean)
}
func ejercicio4() {}
