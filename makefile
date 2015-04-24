setup:
	docker-compose run watcher go get golang.org/x/net/icmp

run:
	docker-compose up
