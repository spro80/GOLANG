package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"net/url"
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

	url := "https://jsonplaceholder.typicode.com/todos/"

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record Usuario

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Println( len(record) )
	fmt.Println( record[0].UserID )
	fmt.Println( record[0].ID )
	fmt.Println( record[0].Title )
	fmt.Println( record[0].Completed )

	var i int
	i = 0
	
	var largeArray int
	largeArray = len( record )
		
	for i < largeArray {
		fmt.Println( i )
		fmt.Println( "    UserID: ", record[i].UserID )
		fmt.Println( "    ID: ", record[i].ID )
		fmt.Println( "    Title: ", record[i].Title )
		fmt.Println( "    Completed: ",record[i].Completed )
		i++
	}
	
}


