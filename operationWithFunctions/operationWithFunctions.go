package main



import(
	"fmt"
	"errors"
	"reflect"
)



func main(){
	fmt.Println("Iniciando en metodo main...")

	//Operations: Suma
	var sumando1 int
	var sumando2 int
	sumando1= 10
	sumando2= 15

	//Operations: Resta
	var minuendo int
	var sustraendo int
	minuendo= 100
	sustraendo= 55

	//Operations: Multiplicacion
	var producto1 int
	var producto2 int
	producto1 = 24
	producto2 = 2

	//Operations: Division
	var dividendo float64
	var divisor float64
	dividendo = 500
	divisor = 0

	sumaTotal := sumarNumeros( sumando1, sumando2)
	fmt.Println("    La Suma Total es: ", sumaTotal)

	resta := restarNumeros( minuendo, sustraendo)
	fmt.Println("    La resta es:", resta)

	multiplicacion := multiplicarNumeros( producto1, producto2 )
	fmt.Println("    La Multiplicacion es:", multiplicacion)

	division, err := dividirNumeros( dividendo, divisor)
	if err == nil {
		fmt.Println( reflect.TypeOf(division) )
		fmt.Println("    La Division es: ", division)
	}else {
		fmt.Println( err )
	}


	fmt.Println("Saliendo de metodo main...")
}



func sumarNumeros( number1 int, number2 int) int{
	return number1 + number2
}

func restarNumeros( number1 int, number2 int) int{
	return number1-number2
}

func multiplicarNumeros( producto1 int, producto2 int) int{
	return producto1 * producto2
}

func dividirNumeros( dividendo float64, divisor float64) (float64, error){

	fmt.Println( "Dividendo: ",dividendo )
	fmt.Println( "Divisor: ",divisor )

	var division float64
	division = 0

	if divisor == 0 {
		fmt.Println("Error al Dividir, el Divisor no puede ser 0.");
		return 0, errors.New( fmt.Sprintf("Can't divide %f by zero", dividendo) )
	}else{
		division= dividendo / divisor
	}


	return division, nil
}

func Imprimir( texto string ) string{
	fmt.Println("Imprime el texto")
	return texto + " Nuevo!!!"
}
