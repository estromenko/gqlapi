
.DEFAULT_GOAL=run

run:
	docker-compose up

runb:
	docker-compose up --build

psql:
	docker-compose run db psql -U postgres -W --dbname postgres -h 0.0.0.0 --port 5432