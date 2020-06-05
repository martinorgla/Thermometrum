package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/configor"
	"log"
	"net/http"
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
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(db:"+string(Config.DB.Port)+")/"+dbName)

	handleError(err)

	return db
}

func getAllTemperatures() []Temperature {
	db := dbConn()

	selDB, err := db.Query("SELECT * FROM temperatures ORDER BY id DESC")
	handleError(err)

	temperature := Temperature{}
	var res []Temperature

	for selDB.Next() {
		var id int
		var room, time string
		var temp, humidity float32

		err = selDB.Scan(&id, &room, &temp, &humidity, &time)

		handleError(err)

		temperature.Temperature = temp
		temperature.Humidity = humidity
		temperature.Room = room

		res = append(res, temperature)
	}

	return res
}

func insertTemperature(temperature Temperature) {
	db := dbConn()

	_, err := db.Query("INSERT INTO temperatures (room, temperature, humidity) VALUES (?, ?, ?)", temperature.Room, temperature.Temperature, temperature.Humidity)
	handleError(err)

	defer db.Close()

	fmt.Println("Temperatuur", temperature.Temperature, "Õhuniiskus", temperature.Humidity, "Ruum", temperature.Room)
}

func handleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	http.HandleFunc("/", router)

	err := configor.Load(&Config, "config.yaml")
	handleError(err)

	fmt.Printf("Starting Thermometrum (+SQL) server in port 8001...\n")
	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatal(err)
	}
}

type Temperature struct {
	Room        string  `json:"room"`
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
}

type Response struct {
	Message string
	Code    int
}
