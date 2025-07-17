run :
	@docker-compose -f config/compose.yml -p deploy up --build -d 
stop :
	@docker-compose -f config/compose.yml -p deploy down  --remove-orphans

dev:
	@PORT=8080 go run .
