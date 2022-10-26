-include .env

DATABASE_URL=postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable

all: build down up

build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

psql:
	docker exec -it tracker__db psql -U postgres -d ${POSTGRES_DB}

migrate:
	docker-compose run --rm api migrate -path internal/db/migrations -database ${DATABASE_URL} -verbose up

downgrade:
	docker-compose run --rm api migrate -path internal/db/migrations -database ${DATABASE_URL} -verbose down 1

makemigration:
	docker-compose run --rm --volume=${PWD}/internal/db/migrations:/app/internal/db/migrations api migrate create -ext sql -dir internal/db/migrations -seq ${NAME}
	sudo chown -R ${USER} internal/db/migrations
