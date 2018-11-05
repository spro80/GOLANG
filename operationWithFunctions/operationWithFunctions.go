package main



import(
	"fmt"
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
	//multiplicacion = producto1*producto2
	//fmt.Println("Multiplicacion: ");
	//fmt.Println( multiplicacion );

	//Operations: Division
	var dividendo int
	var divisor int
	dividendo = 500
	divisor = 0
/*	//division = dividendo / divisor
	//fmt.Println("La division es: ");
	//fmt.Println( division )
*/

	sumaTotal := sumarNumeros( sumando1, sumando2)
	fmt.Println("La Suma Total es: ");
	fmt.Println( sumaTotal );

	resta := restarNumeros( minuendo, sustraendo)
	fmt.Println("La resta es:")
	fmt.Println( resta )

	multiplicacion := multiplicarNumeros( producto1, producto2 )
	fmt.Println("La Multiplicacion es:");
	fmt.Println( multiplicacion );

	division, error := dividirNumeros( dividendo, divisor)
	if error == -1 {
		fmt.Println("El error es -1, no se puede dividir por cero!!!");
	}else {
		fmt.Println("La division es correcta.")
		fmt.Println("La Division es: ");
		fmt.Println( division )
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

func dividirNumeros( dividendo int, divisor int) (int, int){

	var division int
	division = 0

	var errors int
	errors = 0


	if divisor == 0{
		fmt.Println("Error al Dividir, el Divisor no puede ser 0.");
		errors = -1
	}else{
		division= dividendo /divisor
	}


	return division, errors
}

func Imprimir( texto string ) string{
	fmt.Println("Imprime el texto")
	return texto + " Nuevo!!!"
}
