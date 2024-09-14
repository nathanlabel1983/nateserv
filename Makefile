run:
	go run ./cmd/main.go

test:
	go test ./..

compile:
	go build -o bin/nateserv.exe ./cmd/main.go

clean:
	del /Q bin\*

