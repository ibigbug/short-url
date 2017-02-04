. ./scripts/_vars.sh
export PORT=8000
export ADDR=localhost
exec ./bin/short-url-${GOOS}-${GOARCH}
