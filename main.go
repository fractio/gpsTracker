package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/paulmach/go.geojson"
	"github.com/timshannon/bolthold"
	bolt "go.etcd.io/bbolt"
)

var store *bolthold.Store

// Log handles /log and writes to the bolthold
func Log(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	startTimeInt, err := strconv.ParseInt(params["epoch"][0], 10, 64)
	timestamp, err := time.Parse(time.RFC3339, params["time"][0])
	acc, err := strconv.ParseFloat(params["acc"][0], 64)
	lat, err := strconv.ParseFloat(params["lat"][0], 64)
	lng, err := strconv.ParseFloat(params["lon"][0], 64)
	alt, err := strconv.ParseFloat(params["alt"][0], 64)
	spd, err := strconv.ParseFloat(params["spd"][0], 64)
	sat, err := strconv.ParseInt(params["sat"][0], 10, 0)
	dir, err := strconv.ParseFloat(params["dir"][0], 64)
	if err != nil {
		fmt.Println("Err")
		fmt.Println(err)
	}

	location := Location{
		StartTime:          time.Unix(startTimeInt, 0),
		ClientTimeStamp:    timestamp,
		ServerTimeStamp:    time.Now(),
		Accuracy:           acc,
		Lat:                lat,
		Lng:                lng,
		Altitude:           alt,
		Speed:              spd,
		Serial:             params["serial"][0],
		NumberOfSatellites: int(sat),
		Direction:          dir,
		Provider:           params["prov"][0],
	}
	key := location.Serial + "_" + strconv.FormatInt(location.StartTime.Unix(), 10) + "_" + strconv.FormatInt(location.ServerTimeStamp.UnixNano(), 10)
	err = store.Insert(key, location)
	fmt.Println(key)

	if err != nil {
		fmt.Println("Err")
		fmt.Println(err)
		fmt.Println(key)
	}

	w.WriteHeader(http.StatusOK)

}

//All returns all bolts for a thing
func All(w http.ResponseWriter, r *http.Request) {
	var result []Location
	err := store.Find(&result, bolthold.Where("Serial").Eq("ce011711bd1668d80c").Index("Serial"))
	if err != nil {
		fmt.Println("Err")
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(result)

}

//AllGeoJSON returns all bolts for a thing in geojson
func AllGeoJSON(w http.ResponseWriter, r *http.Request) {
	var locations []Location
	err := store.Find(&locations, bolthold.Where("Serial").Eq("ce011711bd1668d80c").Index("Serial"))
	if err != nil {
		fmt.Println("Err")
		fmt.Println(err)
	}

	var coords = make([][]float64, len(locations))
	for i, location := range locations {
		coords[i] = []float64{location.Lng, location.Lat}
	}

	fc := geojson.NewFeatureCollection()
	fc.AddFeature(geojson.NewLineStringFeature(coords))
	rawJSON, err := fc.MarshalJSON()

	w.Write(rawJSON)

}

//Tracks returns all bolts for a thing in geojson
func Tracks(w http.ResponseWriter, r *http.Request) {
	db := store.Bolt()
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("_index:Location:StartTime"))

		c := b.Cursor()

		var startTimes []time.Time

		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			// Create a decoder and receive a value.
			dec := gob.NewDecoder(bytes.NewReader(k))
			var v time.Time
			err := dec.Decode(&v)
			if err != nil {
				log.Fatal("decode:", err)
			}
			startTimes = append(startTimes, v)
		}
		json.NewEncoder(w).Encode(startTimes)

		return nil
	})

}

func main() {
	fmt.Println("GPS Tracking server")
	fmt.Println("Open Bolthold")
	var err error
	store, err = bolthold.Open("./data/db.db", 0666, nil)
	if err != nil {
		fmt.Println("Err")
		fmt.Println(err)
	}

	fmt.Println("Start router")

	router := mux.NewRouter()
	router.HandleFunc("/log", Log).Methods("GET")
	router.HandleFunc("/all", All).Methods("GET")
	router.HandleFunc("/all-geo-json", AllGeoJSON).Methods("GET")
	router.HandleFunc("/tracks", Tracks).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./client/build")))

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(router)))
}
