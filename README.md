geoip-lookup
============

Small server written in Go that responds with GeoIP data for given IP in chosen format (xml/json/html)
[![Build Status](http://188.226.235.163:8080/github.com/3fs/geoip-lookup/status.svg?branch=master)](http://188.226.235.163:8080/github.com/3fs/geoip-lookup)

This product includes GeoLite2 data created by MaxMind, available from
<a href="http://www.maxmind.com">http://www.maxmind.com</a>.

Development environment built using https://github.com/FreakyDazio/vagrant-golang.


## Project setup instructions

1. Clone this repository `git clone git@github.com:3fs/geoip-lookup.git`.
2. Move to geoip-lookup directory `cd geoip-lookup`.
3. Run `vagrant up` to start development environment.
4. Run `vagrant provision` to add all dependencies.
5. Connect to host via `vagrant ssh`.
6. Move to /vagrant directory via `cd /vagrant`.
7. Setup and run project with ` . build.sh`.
8. Go to http://192.168.33.10:8080 to see it in action.
9. Server can be started anytime by using command `geoip-lookupd`.
10. Default port number is 8080. However, port number can be set using first argument of command. Example: `geoip-lookupd 3333`.

## Usage

Any IP can be queried via URL: `http://192.168.33.10:8080/{IP}/{format}`.
Response can be returned in HTML, JSON or XML format

 - http://192.168.33.10:8080/212.58.246.93
 - http://192.168.33.10:8080/212.58.246.93/json
 - http://192.168.33.10:8080/212.58.246.93/xml

If IP is not provided in URL, IP from RemoteAddr is used (visitor's IP).

 - http://192.168.33.10:8080
 - http://192.168.33.10:8080/json
 - http://192.168.33.10:8080/xml
