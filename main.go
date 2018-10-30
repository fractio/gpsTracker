package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// map[acc:[16.0]
//  battery:[98.0]
//  androidId:[a27642f982771040]
//  activity:[]
//  desc:[]
//  spd:[0.0]
//  time:[2018-10-29T20:30:47.000Z]
//  serial:[ce011711bd1668d80c]
//  sat:[9]
//  dir:[0.0]
//  prov:[gps]
//  epoch:[1540845047]
//  lat:[51.42925195758545]
//  lon:[-0.10596947765469795]
// alt:[159.1731263161311]]

// Location stores the GPS coords
type Location struct {
	StartTime          time.Time `boltholdIndex:"StartTime"`
	TimeStamp          time.Time
	Accuracy           float64
	Lat                float64
	Lng                float64
	Altitude           float64
	Speed              float64
	Serial             string `boltholdIndex:"Serial"`
	NumberOfSatallites int
	Direction          float64
	Provider           string
}

func Log(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	fmt.Println(params)
	startTimeInt, err := strconv.ParseInt(params["epoch"][0], 10, 64)
	timestamp, err := time.Parse(time.RFC3339, params["time"][0])
	acc, err := strconv.ParseFloat(params["acc"][0], 64)
	lat, err := strconv.ParseFloat(params["lat"][0], 64)
	lng, err := strconv.ParseFloat(params["lon"][0], 64)
	alt, err := strconv.ParseFloat(params["alt"][0], 64)
	spd, err := strconv.ParseFloat(params["spd"][0], 64)
	sat, err := strconv.ParseInt(params["sat"][0], 10, 0)
	dir, err := strconv.ParseFloat(params["dir"][0], 64)
	fmt.Println(err)
	result := Location{
		StartTime:          time.Unix(startTimeInt, 0),
		TimeStamp:          timestamp,
		Accuracy:           acc,
		Lat:                lat,
		Lng:                lng,
		Altitude:           alt,
		Speed:              spd,
		Serial:             params["serial"][0],
		NumberOfSatallites: int(sat),
		Direction:          dir,
		Provider:           params["prov"][0],
	}

	fmt.Printf("%#v", result)

}

func main() {
	fmt.Println("GPS Tracking server")
	fmt.Println("Open Bolthold")

	// store, err := bolthold.Open("gpsTracker.db", 0666, nil)
	// if err != nil {
	// 	fmt.Println("Err")
	// 	fmt.Println(err)
	// }
	// err = store.Insert("key", &Item{
	// 	Name:    "Test Name",
	// 	Created: time.Now(),
	// })
	// if err != nil {
	// 	fmt.Println("Err")
	// 	fmt.Println(err)
	// }

	fmt.Println("Start router")

	router := mux.NewRouter()
	router.HandleFunc("/log", Log).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
