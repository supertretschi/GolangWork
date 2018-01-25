package main

import (
	"bufio"

	"fmt"
	"log"
	"os"
)

var (
	pfad         string = "B:\\Portable\\go\\work\\src\\Golan\\SqlError\\"
	eingabeText  string = "Bitte Path und Dateiname"
	fehlerEins   string = "Fehler bei der Eingabe: "
	fehlerZwei   string = "Problem mit Datei"
	defaultDatei string = pfad + "dblftb.asc"
)

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}

	return string(ln), err
}
func eingabe() *bufio.Reader {
	fmt.Print(defaultDatei)
	FileNeu := bufio.NewReader(os.Stdin)
	fmt.Print(eingabeText)
	return FileNeu

}
func main() {

	DateiNeu := eingabe()
	datei, err := DateiNeu.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, fehlerEins, err)
	}
	if datei == "\n" {
		datei = defaultDatei
	}
	fmt.Println(datei)

	// file, err := os.Open()
	file, err := os.Open(datei)
	if err != nil {
		fmt.Print(fehlerZwei)
		log.Fatal(err)
	}
	defer file.Close()
	lesen := bufio.NewReader(file)

	gelesen, err := Readln(lesen)
	var count int64 = 1
	for err == nil {
		fmt.Print(count)
		fmt.Println(gelesen)
		gelesen, err = Readln(lesen)
		count++
		if count > 3 {
			break
		}
	}

	//scanner := bufio.newscanner(file)
	//for scanner.scan() {
	//
	//	fmt.println(scanner.scan())
	//
	//}
	//
	//if err := scanner.err(); err != nil {
	//	log.fatal(err)
	//}
}
