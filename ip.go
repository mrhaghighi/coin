package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getUserIpFromApi() {
	// Show message
	fmt.Println("-------------------------------")
	fmt.Println("Getting IP ...")
	fmt.Println("-------------------------------")

	// Getting IP & other details from API
	res, err := http.Get("https://ipinfo.io")
	if err != nil {
		log.Fatalln(err)
	}

	// Close connection
	defer res.Body.Close()

	// Cast response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Print body
	log.Println(string(body))
}