setup:
	@cp src/marmelab.com/uptime/url.json.dist src/marmelab.com/uptime/url.json
	@docker-compose run watcher go get golang.org/x/net/icmp
	@echo "Setup completed!"

run:
	docker-compose up
