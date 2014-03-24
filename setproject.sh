#!/bin/bash

export GOPATH=/vagrant/project
echo "GOPATH set to /vagrant/project"

export PATH=$PATH:$GOPATH/bin
echo "$GOPATH/bin added to PATH"

cd /vagrant
