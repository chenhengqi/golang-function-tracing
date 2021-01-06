APP=gofunctrace

build: dep
	go build -mod=vendor -o $(APP)

dep:
	go mod tidy
	go mod vendor

lint:
	go fmt ./...
	go vet ./...
	gofmt -e -l .
	golint ./...

clean:
	rm -f $(APP)
