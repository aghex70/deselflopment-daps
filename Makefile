.PHONY: local
local: ## Run the application locally
	docker-compose up --build backend frontend nginx redis asynq-worker asynq-client

production: ## Run the application in production
	docker-compose -f docker-compose-prod.yml up --build backend frontend webserver redis asynq-worker asynq-client
