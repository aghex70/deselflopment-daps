.PHONY: local
local: ## Run the application locally
	docker-compose -f docker-compose-local.yml up --build backend frontend webserver asynq-worker asynq-client

production: ## Run the application in production
	docker-compose up --build backend frontend asynq-worker asynq-client certbot
