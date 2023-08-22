run:
	go run main.go

up:
	docker-compose up -d

logs:
	docker-compose logs -f

down:
	docker-compose down

downall:
	docker-compose down --rmi all --volumes --remove-orphans

watch:
	@find ./ -name '*.go' \
	  | entr make run
