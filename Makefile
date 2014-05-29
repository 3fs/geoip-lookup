SHA := $(shell git rev-parse --short HEAD)

build:
		go get github.com/oschwald/geoip2-golang
		go get github.com/gorilla/mux
		mkdir geoip-lookup-${SHA}
		go build -o geoip-lookup-${SHA}/geoip-lookupd geoip-lookup
		cp -R ${GOPATH}/bin/* ${GOPATH}/geodata geoip-lookup-${SHA}/
		tar czf geoip-lookup-${SHA}.tar.gz geoip-lookup-${SHA}
