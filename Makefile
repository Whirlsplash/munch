fmt:
	go fmt github.com/Whirlsplash/munch...

run: fmt
	go run munch.go

build: fmt
	go build
