package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/configor"
	"log"
	"net/http"
	"strconv"
)

func router(w http.ResponseWriter, req *http.Request) {
	// TODO: Implement proper router
	if req.URL.Path != "/api/temperature" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch req.Method {
	case "GET":
		var temperatures []Temperature = getAllTemperatures()

		json, err := json.Marshal(temperatures)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	case "POST":
		decoder := json.NewDecoder(req.Body)
		var temperature Temperature

		err := decoder.Decode(&temperature)
		handleError(err)
		insertTemperature(temperature)

		response := Response{"OK", 200}

		json, err := json.Marshal(response)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := Config.DB.User
	dbPass := Config.DB.Password
	dbName := Config.DB.Name
	dbHost := Config.DB.Host
	dbPort := strconv.Itoa(int(Config.DB.Port))
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)

	handleError(err)

	return db
}

func getAllTemperatures() []Temperature {
	db := dbConn()

	selDB, err := db.Query("SELECT room, temperature, humidity, time FROM temperatures ORDER BY id DESC")
	handleError(err)

	temperature := Temperature{}
	var res []Temperature

	for selDB.Next() {
		err = selDB.Scan(&temperature.Room, &temperature.Temperature, &temperature.Humidity, &temperature.Date)
		handleError(err)
		res = append(res, temperature)
	}

	return res
}

func insertTemperature(temperature Temperature) {
	db := dbConn()

	_, err := db.Query("INSERT INTO temperatures (room, temperature, humidity) VALUES (?, ?, ?)", temperature.Room, temperature.Temperature, temperature.Humidity)
	handleError(err)
	defer db.Close()
	log.Println("Temperatuur", temperature.Temperature, "Õhuniiskus", temperature.Humidity, "Ruum", temperature.Room)
}

func handleError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func main() {
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)
	http.HandleFunc("/", router)

	err := configor.Load(&Config, "config.yaml")
	handleError(err)

	log.Printf("Starting Thermometrum on port 8001...\n")
	log.Println("Ver. " + Config.AppVersion)

	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatal(err)
	}
}

type Temperature struct {
	Room        string  `json:"room"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Date        string
}

type Response struct {
	Message string
	Code    int
}
