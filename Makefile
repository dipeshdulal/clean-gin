include .env

MIGRATE=docker-compose exec web sql-migrate

ifeq ($(p),host)
 	MIGRATE=sql-migrate
endif

migrate-status:
	$(MIGRATE) status

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down 

redo:
	@read -p  "Are you sure to reapply the last migration? [y/n]" -n 1 -r; \
	if [[ $$REPLY =~ ^[Yy] ]]; \
	then \
		$(MIGRATE) redo; \
	fi

create:
	@read -p  "What is the name of migration?" NAME; \
	${MIGRATE} new $$NAME

.PHONY: migrate-status migrate-up migrate-down redo create

