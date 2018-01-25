package main

import (
	"database/sql"
	"fmt"
	_ "github.com/alexbrainman/odbc"
	"log"
	"time"
)

func connectPervasive() *sql.DB {
	// Connect
	// Original db, err := sql.Open("odbc", "DSN=testdb;Uid=root;Pwd=test1234;")
	// String zum Insert in Pervasive
	//Insert into dblftb (dARTNR,lLIEFNR,siBEH,dtDATAB,dtDATBIS,cSTATUS,dEKL,dtDATEK,cArtNrL,siBEH2,lFOLNR,siKONDGRZ,siKONDGRS,dARTNRHO,siLIEFZT,siBESTRY,cBESTTAG,cREGETI,cStueck,cLArt,dUmFact,cSondBest,cSortBer,siDispoVor,bitLiefTag,cBezeichn,siEinraeum,dtNArtDruck,dtGArtDruck,cHerk,dtLetztAend,siWESoBe,dtLetztSAend,cUserkz,siKondGrA,cSHerk,siNachLief,dEAN,dEKGrh,siEKVerwenden,sitkgruppe,sizsgruppe,ceantyp,dtDatEKGr,sSoBeVA,dMinBeMg) values (359072,1325,10,'2016-06-30','2089-01-01','A',222,'2017-08-28','2432334',0,0,3,0,2432334,0,0,'ÿ','1','0','',0,'1','ORDSHT',1,-1,'',0,'1000-01-01','1000-01-01','','1000-01-01',0,'1000-01-01','',0,'',0,0,0,0,0,0,'','1000-01-01','',0);

	db, err := sql.Open("odbc", "DSN=db210dbdemo;Uid=Master;Pwd=test1234;")

	fmt.Print(err)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func main() {

	db := connectPervasive()
	err := db.Ping()

	if err != nil {
		panic(err.Error())
	}
	// Insert a example

	stmt, err := db.Prepare("INSERT INTO example(time) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(time.Now())
	if err != nil {
		log.Fatal(err)
	}

	var (
		id     int
		dbTime time.Time
	)
	// Try to fetch the example
	rows, err := db.Query("select id, time from example where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &dbTime)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, dbTime)
	}
}
