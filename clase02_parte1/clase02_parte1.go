package clase02parte1

import "fmt"

// func Resolucion() {
// 	// Ejercicio 1
// 	impuesto := Ejercicio1()
// 	fmt.Printf("El impuesto total es de %d%%\n", impuesto)
// 	// Ejercicio 2
// 	ejercicio2()
// 	// Ejercicio 3
// 	salario := calcularSalario(40, "C")
// 	fmt.Printf("El sueldo es de %f\n", salario)
// 	// Ejercicio 4
// 	ejercicio4()
// 	// Ejercicio 5
// 	ejercicio5()
// }

// func Ejercicio1() int {
// 	impuesto := CalcularImpuesto(170000)
// 	return impuesto
// }

func CalcularImpuesto(salario float64) int {
	var impuesto int
	switch {
	case salario >= 150000:
		impuesto += 10
		fallthrough
	case salario >= 50000:
		impuesto += 17
	}
	return impuesto
}

func ejercicio2() {
	notas := []int{7, 5, 9, 8, 4, 11}
	promedio := calcularPromedio(notas)
	if promedio > 0 {
		fmt.Printf("El promedio es %f", promedio)
	} else {
		fmt.Println("Notas ingresadas inválidas")
	}
}
func calcularPromedio(notas []int) float64 {
	var resultado float64
	suma, esValido := sumar(notas...)
	if !esValido {
		return 0
	}
	resultado = float64(suma / len(notas))
	return resultado
}

func sumar(numeros ...int) (total int, esValido bool) {
	total = 0
	esValido = true
	for _, numero := range numeros {
		if numero > 10 || numero < 0 {
			esValido = false
		}
		total += numero
	}
	return
}

func calcularSalario(minutosTrabajados int, categoria string) float64 {
	categoriaYSueldoPorHora := map[string]int{"A": 1000, "B": 1500, "C": 3000}
	categoriaYPorcentajeExtra := map[string]float64{"A": 0, "B": 0.2, "C": 0.5}
	sueldo := float64(minutosTrabajados * categoriaYSueldoPorHora[categoria])
	return sueldo * (1 + categoriaYPorcentajeExtra[categoria])
}

func ejercicio4() {
	const (
		minimum = "minimum"
		average = "average"
		maximum = "maximum"
	)
	minFunc := operation(minimum)
	averageFunc := operation(average)
	maxFunc := operation(maximum)
	notas := []int{2, 3, 3, 4, 10, 2, 4, 5}
	minValue := minFunc(notas...)
	averageValue := averageFunc(notas...)
	maxValue := maxFunc(notas...)
	fmt.Printf("El minimo es %d, el promedio es %d y el máximo es %d", minValue, averageValue, maxValue)
}
func operation(operation string) func(...int) int {
	var operations = map[string]func(...int) int{"minimum": getMin, "average": getAverage, "maximum": getMax}
	return operations[operation]
}
func getMin(numbers ...int) int {
	var min = numbers[0]
	for _, number := range numbers {
		if number < min {
			min = number
		}
	}
	return min
}
func getAverage(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total / len(numbers)
}
func getMax(numbers ...int) int {
	var min = numbers[0]
	for _, number := range numbers {
		if number > min {
			min = number
		}
	}
	return min
}

// func ejercicio5() {}
