build: cmd/*/*.go internal/*/*.go
	go build -o rego cmd/rego/*.go

run: build
	./rego

test: cmd/*/*.go internal/*/*.go
	go test internal/set/*.go
	go test internal/deque/*.go
	go test internal/regex/*.go

clean:
	rm rego
