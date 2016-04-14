# This is how we want to name the binary output
BINARY=2d

# These are the values we want to pass for Version and BuildTime
VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X github.com/tommymalmqvist/golang-hourglass/core.Version=${VERSION} -X github.com/tommymalmqvist/golang-hourglass/core.BuildTime=${BUILD_TIME}"

build:
	go build ${LDFLAGS} -o ${BINARY} 2d.go

test:
	sh test.sh | ./2d

clean:
	rm 2d
