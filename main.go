package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type apiResponse struct {
	Id, Kind, LongURL string
}

func main() {

	longURL := os.Args[len(os.Args)-1]

	body := bytes.NewBufferString(fmt.Sprintf(
		`{"longURL":"%s"}`,
		longURL))
// This api is not responding , I have checked with postman.
	var request, err = http.NewRequest(
		"POST",
		"https://godoc.org/google.golang.org/api/urlshortener/v1",
		body)

	request.Header.Add("Content-Type", "application/json")

	client := http.Client{}

	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	outputAsBytes, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	var output apiResponse
	err = json.Unmarshal(outputAsBytes, &output)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", output.Id)

}