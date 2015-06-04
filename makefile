setup:
	@cp src/marmelab.com/uptime/conf.json.dist src/marmelab.com/uptime/conf.json
	@docker-compose run watcher go get golang.org/x/net/icmp 
	@docker-compose run api go get github.com/lib/pq
	@docker-compose run client npm install
	@echo "Setup completed!"

init_db:
	@docker-compose up -d db
	@docker exec uptime_db_1 psql -f /usr/src/db/migration/createDatabase.sql --username=postgres
	@docker exec uptime_db_1 psql -f /usr/src/db/migration/createTable.sql --username=postgres uptime
	@docker-compose kill db
	@echo "init_db completed"

load_fixtures:
	@docker-compose up -d db
	@docker exec uptime_db_1 psql -c '\connect uptime' -f /usr/src/db/fixtures/fixtures.sql --username=postgres uptime
	@docker-compose kill db
	@echo "load_fixtures completed"

run:
	docker-compose up -d webpack
	docker-compose up -d watcher
	docker-compose up -d --no-deps client
	

clear:
	@docker-compose kill
	@docker-compose rm

init_webpack:
	@docker-compose up -d webpack
	@docker exec uptime_webpack_1 npm install
	@docker-compose kill webpack
	@echo "init_webpack completed"

