postgres:
	docker run --name=auth_db -e POSTGRES_USER=olegdjur -e POSTGRES_PASSWORD=qwerty -p 5432:5432 -d postgres

createdb:
	docker exec -it auth_db createdb --username=olegdjur --owner=olegdjur auth_db

migrateup:
	migrate -path migrations -database "postgresql://olegdjur:qwerty@localhost:5432/auth_db?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgresql://olegdjur:qwerty@localhost:5432/auth_db?sslmode=disable" -verbose down

run:
	go run ./cmd/main.go

docker:
	docker-compose up