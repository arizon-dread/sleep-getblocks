package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"gitlab.com/arizon-dread/sleep-getblocks/api"
)

func main() {

	strDelay := os.Getenv("DELAY")

	intDelay, err := strconv.Atoi(strDelay)
	if err == nil {
		api.SetDelay(intDelay)
	} else {
		log.Printf("error reading env var DELAY as int, using 60s as delay. Error was %v\n", err)
		api.SetDelay(60)
	}
	content, err := os.ReadFile("response.xml")
	if err != nil {
		log.Printf("Could not read file b/c: %v\nWill use static response.\n", err)
	} else {
		log.Printf("Will use response.xml as response for every request.\n")
	}
	api.SetText(string(content))

	//fmt.Printf("formated time: %v\n", currentTime)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", api.Healthz)
	mux.HandleFunc("POST /sleep", api.Sleep)
	mux.HandleFunc("POST /getblocks", api.GetBlocks)
	http.ListenAndServe(":8080", mux)
}
