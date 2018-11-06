package main

import (
	"fmt"
	"log"
	//"reflect"
)

//Generador de struct para GO: https://mholt.github.io/json-to-go/
//Se debe copiar el Json retornado por la peticion GET, se copia en la pagina y genera el Struct que se debe generar para peticion GET.
type Usuario []struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("\n###Iniciando en metodo main.");
	
	////CREATE REQUEST
	//url := "https://jsonplaceholder.typicode.com/todos/"
	url := "https://jsonplaceholder.typicode.com"
	path := "/todos"
	urlRequest := createUrlRequest( url, path )


	////BUILD REQUEST
	req, err := buildRequestGet( urlRequest )
	fmt.Println("Valor de req: ", req)
	fmt.Println("Valor de err: ", err)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}


	////SEND REQUEST
	resp, err:= sendRequest( req )
	fmt.Println("Valor de resp: ", resp)
	fmt.Println("Valor de err: ", err)


	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	
	////PROCESS DATA RESPONSE GET
	var record Usuario
	readDataResponse( record, resp )
	
	fmt.Println("\n###Saliendo de metodo main.");
}
