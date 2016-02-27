all:
	go build -v
test:
	go test -v ./...
godep:
	godep save -r ./...
start:
	./settledb -alsologtostderr=true
