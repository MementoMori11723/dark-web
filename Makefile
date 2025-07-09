run :
	@docker-compose -f config/compose.yml -p deploy up --build -d 
	@docker exec -it tor sh -c "cat /var/lib/tor/hidden_service/hostname"
stop :
	@docker-compose -f config/compose.yml -p deploy down  --remove-orphans

dev:
	@go run .
