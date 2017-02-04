#!/bin/sh

. ./scripts/_vars.sh

install_go() {
  apt-get update
  apt-get install -y curl --no-install-recommends
  curl -k https://storage.googleapis.com/golang/go1.7.4.${GOOS}-${GOARCH}.tar.gz | tar -C /usr/local -xz
}

type go > /dev/null
if [ $? -ne 0 ]; then
  install_go
  GOBIN=/usr/local/go/bin/go
fi

$GOBIN build -o ./bin/short-url-${GOOS}-${GOARCH}
