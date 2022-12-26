package clase02parte1

import "testing"

// EJERCICIO 1
func TestCalcularImpuestoMayor150000(t *testing.T) {
	//ARRANGE
	salario := 180000.0
	impuestoEsperado := 27
	//ACT
	impuesto := CalcularImpuesto(salario)
	//ASSERT
	if impuesto != impuestoEsperado {
		t.Error()
	}
}

func TestCalcularImpuestoMayor50000(t *testing.T) {
	//ARRANGE
	salario := 100000.0
	impuestoEsperado := 17
	//ACT
	impuesto := CalcularImpuesto(salario)
	//ASSERT
	if impuesto != impuestoEsperado {
		t.Error()
	}
}

func TestCalcularImpuestoMenor50000(t *testing.T) {
	//ARRANGE
	salario := 10000.0
	impuestoEsperado := 0
	//ACT
	impuesto := CalcularImpuesto(salario)
	//ASSERT
	if impuesto != impuestoEsperado {
		t.Error()
	}
}

// EJERCICIO 2
