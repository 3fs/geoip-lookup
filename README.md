GeoResolver
============

Small server that responds with Geo data based on Maxmind's GeoIP data for given IP. Written in Go-lang.

## Project setup instructions

1. Clone this repository `git clone git@github.com:3fs/geo-resolver.git`.
2. Move to geo-resolver directory `cd geo-resolver`.
3. Run `vagrant up` to start development environment.
4. Run `vagrant provision` to add all dependencies.
5. Connect to host via `vagrant ssh`.
6. Move to /vagrant directory via `cd /vagrant`.
7. Setup and run project with ` . build.sh`.
8. Go to http://192.168.33.10:8080 to see it in action.
9. Server can be started anytime by using command `geoip-3fs`.

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
