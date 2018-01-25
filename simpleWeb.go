package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"github.com/nanobox-io/golang-scribble"
	"encoding/json"
)

var startZeit,timeLaufzeit,timeEnd  = time.Now(),time.Now(),time.Now()

func sayhelloName(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	timeLaufzeit = time.Now()
	fmt.Fprintf(w, "Port Connect erfolgreich!\n\r" ) // send data to client side
	fmt.Fprintf(w,"Zeit :%v \n\r" , timeLaufzeit.Sub(startZeit) )
	fmt.Fprintf(w,"Laufzeit in Sekunden:%v \n\r" ,timeLaufzeit.Second())
	fmt.Fprint(w,"Laufzeit seit letztem polling in Sekunden :",timeLaufzeit.Sub(timeEnd))
	timeEnd =time.Now()
}

func main() {
	type Fish struct{ Name string
		Zeit time.Time }
	var dir = "c:\\temp\\scribbleDB"
	db, err := scribble.New(dir, nil)
		if err != nil {
		fmt.Println("Error", err)
	}
	fish := Fish{}
	if err := db.Write("fishx", "twofish", fish ); err != nil {
		fmt.Println("Error", err)
	}
	for _, name := range []string{"onefish", "twofish", "redfish", "bluefish"} {
		db.Write("fish", name, Fish{Name: name})
	}
	records, err := db.ReadAll("fish")
	if err != nil {
		fmt.Println("Error", err)
	}
	fishies := []Fish{}
	for _, f := range records {
		fishFound := Fish{}
		if err := json.Unmarshal([]byte(f), &fishFound); err != nil {
			fmt.Println("Error", err)
		}
		fishies = append(fishies, fishFound)

	}
	fmt.Println("sequence",fishies)

	http.HandleFunc("/", sayhelloName)       // set router
	err = http.ListenAndServe(":9091", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}


}

