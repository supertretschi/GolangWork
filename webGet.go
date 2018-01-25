package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"encoding/json"
)
var depric = false
var addresseInit = "http://www.google.com/robots.txt"
var address2 = "http://localhost:9091"
var addressCouch = "http://127.0.0.1:5984"
var unmarshal = true
type JObject struct {
	Name  string
	Order string
}
func abgekuendigt() {if depric == true {
	fmt.Printf("%s", "Abgekündigt")
	os.Exit(99)
}}
func fehler(fehl error) {
	if fehl != nil {
		log.Fatal(fehl)
	}
}
func main() {
	abgekuendigt()
	local := true
	addressCouchSwitch := true
	if local == true && addressCouchSwitch == true{
		addresseInit=addressCouch

	}
	res, err := http.Get(addresseInit)
	fehler(err)
	robots, err := ioutil.ReadAll(res.Body)
	if unmarshal == true {
		var animals []JObject

		err := json.Unmarshal(robots, &animals)
		fehler(err)
	}
	res.Body.Close()
	fehler(err)
	fmt.Printf("%s", robots)
	fmt.Printf("%s", )
}
