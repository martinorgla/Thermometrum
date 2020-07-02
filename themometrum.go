package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/configor"
	"log"
	"strconv"
	_ "time"
)

var db *sql.DB

func main() {
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)

	err := configor.Load(&Config, "config.yaml")
	handleError(err)

	openDatabaseConnection()

	log.Printf("Starting Thermometrum on port 8001...\n")
	log.Println("Ver. " + Config.AppVersion)

	setupRouter()
	// go setupRouter()
	//
	// // Leave the app running forever
	// for true {
	// 	time.Sleep(time.Minute * 1)
	// }
}

func openDatabaseConnection() *sql.DB {
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

func getLastTemperature() Temperature {
	selDB, err := db.Query("SELECT room, temperature, humidity, time FROM temperatures ORDER BY id DESC LIMIt 1")
	handleError(err)
	defer db.Close()

	temperature := Temperature{}

	for selDB.Next() {
		err = selDB.Scan(&temperature.Room, &temperature.Temperature, &temperature.Humidity, &temperature.Date)
		handleError(err)
	}

	return temperature
}

func getAllTemperatures() []Temperature {
	selDB, err := db.Query("SELECT room, temperature, humidity, time FROM temperatures ORDER BY id DESC")
	handleError(err)
	defer db.Close()

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

type Temperature struct {
	Room        string  `json:"room"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Date        string  `json:"date"`
}

type Response struct {
	Message string
	Code    int
}
