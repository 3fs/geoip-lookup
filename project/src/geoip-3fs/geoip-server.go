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
	"strings"
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

// Load the database file
var db, liberr = geoip2.Open("geodata/GeoLite2-City.mmdb")

// Load templates.
var htmlTemplates = template.Must(template.ParseFiles("bin/templates/index.html", "bin/templates/error.html"))

func getData(r *http.Request) interface{} {
	vars := mux.Vars(r)
	requestIp := ""
	if val, ok := vars["ip"]; ok {
		// IP is provided in URL.
		requestIp = val
	} else {
		// Take IP from Remote Address.
		address := strings.Split(r.RemoteAddr, ":")
		requestIp = address[0]
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
	retData := getData(r)
	msg, _ := xml.Marshal(retData)
	w.Write(msg)
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

func main() {

	if liberr != nil {
		// Country database was not loaded successfully
		fmt.Printf("Error: %s\n", liberr.Error())
		return
	}

	r := mux.NewRouter()
	r.HandleFunc("/{ip}/json", jsonHandler)
	r.HandleFunc("/json", jsonHandler)
	r.HandleFunc("/{ip}/xml", xmlHandler)
	r.HandleFunc("/xml", xmlHandler)
	r.HandleFunc("/{ip}", htmlHandler)
	r.HandleFunc("/", htmlHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("bin/assets")))

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
