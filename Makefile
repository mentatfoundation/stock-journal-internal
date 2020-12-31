server-local:
	go build main.go && ./main

client-local:
	cd client && npm run start

