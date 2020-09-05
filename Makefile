.PHONY : build
build :
	go build -o ./bin/gonel ./cmd/gonel/*.go

.PHONY : run
run : build
	./bin/gonel

.PHONY : test
test :
	go test ./...

.PHONY : install
install : 
	cp ./bin/gonel /usr/local/bin

