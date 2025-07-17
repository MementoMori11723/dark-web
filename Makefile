run :
	@docker-compose -f config/compose.yml -p darkweb up --build -d 
stop :
	@docker-compose -f config/compose.yml -p darkweb down  --remove-orphans

dev:
	@PORT=8080 go run .
