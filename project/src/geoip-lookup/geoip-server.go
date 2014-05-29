package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/oschwald/geoip2-golang"
	"html/template"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

type ErrorMessage struct {
	Error string
}

type ReturnData struct {
	Ip          net.IP
	CountryCode string
	CountryName string
	City        string
	Latitude    float64
	Longitude   float64
}

// Load templates.
var htmlTemplates = template.Must(template.ParseFiles("templates/index.html", "templates/error.html"))

// Load the database file
var db, liberr = geoip2.Open("geodata/GeoLite2-City.mmdb")

// Create signal channel for reloading database.
var sigs = make(chan os.Signal, 1)

func getData(r *http.Request) interface{} {
	vars := mux.Vars(r)
	requestIp := ""
	if val, ok := vars["ip"]; ok {
		// IP is provided in URL.
		requestIp = val
	} else {
		// Take IP from Remote Address.
		address := r.Header.Get("X-Real-IP")
		if address == "" {
			address = r.Header.Get("X-Forwarded-For")
		}
		if address == "" {
			address = r.RemoteAddr
			idx := strings.LastIndex(address, ":")
			address = address[:idx]
		}
		requestIp = address
	}

	ip := net.ParseIP(requestIp)

	record, err := db.City(ip)
	if err != nil {
		return ErrorMessage{Error: "Unknown IP"}
	}

	return ReturnData{
		Ip:          ip,
		CountryCode: record.Country.IsoCode,
		CountryName: record.Country.Names["en"],
		City:        record.City.Names["en"],
		Latitude:    record.Location.Latitude,
		Longitude:   record.Location.Longitude}
}

func xmlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	retData := getData(r)
	msg, _ := xml.Marshal(retData)
	w.Write(msg)
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	retData := getData(r)
	msg, _ := json.Marshal(retData)
	w.Write(msg)
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	retData := getData(r)

	switch retData.(type) {
	case ReturnData:
		htmlTemplates.ExecuteTemplate(w, "index.html", retData)
	default:
		htmlTemplates.ExecuteTemplate(w, "error.html", "")
	}
}

func validatePortNumber(port string) bool {
	i, _ := strconv.ParseInt(port, 0, 64)

	return i > 0 && i < 65535
}

func reloadDatabaseListener() {
	// Wait for appropriate signal.
	<-sigs

	// Load the database file.
	db, liberr = geoip2.Open("geodata/GeoLite2-City.mmdb")

	if liberr != nil {
		// Country database was not loaded successfully.
		fmt.Printf("Error: %s\n", liberr.Error())
		return
	}

	fmt.Println("Database reloaded.")
	go reloadDatabaseListener()
}

func main() {
	// Check if data loaded properly.
	if liberr != nil {
		// Country database was not loaded successfully
		fmt.Printf("Error: %s\n", liberr.Error())
		return
	}

	// Use SIGHUP signal for reloading database.
	signal.Notify(sigs, syscall.SIGHUP)
	go reloadDatabaseListener()

	r := mux.NewRouter()
	r.HandleFunc("/{ip}/json", jsonHandler)
	r.HandleFunc("/json", jsonHandler)
	r.HandleFunc("/{ip}/xml", xmlHandler)
	r.HandleFunc("/xml", xmlHandler)
	r.HandleFunc("/{ip}", htmlHandler)
	r.HandleFunc("/", htmlHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("assets")))

	http.Handle("/", r)

	port := ":8080"

	// Check if port was provided via argument and validate port number.
	if len(os.Args) > 1 && validatePortNumber(os.Args[1]) == true {
		port = fmt.Sprintf(":%s", os.Args[1])
	}

	http.ListenAndServe(port, nil)
}
