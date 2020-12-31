start-server:
	go build main.go && ./main

start-client:
	cd client && npm run start

build-client:
	cd client && npm run build

package:
	make build-client
	go build main.go

run: 
	./main