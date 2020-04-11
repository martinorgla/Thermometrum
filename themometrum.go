package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)

func router(w http.ResponseWriter, req *http.Request) {
    // TODO: Implement proper router
    if req.URL.Path != "/api/temperature" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    switch req.Method {
    case "GET":
        // fmt.Fprintf(w, req.URL.Path)
        // http.ServeFile(w, r, "form.html")

    case "POST":
        decoder := json.NewDecoder(req.Body)
        var temperature Temperature

        err := decoder.Decode(&temperature);

        if err != nil {
            log.Fatal(err)
        }

        json, err := json.Marshal(temperature)

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

func main() {
    http.HandleFunc("/", router)

    fmt.Printf("Starting Thermometrum server in port 8001...\n")
    if err := http.ListenAndServe(":8001", nil); err != nil {
        log.Fatal(err)
    }
}

type Temperature struct {
    Room string `json:"room"`
	Temperature float32 `json:"temperature"`
	Humidity float32 `json:"humidity"`
}
