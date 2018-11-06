package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)



func createUrlRequest(url string , path string ) string{
	
	urlRequest :=  url + path

	return urlRequest
}


func buildRequestGet( urlRequest string ) (*http.Request, error) {
	
	fmt.Println("\n###Iniciando en metodo buildRequestGet.");
	
	req, err := http.NewRequest("GET", urlRequest, nil)
	
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}
	
	fmt.Println("###Saliendo de metodo buildRequestGet.");
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


func readDataResponse( record Usuario, resp *http.Response ) int {
	
	fmt.Println("\n###Iniciando en metodo readDataResponse.");
	
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	
	var i int
	i = 0	
	var largeArray int
	largeArray = len( record )
	
	for i < largeArray {
		fmt.Print( "    UserID: ", record[i].UserID )
		fmt.Print( " - ID: ", record[i].ID )
		fmt.Print( " - Title: ", record[i].Title )
		fmt.Print( " - Completed: ",record[i].Completed )
		fmt.Println("\n");
		i++
	}
	
	fmt.Println("###Saliendo de metodo readDataResponse.");
	return 0
}
