.PHONY: local
local: ## Run the application locally
	docker-compose up --build backend

production: ## Run the application in production
	docker-compose up --build backend frontend asynq-worker asynq-client certbot

stop: ## Stop the application
	docker-compose down

start:
	docker-compose up --build backend frontend
