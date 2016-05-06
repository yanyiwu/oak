build:
	go build -v ./...
test:
	go test $(glide novendor)
start:
	./settledb -alsologtostderr=true
