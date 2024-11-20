build: 
	go build -o numutils main.go

test:
	go test -v ./...	

clean:
	rm -f numutils

.PHONY: build test clean