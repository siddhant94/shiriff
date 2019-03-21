package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"
)

type pageVariables struct {
	Date string
	Time string
}

type mlabReq struct {
	Query string
}

type mlabResponse struct {
	db          string
	collections string
	objects     string
	indexes     int
	avgObjSize  string
	dataSize    float64
	storageSize float64
	indexSize   float64
	fileSize    float64
}

const baseURL = "https://api.mlab.com/api/1"
const dbStats = "/databases/avianadb/runCommand" //Make avianadb name configurational
const apiKey = "your-api-key"

//{ dbStats: 1, scale: 1024 }

func main() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	now := time.Now()              // find the time right now
	homePageVars := pageVariables{ //store the date and time in a struct
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}
	// Get dbStats() response from mlab api.
	var str strings.Builder
	for _, c := range baseURL {
		str.WriteString(string(c))
	}
	for _, c := range dbStats {
		str.WriteString(string(c))
	}
	str.WriteString("?apiKey=")
	for _, c := range apiKey {
		str.WriteString(string(c))
	}

	value := mlabReq{`{"dbStats": 1, "scale": 1024}`}
	// jsonValue, err := json.Marshal(value.Query)
	fmt.Println("Marshalled JSON: ", value.Query)
	reqBody := bytes.NewBuffer([]byte(value.Query))
	mlabResp, err := http.Post(str.String(), "application/json", reqBody)
	fmt.Println(mlabResp.Status)
	if err != nil {
		// handle error
	}
	defer mlabResp.Body.Close()

	responseBody, err := ioutil.ReadAll(mlabResp.Body)
	fmt.Println(string(responseBody))
	// responseBodyString := string(responseBody)
	//fmt.Printf("%s", responseBodyString)

	var mlabData mlabResponse
	error := json.Unmarshal(responseBody, &mlabData)
	if error != nil {
		log.Print("Response unmarshalling error: ", error)
	}
	fmt.Println("Unmarshalled")
	fmt.Println(mlabData)

	t, err := template.ParseFiles("homepage.html") //parse the html file homepage.html
	if err != nil {                                // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, homePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                  // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
