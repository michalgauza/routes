package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
)

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc(routesPath, routes)
	port := os.Getenv("PORT")
	if port == "" {
		port = myPort
		fmt.Println("port:" + port)
	}
	err := http.ListenAndServe(":"+port, router)
	log.Fatal(err)
}

func routes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")

	myResponse := MyResponse{}
	setupSource(&myResponse, r)
	setupRoutes(&myResponse, r)
	sort.Sort(myResponse.Routes)
	json.NewEncoder(w).Encode(myResponse)
}

func setupDurationAndDistance(myResponse *MyResponse, index int) {
	url := getOSMRUrl(myResponse.Source, myResponse.Routes[index].Destination)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		myResponse.Err = err.Error()
		return
	} else {
		osmrResponse, err := getOSMRResponse(response)
		if len(osmrResponse.Routes) > 0 {
			myResponse.Routes[index].Distance = osmrResponse.Routes[0].Distance
			myResponse.Routes[index].Duration = osmrResponse.Routes[0].Duration
		}
		myResponse.Err = err
	}
}

func getOSMRResponse(response *http.Response) (OSMRResponse, string) {
	data, err := ioutil.ReadAll(response.Body)
	osmrResponse := OSMRResponse{}
	if err != nil {
		fmt.Println(err)
		return osmrResponse, err.Error()
	}
	err = json.Unmarshal(data, &osmrResponse)
	if err != nil {
		fmt.Println(err)
		return osmrResponse, err.Error()
	}
	if osmrResponse.Code != osmrResponseOkCode {
		fmt.Println(osmrResponse.Code)
		return osmrResponse, osmrResponse.Code
	}
	return osmrResponse, ""
}

func setupSource(myResponse *MyResponse, r *http.Request) {
	src, ok := r.URL.Query()[srcKey]
	if !ok || len(src[0]) < 1 {
		fmt.Println(missingSrcErr)
		myResponse.Err = missingSrcErr
		return
	}
	myResponse.Source = src[0]
}

func setupRoutes(myResponse *MyResponse, r *http.Request) {
	dst, ok := r.URL.Query()[dstKey]
	if !ok || len(dst[0]) < 1 {
		fmt.Println(missingDstErr)
		myResponse.Err = missingDstErr
		return
	}
	for index, val := range dst {
		myResponse.Routes = append(myResponse.Routes, Route{Destination: val})
		setupDurationAndDistance(myResponse, index)
	}
}

func main() {
	handleRequests()
}
