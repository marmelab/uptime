setup:
	@cp src/marmelab.com/uptime/conf.json.dist src/marmelab.com/uptime/conf.json
	@docker-compose run watcher go get golang.org/x/net/icmp 
	@docker-compose run api go get github.com/lib/pq
	@docker-compose run webpack npm install
	@echo "Setup completed!"
init_db:
	@docker-compose up -d db
	@docker exec  uptime2_db_1 psql -f /usr/src/db/migration/createTable.sql --username=postgres
	@docker-compose kill db
	@echo "init_db completed"
load_fixtures:
	@docker-compose up -d db
	@docker exec  uptime2_db_1 psql -f /usr/src/db/fixtures/fixtures.sql --username=postgres
	@docker-compose kill db
	@echo "load_fixtures completed"
init_node_server:
	@docker-compose run webpack export NODE_ENV=development
run:
	docker-compose up
clear:
	@docker-compose kill
	@docker-compose rm
