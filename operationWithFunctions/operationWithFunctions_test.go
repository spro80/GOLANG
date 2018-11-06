package main

import(
	"testing"
)


func TestOperationWithFunctions_SumarNumeros(t *testing.T) {
		
		sumaTotal :=  sumarNumeros( 1, 2 )
		sumaExpected := 3

		if sumaTotal != sumaExpected {
			t.Errorf("El Resultado de la Suma es: [%d] y se esperaba [%d] ",sumaTotal, sumaExpected )
		}
				
}


func TestOperationWithFunctions_RestarNumeros(t *testing.T) {
		
		restar :=  restarNumeros( 4, 2 )
		restaExpected := 2

		if restar != restaExpected {
			t.Errorf("El Resultado de la Resta es: [%d] y se esperaba [%d] ",restar, restaExpected )
		}
				
}


func TestOperationWithFunctions_MultiplicarNumeros(t *testing.T) {
		
		multiplicar :=  multiplicarNumeros( 4, 2 )
		multiplicacionExpected := 8

		if multiplicar != multiplicacionExpected {
			t.Errorf("El Resultado de la Multiplicacion es: [%d] y se esperaba [%d] ",multiplicar, multiplicacionExpected )
		}
				
}


func TestOperationWithFunctions_DividirNumeros(t *testing.T) {
		
		dividir, _ :=  dividirNumeros( 500, 10 )
		var divisionExpected float64
		divisionExpected = 50

		if dividir != divisionExpected {
			t.Errorf("El Resultado de la Division es: [%f] y se esperaba [%f] ",dividir, divisionExpected )
		}
				
}
