run:
	go run cmd/main.go

up:
	docker run --name pg --rm -p 5432:5432 -e POSTGRES_USER=user -e POSTGRES_PASSWORD=test -e POSTGRES_DB=goods -d postgres:latest
	docker run -d --name sel --rm -p 4444:4444 -v /dev/shm:/dev/shm selenium/standalone-chrome

down:
	docker stop pg
	docker stop sel

bash:
	docker exec -it pg bash

selenium:
	 docker run -d --name sel --rm -p 4444:4444 -v /dev/shm:/dev/shm selenium/standalone-chrome

stop:
	docker stop sel