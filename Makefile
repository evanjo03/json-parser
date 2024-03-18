run:
	go run ./cmd/json-parser ./test/simple.json 

build:
	go build -o ./bin/json-parser ./cmd/json-parser
