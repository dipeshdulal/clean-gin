
MIGRATE=docker-compose exec web migrate -path=migration -database "mysql://root:root@tcp(database:3306)/clean_gin" -verbose

migrate-up:
		$(MIGRATE) up
migrate-down:
		$(MIGRATE) down 
force:
		@read -p  "Which version do you want to force?" VERSION; \
		$(MIGRATE) force $$VERSION

goto:
		@read -p  "Which version do you want to migrate?" VERSION; \
		$(MIGRATE) goto $$VERSION

drop:
		$(MIGRATE) drop

create:
		@read -p  "What is the name of migration?" NAME; \
		migrate create -ext sql -seq -dir migration  $$NAME

.PHONY: migrate-up migrate-down force goto drop create

