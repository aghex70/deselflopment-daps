.PHONY: local
local: ## Run the application locally
	docker-compose up --build

production: ## Run the application in production
	docker-compose up --build backend frontend asynq-worker asynq-client certbot
