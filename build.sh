#!/bin/bash

pathmunge () {
  if ! echo $PATH | /bin/egrep -q "(^|:)$1($|:)" ; then
    if [ "$2" = "after" ] ; then
      PATH=$PATH:$1
    else
      PATH=$1:$PATH
    fi
  fi
}

echo -e "\e[33mBuilding GeoResolver\e[0m"

export GOPATH=/vagrant/project
echo -e "\e[33m[1] GOPATH set to /vagrant/project.\e[0m"

pathmunge $GOPATH/bin
echo -e "\e[33m[2] $GOPATH/bin added to PATH.\e[0m"

/usr/local/go/bin/go get github.com/oschwald/geoip2-golang
echo -e "\e[33m[3] GO: geoip2-golang package installed.\e[0m"

# Check if database exists and if not - download it.
if [ ! -f  "project/geodata/GeoLite2-City.mmdb" ] ; then
    (cd $GOPATH/geodata \
    && wget -q http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.mmdb.gz \
    && gunzip -f  GeoLite2-City.mmdb.gz \
    && echo -e "\e[33m[3.1] GeoLite2-City database downloaded.\e[0m")
fi

/usr/local/go/bin/go get github.com/gorilla/mux
echo -e "\e[33m[4] GO: gorilla/mux package installed.\e[0m"

(cd $GOPATH \
    && go install geoip-3fs \
    && echo  -e "\e[33m[5] Project built." \
    && echo -e '[6] Running server on port 8080. Quit: Ctrl+C.\e[0m' \
    && $GOPATH/bin/geoip-3fs > /dev/tty)

