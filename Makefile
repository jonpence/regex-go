build: src/*.go
	go build -o regex src/*.go

run: src/*.go
	go run src/*.go

test: src/*.go test/*.go
	go test test/*.go

clean:
	rm regex
