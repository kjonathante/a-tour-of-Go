# a-tour-of-Go
A Tour of Go

Running code inside docker
docker run -ti --rm -v "$PWD":/usr/go -w /usr/src/myapp golang:1.8 go run basics/13-type-conversions.go

docker run -ti --rm -v "$PWD":/go golang:1.8 go run 18-exercise-stringers.go