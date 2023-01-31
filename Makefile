include .env
export


start:
	@go run main.go


up:
	@docker-compose up -d


down:
	@docker-compose down


cli:
	@mysql -h ${DB_HOST}
