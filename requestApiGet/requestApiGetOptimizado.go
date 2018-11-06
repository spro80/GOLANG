package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"net/url"
	"reflect"
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
	
	//CREATE REQUEST
	//url := "https://jsonplaceholder.typicode.com/todos/"
	url := "https://jsonplaceholder.typicode.com"
	path := "/todos"
	urlRequest := createUrlRequest( url, path )


	//BUILD REQUEST
	//var statusBuildRequest int
	//statusBuildRequest = buildRequest( urlRequest )
	req, err := buildRequest( urlRequest )
	fmt.Println("Valor de req: ", req)
	fmt.Println("Valor de err: ", err)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}


	//SEND REQUEST:
	resp, err:= sendRequest( req )
	fmt.Println("Valor de resp: ", resp)
	fmt.Println("Valor de err: ", err)
	


	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()


	var record Usuario
	reflect.TypeOf( resp )

	//Procesa response data de API REST
	readData( record, resp )


	
}



func createUrlRequest(url string , path string ) string{
	
	urlRequest :=  url + path

	return urlRequest
}


func buildRequest( urlRequest string ) (*http.Request, error) {
	
	fmt.Println("\n###Iniciando en metodo buildRequest.");
	
	req, err := http.NewRequest("GET", urlRequest, nil)
	
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}
	
	fmt.Println("###Saliendo de metodo buildRequest.");
	return req, err	
}


func sendRequest( req *http.Request ) ( *http.Response , error ) {

	fmt.Println("\n###Iniciando en metodo sendRequest.");
	
	client := &http.Client{}
	
	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do( req )
	if err != nil {
		log.Fatal("ERROR en metodo sendRequest --> Do: ", err)
	}
	
	fmt.Println("###Saliendo de metodo sendRequest.");
	
	return resp, err
	
}

func readData( record Usuario, resp *http.Response ) int {
	
	fmt.Println("\n###Iniciando en metodo readData.");
	
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	
	/*fmt.Println( len(record) )
	fmt.Println( record[0].UserID )
	fmt.Println( record[0].ID )
	fmt.Println( record[0].Title )
	fmt.Println( record[0].Completed )
	*/
	
	var i int
	i = 0
	
	var largeArray int
	largeArray = len( record )
	
	
	for i < largeArray {
		//fmt.Print( i )
		fmt.Print( "    UserID: ", record[i].UserID )
		fmt.Print( " - ID: ", record[i].ID )
		fmt.Print( " - Title: ", record[i].Title )
		fmt.Print( " - Completed: ",record[i].Completed )
		fmt.Println("\n");
		i++
	}
	
	
	fmt.Println("###Saliendo de metodo readData.");
	return 0
}
