export GOARCH="386"
if [ `uname -m` = "x86_64" ]; then
  GOARCH="amd64"
fi

export GOOS=`uname -s | tr [:upper:] [:lower:]`
GOBIN=go


