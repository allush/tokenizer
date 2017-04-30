ENV ?= tokenizer.local

logs:
	docker-compose -f config/${ENV}/main.yml logs --tail 100 -f -t

stop:
	docker-compose -f config/${ENV}/main.yml stop

start:
	docker-compose -f config/${ENV}/main.yml up -d \
		--force-recreate \
		--remove-orphans

	while true; do 													\
		docker exec postgres.${ENV} psql tokenizer user -c	'SELECT * FROM pg_settings LIMIT 1' >/dev/null 2>&1;	\
		if [ $${?} = 0 ] ; then										\
			break;													\
		fi;															\
		echo -n . ;													\
		sleep 1;													\
	done

build:
	docker-compose -f config/${ENV}/main.yml build

ps:
	docker-compose -f config/${ENV}/main.yml ps

migrate:
	docker exec -i postgres.${ENV} psql tokenizer user < db/schema.sql

seed:
	docker exec -i postgres.${ENV} psql tokenizer user < db/seed.sql

up:
	make build
	make start
	make migrate
	make seed

	# ====================
	# Приложение запущено!
	# ====================
