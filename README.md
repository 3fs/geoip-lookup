GeoResolver
============

Small server that responds with Geo data based on Maxmind's GeoIP data for given IP. Written in Go-lang.

## Project setup instructions

1. Run `vagrant up` to start development environment.
2. Run `vagrant provision` to add all dependencies.
3. Connect to host via `vagrant ssh`.
4. Move to /vagrant directory via `cd /vagrant`.
5. Setup and run project with ` . build.sh`.
6. Go to http://192.168.33.10:8080 to see it in action.
7. Server can be started anytime by using command `geoip-3fs`.

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
